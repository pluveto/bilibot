package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

var botConfig BotConfig

func main() {
	fmt.Println("=== Bilibot v1.0 by Pluvet ===")
	botConfigFile, err := os.Open("bot_config.json")
	if err != nil {
		fmt.Printf("Error: Unable to open bot config file. (%s)\n", err.Error())
		os.Exit(-1)
	}
	bytes, err := ioutil.ReadAll(botConfigFile)
	if err != nil {
		fmt.Printf("Error: Unable to read bot config file. (%s)\n", err.Error())
		os.Exit(-1)
	}
	defer botConfigFile.Close()
	err = json.Unmarshal(bytes, &botConfig)
	if err != nil {
		fmt.Printf("Error: Unable to unmarshal bot config file. (%s)\n", err.Error())
		os.Exit(-1)
	}
	fmt.Printf("jct: %s\n", botConfig.Jct)

	srv := MessageService{Config: botConfig}

	t := time.Now()
	err = srv.SendText(Receiver{ID: 3641797, Type: 1}, "这是一条测试消息。当前时间："+t.Format("2006-01-02 15:04:05"))

	if err != nil {
		fmt.Printf("Unable to send message. (%s)\n", err.Error())
	}

}
