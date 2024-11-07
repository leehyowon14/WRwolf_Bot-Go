package util

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func HandleError(err error, errorString string) {
	if err != nil {
		log.Panicln(errorString+"(", err, ")")
	}
}

func HandleCommandError(session *discordgo.Session, message *discordgo.MessageCreate, err error, errorString string, noPanic ...bool) {
	description := errorString
	if err != nil {
		description += "(" + err.Error() + ")"
	}
	session.ChannelMessageSendEmbedReply(message.ChannelID, &discordgo.MessageEmbed{
		Color:       0xED4245,
		Title:       "Error",
		Description: description,
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Developed by. Wonny._.lee",
		},
	}, message.Reference())
	if len(noPanic) != 0 && !noPanic[0] {
		log.Panicln(description)
	}
}
