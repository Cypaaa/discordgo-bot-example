package commands

import (
	"github.com/bwmarrin/discordgo"
)

var RepeatCommand = &Command{
	Command: &discordgo.ApplicationCommand{
		Name:        "repeat",
		Description: "repeat a message",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "message",
				Description: "A message to repeat",
				Required:    true,
			},
		},
	},
	Exec: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		options := GetOptionMap(i.ApplicationCommandData().Options)
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags:   discordgo.MessageFlagsEphemeral,
				Content: options["message"].StringValue(),
			},
		})
	},
}
