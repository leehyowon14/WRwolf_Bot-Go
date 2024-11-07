package messageCommand

import (
	"WRwolf_bot-Go/util"
	"time"

	"github.com/bwmarrin/discordgo"
)

func (m *commandModule) Avatar() {
	m.Name = "avatar"
	m.Alias = []string{"아바타", "프사"}
	m.Category = "기타"
	m.Description = "유저의 프로필 사진을 불러옵니다."
	m.Usage = "avatar [없음|@멘션|유저ID]"
	m.Example = "아바타 @user"
	m.IsRequireAdminPermission = false
	m.CommandFunc = func(session *discordgo.Session, message *discordgo.MessageCreate, args []string) {
		if len(args) == 0 {
			args = append(args, message.Author.ID)
		} else if args[0] != "@me" {
			args[0] = util.GetUserNameFromMention(args[0])
		}
		user, err := session.User(args[0])

		if err != nil {
			session.ChannelMessageSendReply(message.ChannelID, "해당 유저를 찾을 수 없습니다.", message.Reference())
			return
		}
		userAvatarURL := user.AvatarURL("2048")
		session.ChannelMessageSendEmbedReply(message.ChannelID, &discordgo.MessageEmbed{
			Image: &discordgo.MessageEmbedImage{
				URL: userAvatarURL,
			},
			Title: user.Username + "의 프로필 사진",
			Color: 0x00ff00,
			Footer: &discordgo.MessageEmbedFooter{
				Text: "Developed by. Wonny._.lee",
			},
			Timestamp: time.Now().UTC().Format("2006-01-02 15:04:05"),
		}, message.Reference())
	}
}
