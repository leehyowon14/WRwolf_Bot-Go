package clientEvent

import (
	guildEvent "WRwolf_bot-Go/events/guild"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

func Ready(s *discordgo.Session, r *discordgo.Ready) {

	fmt.Println(r.User.Username + " is ready")

	waitAmount := 600
	for {
		time.Sleep(time.Duration(waitAmount) * time.Second)

		if time.Now().Unix()-guildEvent.UserJoinMapLastUpdateTime >= int64(waitAmount) {
			guildEvent.SaveUserJoinMap()
		} else {
			waitAmount = int(600 - (time.Now().Unix() - guildEvent.UserJoinMapLastUpdateTime))
		}
	}
}
