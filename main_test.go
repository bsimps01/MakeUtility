package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func TestAPI(t *testing.T) {
	godotenv.Load()
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	if api != slack.New(os.Getenv("SLACK_BOT_TOKEN")) {
		t.Errorf("Your slackbot was not connecting", api)
	}
}

func TestChannel(b *testing.T) {
	godotenv.Load()
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelID := "C028YGHJH6U"
	b.Errorf("Channel Id was never used", channelID, api)

}

func BenchmarkCalculate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		args := os.Args[1:]
		fmt.Println(args)
	}
}
