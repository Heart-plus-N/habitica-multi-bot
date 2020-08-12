package bot

import (
	"fmt"
	"log"
	"strings"
	"time"

	. "github.com/Heart-plus-N/habitica-multi-bot/observer_pattern"
	// Using json parser because I know I'll
	// only need a handful of params from the webhook.
	// Should be faster than marshalling everyting into a struct.
	// Also, coupling the struct definitions is brittle!
	"github.com/buger/jsonparser"
	"gitlab.com/bfcarpio/gabit"
)

type Bot struct {
	Name string
}

func (b Bot) Initiate(at ActivityType, body []byte, sc SharedConfig) {
	log.Println("Bot Utils")

	// Get the new chat message
	value, err := jsonparser.GetString(body, "chat", "unformattedText")
	if err != nil {
		log.Println("Couldn't parse unformattedText")
		log.Println(err)
		return
	}

	// Split the message so we can use it and identify the parts of
	// a command.
	valueSplit := strings.Fields(value)

	if valueSplit[0] == "@Utility_Bot" {
		// Since we know it's a command we need to
		// find a group to post the message in.
		group, err := jsonparser.GetString(body, "group", "id")
		// Find the user so we can ping them back!
		user, err := jsonparser.GetString(body, "chat", "username")
		if err != nil {
			log.Println(err)
			return
		}

		// Responsd in the chat with a nice message and ping the user who called on the bot
		responseMessage := fmt.Sprintf("@%s Utility_Bot is up as of: %s", user, time.Now().UTC().String())
		// log.Println(responseMessage)

		api := gabit.NewHabiticaAPI(nil, "", nil)
		api.Authenticate(sc.HabiticaUsername, sc.HabiticaPassword)

		_, err = api.PostMessage(group, responseMessage)
		if err != nil {
			log.Println(err)
		}
	}

}

func (b Bot) GetInterest() ActivityType {
	return GroupChatEvent
}
