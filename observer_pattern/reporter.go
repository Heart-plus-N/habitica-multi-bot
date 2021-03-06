package observer_pattern

import (
	"log"
	"time"
)

type reporter struct {
	observers []Observer
	config    SharedConfig
}

func NewReporter(sc SharedConfig) reporter {
	return reporter{config: sc}
}

func (r *reporter) Subscribe(o Observer) {
	r.observers = append(r.observers, o)
	log.Println(r.observers)
}

func (r reporter) Notify(e ActivityType, body []byte) {
	log.Println("Notifying Observers!")
	for _, o := range r.observers {
		go o.Initiate(e, body, r.config)
		time.Sleep(500 * time.Millisecond)
	}
}
