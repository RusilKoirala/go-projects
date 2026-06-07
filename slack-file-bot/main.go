package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "your-bot-token")
	os.Setenv("SLACK_APP_TOKEN", "your-app-token")
	os.Setenv("CHANNEL_ID", "your-channel-id")
	token := os.Getenv("SLACK_BOT_TOKEN")
	if token == "" {
		log.Fatal("SLACK_BOT_TOKEN environment variable not set")
	}
	channelID := os.Getenv("CHANNEL_ID")
	if channelID == "" {
		log.Fatal("CHANNEL_ID environment variable not set")
	}

	api := slack.New(token)
	fileArr := []string{"file-name"}

	for i := 0; i < len(fileArr); i++ {
		content, err := ioutil.ReadFile(fileArr[i])
		if err != nil {
			log.Println(err)
			return
		}
		params := slack.UploadFileParameters{
			Channel:  channelID,
			Filename: fileArr[i],
			FileSize: len(content),
			Reader:   bytes.NewReader(content),
		}
		file, err := api.UploadFile(params)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Printf("Name: %s, URL : %s\n", file.Title, file.ID)
	}
}
