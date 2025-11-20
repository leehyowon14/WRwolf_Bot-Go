package messageCommand

import (
	"sort"

	"github.com/bwmarrin/discordgo"
)

func (m *commandModule) Help() {
	m.Name = "help"
	m.Alias = []string{"도움말", "도움", "help"}
	m.Category = "기타"
	m.Description = "명령어 도움말을 확인합니다."
	m.Usage = "help"
	m.Example = "도움말"
	m.IsRequireAdminPermission = false
	m.CommandFunc = func(session *discordgo.Session, message *discordgo.MessageCreate, args []string) {
		session.ChannelMessageSendEmbed(message.ChannelID, generateHelpEmbed())
	}
}

func generateHelpEmbed() *discordgo.MessageEmbed { // WIP
	embed := &discordgo.MessageEmbed{
		Title:       "도움말",
		Description: "명령어 도움말을 확인합니다.",
		Color:       0x00ff00,
		Fields:      []*discordgo.MessageEmbedField{},
	}

	categoryMap := make(map[string][]*discordgo.MessageEmbedField)

	for _, command := range CommandArr {
		field := &discordgo.MessageEmbedField{
			Name:   command.Name,
			Value:  command.Description,
			Inline: true,
		}
		categoryMap[command.Category] = append(categoryMap[command.Category], field)
	}
	var sortedCategories []string
	for category := range categoryMap {
		sortedCategories = append(sortedCategories, category)
	}

	//카테고리 정렬
	sort.Slice(sortedCategories, func(i, j int) bool {
		return sortedCategories[i] < sortedCategories[j]
	})
	for category, fields := range categoryMap {
		sort.Slice(fields, func(i, j int) bool {
			return fields[i].Name < fields[j].Name
		})
		categoryMap[category] = fields
	}

	for _, category := range sortedCategories {
		fields := categoryMap[category]
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:  category,
			Value: "\u200B",
		})
		embed.Fields = append(embed.Fields, fields...)
	}

	for category, fields := range categoryMap {
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:  category,
			Value: "\u200B",
		})
		embed.Fields = append(embed.Fields, fields...)
	}

	for _, command := range CommandArr {
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:   command.Name,
			Value:  command.Description,
			Inline: true,
		})
	}
	return embed
}
