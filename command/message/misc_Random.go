package messageCommand

import (
	"WRwolf_bot-Go/util"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func (m *commandModule) Random() {
	m.Name = "random"
	m.Alias = []string{"선택", "골라", "랜덤"}
	m.Category = "기타"
	m.Description = "쉼표로 구분되어 주어진 선택지를 무작위로 선택합니다."
	m.Usage = "random [선택지1, 선택지2, 선택지3, ...]"
	m.Example = "random 사과, 배, 바나나, 포도"
	m.IsRequireAdminPermission = false
	m.CommandFunc = func(session *discordgo.Session, message *discordgo.MessageCreate, args []string) {
		if len(args) == 0 {
			util.HandleCommandError(session, message, nil, "선택지를 입력해주세요.")
			return
		}
		choices := strings.Split(strings.Join(args, " "), ",")
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		choice := strings.TrimSpace(choices[r.Intn(len(choices))])

		session.ChannelMessageSendEmbedReply(message.ChannelID, &discordgo.MessageEmbed{
			Title:       "🎲 선택 결과",
			Description: fmt.Sprintf("**%s**", choice),
			Footer: &discordgo.MessageEmbedFooter{
				Text: "Developed by. Wonny._.lee",
			},
		}, message.Reference())
	}
}
