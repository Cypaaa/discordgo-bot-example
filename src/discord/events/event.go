package events

import (
	"log"
	"reflect"

	"github.com/bwmarrin/discordgo"
)

// slice of event
var EventList = Events{
	ReadyEvent,
	InteractionCreateEvent,
}

type Events []*Event

type Event struct {
	Name string
	Once bool
	Exec interface{}
}

func (e Event) Handle(args ...interface{}) {
	v := reflect.ValueOf(e.Exec)
	if v.Kind() != reflect.Func {
		log.Fatal(e.Name + " event: Exec is not a func")
		return
	}
	// Build arguments for Exec
	in := make([]reflect.Value, len(args))
	for i, arg := range args {
		in[i] = reflect.ValueOf(arg)
	}
	// Call Exec
	v.Call(in)
}

func (es Events) Reload(s *discordgo.Session) {
	for _, e := range es {
		if e.Once {
			s.AddHandlerOnce(e.Exec)
		} else {
			s.AddHandler(e.Exec)
		}
	}
}
