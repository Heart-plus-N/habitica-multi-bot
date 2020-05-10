package observer_pattern

type ActivityType uint8

const (
	TaskEvent      ActivityType = 8
	GroupChatEvent ActivityType = 4
	UserEvent      ActivityType = 2
	QuestEvent     ActivityType = 1
)

type Observer interface {
	Initiate(ActivityType, interface{})
	GetInterest() ActivityType
}
