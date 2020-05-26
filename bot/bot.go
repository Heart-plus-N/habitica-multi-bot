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

	// Get the new chat message
	value, err := jsonparser.GetString(body, "unformattedText")
	if err != nil {
		log.Println("Couldn't parse body")
		log.Println(err)
	}

	// Split the message so we can use it and identify the parts of
	// a command.
	valueSplit := strings.Fields(value)

	if valueSplit[0] == "!ub" {
		// Since we know it's a command we need to
		// find a group to post the message in.
		group, err := jsonparser.GetString(body, "group", "id")

		api := gabit.NewHabiticaAPI(nil, "", nil)
		api.Authenticate(sc.HabiticaUsername, sc.HabiticaPassword)
		_, err = api.PostMessage(group, "Utility_Bot is up as of: "+time.Now().String())
		if err != nil {
			log.Println(err)
		}
	}

}

func (b Bot) GetInterest() ActivityType {
	return GroupChatEvent
}
