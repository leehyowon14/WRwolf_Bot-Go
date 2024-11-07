package guildEvent

import (
	"WRwolf_bot-Go/util"
	"fmt"
	"os"
	"reflect"
	"strings"
)

// 	"ChannelCreate":                       *discordgo.ChannelCreate,
// 	"ChannelUpdate":                       *discordgo.ChannelUpdate,
// 	"ChannelDelete":                       *discordgo.ChannelDelete,
// 	"ChannelPinsUpdate":                   *discordgo.ChannelPinsUpdate,
// 	"ThreadCreate":                        *discordgo.ThreadCreate,
// 	"ThreadUpdate":                        *discordgo.ThreadUpdate,
// 	"ThreadDelete":                        *discordgo.ThreadDelete,
// 	"ThreadListSync":                      *discordgo.ThreadListSync,
// 	"ThreadMemberUpdate":                  *discordgo.ThreadMemberUpdate,
// 	"ThreadMembersUpdate":                 *discordgo.ThreadMembersUpdate,
// 	"GuildCreate":                         *discordgo.GuildCreate,
// 	"GuildUpdate":                         *discordgo.GuildUpdate,
// 	"GuildDelete":                         *discordgo.GuildDelete,
// 	"GuildBanAdd":                         *discordgo.GuildBanAdd,
// 	"GuildBanRemove":                      *discordgo.GuildBanRemove,
// 	"GuildMemberAdd":                      *discordgo.GuildMemberAdd,
// 	"GuildMemberUpdate":                   *discordgo.GuildMemberUpdate,
// 	"GuildMemberRemove":                   *discordgo.GuildMemberRemove,
// 	"GuildRoleCreate":                     *discordgo.GuildRoleCreate,
// 	"GuildRoleUpdate":                     *discordgo.GuildRoleUpdate,
// 	"GuildRoleDelete":                     *discordgo.GuildRoleDelete,
// 	"GuildEmojisUpdate":                   *discordgo.GuildEmojisUpdate,
// 	"GuildMembersChunk":                   *discordgo.GuildMembersChunk,
// 	"GuildIntegrationsUpdate":             *discordgo.GuildIntegrationsUpdate,
// 	"StageInstanceEventCreate":            *discordgo.StageInstanceEventCreate,
// 	"StageInstanceEventUpdate":            *discordgo.StageInstanceEventUpdate,
// 	"StageInstanceEventDelete":            *discordgo.StageInstanceEventDelete,
// 	"GuildScheduledEventCreate":           *discordgo.GuildScheduledEventCreate,
// 	"GuildScheduledEventUpdate":           *discordgo.GuildScheduledEventUpdate,
// 	"GuildScheduledEventDelete":           *discordgo.GuildScheduledEventDelete,
// 	"GuildScheduledEventUserAdd":          *discordgo.GuildScheduledEventUserAdd,
// 	"GuildScheduledEventUserRemove":       *discordgo.GuildScheduledEventUserRemove,
// 	"MessageCreate":                       *discordgo.MessageCreate,
// 	"MessageUpdate":                       *discordgo.MessageUpdate,
// 	"MessageDelete":                       *discordgo.MessageDelete,
// 	"MessageReactionAdd":                  *discordgo.MessageReactionAdd,
// 	"MessageReactionRemove":               *discordgo.MessageReactionRemove,
// 	"MessageReactionRemoveAll":            *discordgo.MessageReactionRemoveAll,
// 	"PresencesReplace":                    *discordgo.PresencesReplace,
// 	"PresenceUpdate":                      *discordgo.PresenceUpdate,
// 	"Resumed":                             *discordgo.Resumed,
// 	"TypingStart":                         *discordgo.TypingStart,
// 	"UserUpdate":                          *discordgo.UserUpdate,
// 	"VoiceServerUpdate":                   *discordgo.VoiceServerUpdate,
// 	"VoiceStateUpdate":                    *discordgo.VoiceStateUpdate,
// 	"MessageDeleteBulk":                   *discordgo.MessageDeleteBulk,
// 	"WebhooksUpdate":                      *discordgo.WebhooksUpdate,
// 	"InteractionCreate":                   *discordgo.InteractionCreate,
// 	"InviteCreate":                        *discordgo.InviteCreate,
// 	"InviteDelete":                        *discordgo.InviteDelete,
// 	"ApplicationCommandPermissionsUpdate": *discordgo.ApplicationCommandPermissionsUpdate,
// 	"AutoModerationRuleCreate":            *discordgo.AutoModerationRuleCreate,
// 	"AutoModerationRuleUpdate":            *discordgo.AutoModerationRuleUpdate,
// 	"AutoModerationRuleDelete":            *discordgo.AutoModerationRuleDelete,
// 	"AutoModerationActionExecution":       *discordgo.AutoModerationActionExecution,
// 	"GuildAuditLogEntryCreate":            *discordgo.GuildAuditLogEntryCreate,

type event struct {
	EventName string
	EventFunc interface{}
}

var GuildEventArr = make([]*event, 0)

func init() {
	files, err := os.ReadDir("events/guild")
	util.HandleError(err, "Error reading guild events directory")

	reflectType := reflect.TypeOf(&event{})
	for _, file := range files {
		methodName := strings.Split(file.Name(), ".")[0]
		method, isExist := reflectType.MethodByName(methodName)
		if isExist {
			eventStruct := &event{}
			method.Func.Interface().(func(eventStruct *event))(eventStruct)
			GuildEventArr = append(GuildEventArr, eventStruct)
			fmt.Println(eventStruct.EventName + " is registering")
		}
	}
}
