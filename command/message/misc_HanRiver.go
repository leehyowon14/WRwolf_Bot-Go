package messageCommand

import (
	"WRwolf_bot-Go/util"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"

	"github.com/bwmarrin/discordgo"
)

var hangangPosition = "탄천"

func (m *commandModule) HanRiver() {
	m.Name = "river"
	m.Alias = []string{"한강", "수온", "온도", "한강수온", "한강온도"}
	m.Category = "기타"
	m.Description = "한강의 수온을 확인합니다."
	m.Usage = "river"
	m.Example = "한강"
	m.IsRequireAdminPermission = false
	m.CommandFunc = func(session *discordgo.Session, message *discordgo.MessageCreate, args []string) {
		res, err := http.Get("https://api.hangang.life")
		defer func() {
			recover()
		}()
		util.HandleCommandError(session, message, err, "한강 데이터 가져오는 중 오류 발생")
		if res.StatusCode != 200 {
			util.HandleCommandError(session, message, nil, "한강 데이터 서버 오류 발생")
		}
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		util.HandleCommandError(session, message, err, "한강 데이터 읽는 중 오류 발생")

		var data map[string]interface{}
		err = json.Unmarshal(body, &data)
		util.HandleCommandError(session, message, err, "한강 데이터 언마샬 중 오류 발생")
		hangangData := data["DATAs"].(map[string]interface{})["DATA"].(map[string]interface{})["HANGANG"].(map[string]interface{})[hangangPosition]
		if hangangData == nil {
			hangangPosition = getTempCheckPosition(session, message, string(body))
			log.Println(hangangPosition)
			hangangData = data["DATAs"].(map[string]interface{})["DATA"].(map[string]interface{})["HANGANG"].(map[string]interface{})[hangangPosition]
			fmt.Println(hangangPosition)
			fmt.Println(hangangData)
			if hangangData == nil {
				util.HandleCommandError(session, message, nil, "한강 데이터 측정 위치 오류 발생")
			}
		}
		hangangDataMap := hangangData.(map[string]interface{})

		session.ChannelMessageSendEmbed(message.ChannelID, &discordgo.MessageEmbed{
			Color:       0x4fe8a3,
			Title:       "한강 수온",
			Description: fmt.Sprintf("측정 위치: %s, pH농도: %.1f", hangangPosition, hangangDataMap["PH"]),
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   ":droplet: " + fmt.Sprintf("%.1f", hangangDataMap["TEMP"].(float64)),
					Value:  "측정 시각: " + hangangDataMap["LAST_UPDATE"].(string),
					Inline: true,
				},
			},
			Footer: &discordgo.MessageEmbedFooter{
				Text: "Developed by. Wonny._.lee",
			},
		})
	}
}

func getTempCheckPosition(session *discordgo.Session, message *discordgo.MessageCreate, body string) string {
	reg, err := regexp.Compile("[가-힣]+")
	util.HandleCommandError(session, message, err, "한강 데이터 읽는 중 오류 발생")
	return reg.FindString(string(body))
}
