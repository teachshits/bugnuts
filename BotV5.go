package main
// The v4 Bot -- Marginally less Terrible!!!!

import (
	"fmt"
	"os"
	"rand"
	"sort"
	"log"
)

type BotV5 struct {
	P *Parameters

	Explore map[Location]*Target

	IdleAnts []int
}

//NewBot creates a new instance of your bot
func NewBotV5(s *State) Bot {
	mb := &BotV5{
		P:       ParameterSets["V5"],
		Explore: make(map[Location]*Target, 10),
	}

	mb.MakeExplorers(s, 1.25)
	return mb
}

func (bot *BotV5) MakeExplorers(s *State, scale float64) {

	// Set an initial group of targetpoints
	stride := int(float64(s.Radius) * scale)
	if stride <= 0 {
		stride = 2 * s.Radius
	}
	for r := 5; r < s.Rows; r += stride {
		for c := 5; c < s.Cols; c += stride {
			loc := s.Map.ToLocation(Point{r: r, c: c})
			bot.Explore[loc] = &Target{item: EXPLORE, loc: loc, count: 1, pri: bot.P.Priority[EXPLORE]}
		}
	}
}

func (bot *BotV5) ExploreUpdate(s *State) {
	// Any explore point which is visible should be nuked
	for loc, _ := range bot.Explore {
		if s.Map.Seen[loc] == s.Turn {
			bot.Explore[loc] = &Target{}, false
		}
	}
}

func (bot *BotV5) DoTurn(s *State) os.Error {
	// TODO this still seems clunky.  need to figure out a better way
	s.FoodUpdate(bot.P.ExpireFood)
	bot.ExploreUpdate(s)

	tset := TargetSet{}

	// Run through explore targets
	for loc, tgt := range bot.Explore {
		tgt.count = 1
		tset[loc] = tgt
		if Debug > 4 {
			log.Printf("Setting explore target at %v: %v", s.Map.ToPoint(loc), tgt)
		}
	}

	// Generate list of food and enemy hill points.
	for _, loc := range s.FoodLocations() {
		if Debug > 4 {
			log.Printf("adding target %v(%d) food pri %d", s.Map.ToPoint(loc), loc, bot.P.Priority[FOOD])
		}
		tset.Add(FOOD, loc, 1, bot.P.Priority[FOOD])
	}

	// TODO handle different priorities for different enemies.
	eh := s.EnemyHillLocations()
	idle := 0
	if s.Turn > 10 && len(eh) > 0 {
		//log.Printf("IDLE: %d : %d : %d : %v", s.Turn, len(bot.IdleAnts), len(eh), bot.IdleAnts)
		idle = bot.IdleAnts[s.Turn-2] / len(eh)
		if idle > 15 {
			idle = 15
		}
	}
	for _, loc := range eh {
		tset.Add(HILL1, loc, 1+idle, bot.P.Priority[HILL1])
	}

	if Debug > 4 {
		log.Printf("Target set %v", tset)
	}

	// List of available ants
	ants := make(map[Location]int, len(s.Ants[0]))
	for k, v := range s.Ants[0] {
		ants[k] = v
	}

	// TODO remove dedicated ants eg sentinel, capture, defense guys

	moves := make(map[Location]Direction, len(ants))

	for iter := 0; iter < 15 && len(ants) > 0 && tset.Pending() > 0; iter++ {
		if Debug > 4 {
			log.Printf("Location iteration %d, ants: %d, tset.Pending %d", iter, len(ants), tset.Pending())
		}

		// TODO: Here should update map for fixed ants.
		f, _, _ := MapFill(s.Map, tset.Active())

		// Make hills infinitely unattractive
		/* This DOES NOT WORK - can make the hill a local minima.
			for _, loc := range s.MyHillLocations() {
			f.Depth[loc] = math.MaxUint16
		}
		*/

		// Build list of locations sorted by depth
		ll := make(map[int][]Location)
		var dl []int
		for loc, _ := range ants {
			depth := int(f.Depth[loc])
			if _, ok := ll[depth]; !ok {
				ll[depth] = make([]Location, 0)
				dl = append(dl, int(depth))
			}
			ll[depth] = append(ll[depth], loc)
		}
		sort.Sort(IntSlice(dl))

		for _, depth := range dl {
			// loop over depths
			for _, loc := range ll[depth] {
				// for a given depth loop over ant locations

				// p := s.Map.ToPoint(loc)
			STEP:
				for _, d := range rand.Perm(4) {
					// find a direction we can step in thats stepable.
					np := s.Map.PointAdd(s.Map.ToPoint(loc), Steps[d])
					nl := s.Map.ToLocation(np)

					if f.Depth[nl] < uint16(depth) &&
						(s.Map.Grid[nl] == LAND || s.Map.Grid[nl].IsEnemyHill() || s.Map.Grid[nl] == FOOD) {
						// We have a valid next step, path in to dest and see if
						// We should remove ant and possibly target.

						// Need to handle the special case where food spawns next to us
						// in that case we don't move, mark the food as a block, and
						// remove ourself and the food from the list

						if s.Map.Grid[nl] == FOOD {
							// We are next to food, remove this ant from the 
							// available list, mark this food as taken, and mark the 
							// food as a block.
							s.Map.Grid[nl] = BLOCK
							moves[loc] = Direction(d)
							if Debug > 4 {
								log.Printf("Removing %v food adjacent", s.Map.ToPoint(loc))
							}
							ants[loc] = 0, false
							tset[nl].count = 0
						} else {
							endloc, steps := f.PathIn(nl)
							tgt, ok := tset[endloc]
							if Debug > 4 {
								log.Printf("Found target %v -> %v, end %v %d steps, tgt: %v", s.Map.ToPoint(loc), s.Map.ToPoint(nl), s.Map.ToPoint(endloc), steps, tgt)
							}
							if !ok {
								if Debug > 3 {
									log.Printf("pathin from %v(%d)->%v to %v(%d) no matching target.",
										s.Map.ToPoint(loc), loc, s.Map.ToPoint(nl), s.Map.ToPoint(endloc), endloc)
								}
							} else if steps >= 0 && tgt.count > 0 {
								moves[loc] = Direction(d)
								tgt.count -= 1
								s.Map.Grid[nl] = BLOCK
								ants[loc] = 0, false
							}
						}
						break STEP
					}
				}
			}
		}

	}
	// TODO If we have more ants than targets we have bored ants, uptick the # in on hills, add some exploring etc.

	//log.Printf("%d ants with nothing to do", len(ants))
	bot.IdleAnts = append(bot.IdleAnts, len(ants))

	for loc, d := range moves {
		p := s.Map.ToPoint(loc)
		fmt.Fprintf(os.Stdout, "o %d %d %s\n", p.r, p.c, DirectionChar[d])
	}

	fmt.Fprintf(os.Stdout, "go\n")

	// TODO tiebreak on global goals.

	return nil
}