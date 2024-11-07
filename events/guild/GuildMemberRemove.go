package guildEvent

import (
	"WRwolf_bot-Go/util"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func (e *event) GuildMemberRemove() {
	e.EventName = "GuildMemberRemove"
	e.EventFunc = func(s *discordgo.Session, m *discordgo.GuildMemberRemove) {
		defer func() {
			recover()
		}()
		member, err := s.GuildMember(m.GuildID, m.User.ID)
		util.HandleError(err, "Failed to get member")
		fmt.Println(member.User.Username)
	}
}
