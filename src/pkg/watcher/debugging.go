package watcher

import (
	"log"
)

const (
	DBG_GatherTime = iota + 1
	DBG_TurnTime
	DBG_Timeouts
	DBG_Iterations
	DBG_Start
	DBG_Threat
	DBG_Movement
	DBG_MoveErrors
	DBG_Combat
	DBG_PathIn
	DBG_PathInDefense
	DBG_Targets
	DBG_Sample
	DBG_BorderTargets
	DBG_Symmetry
	DBG_Metrics
	DBG_Statistics
	maxDBG
)

var Debug = make([]bool, maxDBG)

var WS *Watches

// DebugLevels are cumulative, but a 0 resets all to false and a -Level unsets.
var debugLevels = [][]int{
	0: {DBG_GatherTime},
	1: {DBG_TurnTime, DBG_Combat},
	2: {DBG_Iterations, DBG_Symmetry},
	3: {DBG_Targets, DBG_Threat},
	4: {0, DBG_Movement},
}

func SetDebugLevel(dlev int) {
	if dlev >= len(debugLevels) {
		log.Panicf("Max defined debug level is %d", len(debugLevels)-1)
	}
	for d := 0; d <= dlev; d++ {
		for _, dbg := range debugLevels[d] {
			if dbg == 0 {
				for i := range Debug {
					Debug[i] = false
				}
			} else if dbg > 0 {
				Debug[dbg] = true
			} else {
				Debug[dbg] = false
			}
		}
	}
}
