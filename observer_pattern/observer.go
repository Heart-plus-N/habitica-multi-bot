package observer_pattern

type ActivityType uint8

const (
	TaskEvent      ActivityType = 8
	GroupChatEvent ActivityType = 4
	UserEvent      ActivityType = 2
	QuestEvent     ActivityType = 1
)

// Observer has all the functions that will be used by a Reporter
// These functions should contain limited business logic;
// instead, they should wrap functions on another package.
// This way, the other package can be tested independently
// of the interface (and the bot).
type Observer interface {
	Initiate(ActivityType, []byte, SharedConfig)
	GetInterest() ActivityType
}
