package config

import (
	"encoding/json"
	"fmt"
	"os"
)

var (
	Token     string
	BotPrefix string

	config *Config
)

type Config struct {
	Token     string `json:"token"`
	BotPrefix string `json:"botPrefix`
}

func ReadConfig() error {
	fmt.Println("Reading config.json")
	file, err := os.ReadFile("./config.json")

	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println("Error unmarshalling config.json")
		return err
	}

	Token = config.Token
	BotPrefix = config.BotPrefix

	return nil
}
