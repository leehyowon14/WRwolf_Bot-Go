package guildEvent

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func (e *event) MessageCreate() {
	e.EventName = "MessageCreate"
	e.EventFunc = func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}
		fmt.Println(m.Content)
	}
}
