package bot

import (
	"log"
	"strings"
	"time"

	. "github.com/Heart-plus-N/habitica-multi-bot/observer_pattern"
	"github.com/buger/jsonparser"
	"gitlab.com/bfcarpio/gabit"
)

type Bot struct {
	Name string
}

func (b Bot) Initiate(at ActivityType, body []byte, sc SharedConfig) {
	log.Println("Bot Utils")
	value, err := jsonparser.GetString(body, "unformattedText")
	if err != nil {
		log.Println("Couldn't parse body")
		log.Println(err)
	}

	valueSplit := strings.Fields(value)

	if valueSplit[0] == "!ub" {
		api := gabit.NewHabiticaAPI(nil, "", nil)
		api.Authenticate(sc.HabiticaUsername, sc.HabiticaPassword)
		_, err := api.PostMessage("party", "Utility_Bot is up as of: "+time.Now().String())
		if err != nil {
			log.Println(err)
		}
	}

}

func (b Bot) GetInterest() ActivityType {
	return GroupChatEvent
}
