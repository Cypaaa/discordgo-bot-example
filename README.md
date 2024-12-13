## Setup
Set your discord token in `./src/.env` under the variable `DISCORD_TOKEN`

Run `go run src/raven.go` to start your bot.

__Remember to rename raven.go as you wish and update the project name.<br />
I set it to Raven since my base project was name Raven.__

## Add commands:
- make a new command file under `./src/discord/commands`
- make an instance of Command struct
- in `./src/discord/commands/command.go`, add your instance of Command in `CommandList` variable

## Add events:
- make a new event file under `./src/discord/events`
- make an instance of Event struct
- in `./src/discord/events/event.go`, add your instance of Event in `EventList` variable
