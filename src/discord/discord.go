package discord

import (
	"log"
	"os"
	"os/signal"
	"raven/src/discord/commands"
	"raven/src/discord/events"
	"syscall"

	"github.com/bwmarrin/discordgo"
	_ "github.com/joho/godotenv/autoload"
)

func Init() {
	s, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	s.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentsGuildVoiceStates
	// events should be registyered BEFORE opening connection
	events.EventList.Reload(s)

	err = s.Open()
	if err != nil {
		log.Fatal("Authentication failed:", err)
	}
	defer s.Close()

	// commands should be send after since it relies on s
	commands.CommandList.Reload(s)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-c
}
