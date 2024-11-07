package main

import (
	handler "WRwolf_bot-Go/handlers"
	"WRwolf_bot-Go/util"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func run() {
	discordSession, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))
	util.HandleError(err, "세션 생성 중 에러 발생")

	//discordApplication := &discordgo.Application{}
	util.HandleError(err, "애플리케이션 정보 가져오기 중 에러 발생")

	//Handler 등록
	handler.RegisterEventHandlers(discordSession)
	handler.RegisterMessageCommandHandlers(discordSession)

	discordSession.Open()
	defer discordSession.Close()

	//인터럽트 처리
	interruptChannel := make(chan os.Signal, 1)
	signal.Notify(interruptChannel, os.Interrupt)
	<-interruptChannel
	log.Println("Bot is shutting down")
}

func main() {
	err := godotenv.Load(".env")
	util.HandleError(err, "환경 변수 로드 중 에러 발생")
	run()
}
