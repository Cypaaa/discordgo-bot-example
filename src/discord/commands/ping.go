package commands

import (
	"github.com/bwmarrin/discordgo"
)

var PingCommand = &Command{
	Command: &discordgo.ApplicationCommand{
		Name:        "ping",
		Description: "Reply Pong!",
	},
	Exec: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Pong!",
			},
		})
	},
}
