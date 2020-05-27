package bot

import (
	"fmt"
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

	if valueSplit[0] == "@Utility_Bot" {
		// Since we know it's a command we need to
		// find a group to post the message in.
		group, err := jsonparser.GetString(body, "group", "id")
		log.Println(group)

		responseMessage := fmt.Sprintf("Utility_Bot is up as of: %s", time.Now().UTC().String())

		api := gabit.NewHabiticaAPI(nil, "", nil)
		_, err = api.Authenticate(sc.HabiticaUsername, sc.HabiticaPassword)
		if err != nil {
			log.Println(err)
		}
		_, err = api.PostMessage(group, responseMessage)
		if err != nil {
			log.Println(err)
		}
	}

}

func (b Bot) GetInterest() ActivityType {
	return GroupChatEvent
}
