package events

import (
	"fmt"
	"raven/src/discord/commands"

	"github.com/bwmarrin/discordgo"
)

var InteractionCreateEvent = &Event{
	Name: "interactionCreate", // required for logs only
	Once: false,
	Exec: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if c, ok := commands.CommandList.Find(i.Interaction.ApplicationCommandData().ID); ok {
			c.Exec(s, i)
		}
	},
}
