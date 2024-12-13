package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

// map of command with their IDs upfront. The ID placeholder can be anything
// as it will be replaced after command registration
var CommandList = Commands{
	"PingCommand": PingCommand,
}

type Commands map[string]*Command

type Command struct {
	Command *discordgo.ApplicationCommand
	Exec    func(*discordgo.Session, *discordgo.InteractionCreate)
}

func (cs Commands) Find(s string) (*Command, bool) {
	if c, ok := cs[s]; ok {
		return c, ok
	}
	return nil, false
}

func (cs *Commands) Reload(s *discordgo.Session) {
	// new cs, will replace cs at the end of the function
	var ncs = Commands{}
	// registered existing commands
	var recs, _ = s.ApplicationCommands(s.State.User.ID, "")
	// map of existing commands identifiable by their names
	var ecs = map[string]*discordgo.ApplicationCommand{}
	for _, ec := range recs {
		ecs[ec.Name] = ec
	}

	for _, c := range *cs {
		// if ec exists, update it, else create it
		if ec, ok := ecs[c.Command.Name]; ok {
			uc, err := s.ApplicationCommandEdit(s.State.User.ID, "", ec.ID, c.Command)
			if err != nil {
				log.Fatal("Error creating command:", err)
			}
			c.Command = uc
		} else {
			cc, err := s.ApplicationCommandCreate(s.State.User.ID, "", c.Command)
			if err != nil {
				log.Fatal("Error creating command:", err)
			}
			c.Command = cc
		}

		// save the command under its name in the map
		// and remove the placeholder id entry
		ncs[c.Command.ID] = c
	}
	*cs = ncs
}

func GetOptionMap(options []*discordgo.ApplicationCommandInteractionDataOption) map[string]*discordgo.ApplicationCommandInteractionDataOption {
	var om = map[string]*discordgo.ApplicationCommandInteractionDataOption{}
	for _, o := range options {
		om[o.Name] = o
	}
	return om
}
