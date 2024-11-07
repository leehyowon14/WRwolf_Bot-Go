package messageCommand

import (
	"math/rand"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

func (m *commandModule) Dice() {
	m.Name = "dice"
	m.Alias = []string{"주사위"}
	m.Category = "기타"
	m.Description = "주사위를 굴립니다."
	m.Usage = "dice"
	m.Example = "주사위"
	m.IsRequireAdminPermission = false
	m.CommandFunc = func(session *discordgo.Session, message *discordgo.MessageCreate, args []string) {
		session.ChannelMessageSend(message.ChannelID, "🎲 "+strconv.Itoa(rand.Intn(6)+1))
	}
}
