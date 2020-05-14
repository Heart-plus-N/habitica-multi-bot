package quest_queue

import (
	"log"

	. "github.com/Heart-plus-N/habitica-multi-bot/observer_pattern"
)

type QuestQueue struct {
	Name     string
	interest ActivityType
}

func (q QuestQueue) Initiate(at ActivityType, body interface{}) {
	log.Println(q.Name)
}

func (q QuestQueue) GetInterest() ActivityType {
	return q.interest
}
