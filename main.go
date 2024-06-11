package main

import (
	"fmt"
	"github.com/kerimcanbalkan/discord-bot/bot"
	"github.com/kerimcanbalkan/discord-bot/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err)
		return
	}

	bot.Start()

	<-make(chan struct{})

	return
}
