package events

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var ReadyEvent = &Event{
	Name: "ready",
	Once: false,
	Exec: func(s *discordgo.Session, r *discordgo.Ready) {
		fmt.Println("Client ready and logged as", s.State.User.Username)
	},
}
