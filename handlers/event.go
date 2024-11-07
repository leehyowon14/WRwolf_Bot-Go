package handler

import (
	clientEvent "WRwolf_bot-Go/events/client"
	guildEvent "WRwolf_bot-Go/events/guild"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func RegisterEventHandlers(discordSession *discordgo.Session) {
	for _, event := range guildEvent.GuildEventArr {
		discordSession.AddHandler(event.EventFunc)
		fmt.Println(event.EventName + " is registered")
	}
	discordSession.AddHandler(clientEvent.Ready)
	fmt.Println("Ready is registered")
}
