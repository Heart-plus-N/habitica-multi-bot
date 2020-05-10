package observer_pattern

import (
	"log"
	"time"
)

type Reporter struct {
	observers []Observer
}

func (r *Reporter) Subscribe(o Observer) {
	r.observers = append(r.observers, o)
	log.Println(r.observers)
}

func (r *Reporter) Notify(e ActivityType, body interface{}) {
	for _, o := range r.observers {
		go o.Initiate(e, body)
		time.Sleep(500 * time.Millisecond)
	}
}
