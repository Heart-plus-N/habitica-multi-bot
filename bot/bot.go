package bot

import (
	"database/sql"

	. "gitlab.com/bfcarpio/gabit"
)

// Used with "Interests"
type EventType uint8

const (
	AllEvents EventType = 255

	// Unused Event ActivityType = 128
	// Unused Event ActivityType = 64
	// Unused Event ActivityType = 32
	// Unused Event ActivityType = 16

	TaskEvent      EventType = 8
	GroupChatEvent EventType = 4
	UserEvent      EventType = 2
	QuestEvent     EventType = 1
	NoEvents       EventType = 0
)

type BotFactory interface {
	NewBot(SharedConfig, interface{}) Bot
}

type Bot interface {
	// Bots should return an EventType of what
	// they are able and willing to process.
	GetInterest() EventType
	// Gets called when a matching activity type is received.
	Execute(Webhook)
}

// Pass in all EventTypes that a bot is interested in to create
// an aggregate EventType.
func BuildInterest(es ...EventType) EventType {
	var acc uint8
	acc = 0
	for _, e := range es {
		acc = acc + uint8(e)
	}
	return EventType(acc)
}

type SharedConfig struct {
	AccountUsername string
	Api             *HabiticaAPI
	Db              *sql.DB `mapstructure:"Db"`
}
