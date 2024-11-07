package messageCommand

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

func (m *commandModule) Ping() {
	m.Name = "ping"
	m.Alias = []string{"핑"}
	m.Category = "기타"
	m.Description = "서버의 핑을 확인합니다."
	m.Usage = "ping"
	m.Example = "핑"
	m.IsRequireAdminPermission = false
	m.CommandFunc = func(session *discordgo.Session, message *discordgo.MessageCreate, args []string) {
		msg, _ := session.ChannelMessageSend(message.ChannelID, "🏓 Ping!")
		session.ChannelMessageEdit(msg.ChannelID, msg.ID, fmt.Sprintf("🏓 Pong! (%dms)", time.Since(msg.Timestamp).Milliseconds()))
	}
}
