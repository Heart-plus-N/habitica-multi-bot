package main

import (
	"errors"
	"fmt"

	. "github.com/Heart-plus-N/habitica-multi-bot/bot"
	log "github.com/amoghe/distillog"
	. "gitlab.com/bfcarpio/gabit"
)

type reporter struct {
	bots []Bot
}

func translateWebhookType(w WebhookType) (EventType, error) {
	log.Debugln("Translating WebhookType: %s", w)
	switch w {
	case TaskActivity:
		return TaskEvent, nil
	case GroupChatReceived:
		return GroupChatEvent, nil
	case UserActivity:
		return UserEvent, nil
	case QuestActivity:
		return QuestEvent, nil
	default:
		return 0, errors.New(fmt.Sprintf("unknown WebhookType: %s", w))
	}

}

func interested(a, i EventType) bool {
	return a&i > 0
}

func (r reporter) Notify(hook Webhook) {
	hookInterest, err := translateWebhookType(hook.Type)
	if err != nil {
		log.Warningln("Couldn't translate WebhookType to activity type")
		return
	}
	log.Debugln("Notifying Bots!")
	for _, b := range r.bots {
		if interested(hookInterest, b.GetInterest()) {
			go b.Execute(hook)
		}
	}
}

func (r *reporter) Add(b Bot) error {
	log.Debugln("Adding Bot: ", b)
	r.bots = append(r.bots, b)
	return nil
}
