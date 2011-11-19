package MyBot

import (
	"log"
	. "bugnuts/state"
	. "bugnuts/parameters"
)

type ABot struct {
	Key    string
	Desc   string
	PSet   string
	NewBot func(*State, *Parameters) Bot
}

var botList map[string]ABot

func RegisterABot(bot ABot) {
	botList[bot.Key] = bot
}

func UnregisterABot(key string) {
	botList[key] = ABot{}, false
}

func BotList() []string {
	bl := make([]string, 0, len(botList))
	for k, b := range botList {
		bl = append(bl, k+": "+b.Desc)
	}

	return bl
}

func BotGet(k string) ABot {
	b, ok := botList[k]
	if !ok {
		log.Panicf("Invalid bot %s", k)
	}

	return b
}
