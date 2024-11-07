package handler

import (
	messageCommand "WRwolf_bot-Go/command/message"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var CommandMap = make(map[string]interface{})

func RegisterMessageCommandHandlers(s *discordgo.Session) {
	for _, command := range messageCommand.CommandArr {
		CommandMap[command.Name] = command.CommandFunc
		for _, alias := range command.Alias {
			CommandMap[alias] = command.CommandFunc
		}
	}

	s.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		command := strings.Split(m.Content, " ")[0]
		if command, ok := CommandMap[command]; ok {
			command.(func(s *discordgo.Session, m *discordgo.MessageCreate, args []string))(s, m, strings.Split(m.Content, " ")[1:])
		}
	})
}
