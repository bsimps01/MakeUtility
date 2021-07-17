package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main() {
	args := os.Args[1:]
	fmt.Println(args)
	godotenv.Load()
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelID, timestamp, err := api.PostMessage(
		"C028YGHJH6U",
		slack.MsgOptionText("*GITHUB UPDATES POSTED!* find the repo at: https://github.com/bsimps01/MakeUtility", false),
	)

	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)

	//jenkinsBuild()

}

func jenkinsBuild() {
	args := os.Args[1:]
	fmt.Println(args)
	godotenv.Load()
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	preText := "*Github Updates posted!*"
	githubURL := "*Build URL:*" + args[0]
	buildResult := "*" + args[1] + "*"
	buildNumber := "*" + args[2] + "*"
	jobName := "*" + args[3] + "*"

	if buildResult == "*SUCCESS*" {
		buildResult = buildResult + " :white_check_mark:"
	} else {
		buildResult = buildResult + " :x:"
	}

	dividerSection1 := slack.NewDividerBlock()
	jenkinsBuildDetails := jobName + " #" + buildNumber + " - " + buildResult + "\n" + githubURL
	preTextField := slack.NewTextBlockObject("mrkdwn", preText+"\n\n", false, false)
	jenkinsBuildDetailsField := slack.NewTextBlockObject("mrkdwn", jenkinsBuildDetails, false, false)

	jenkinsBuildDetailsSection := slack.NewSectionBlock(jenkinsBuildDetailsField, nil, nil)
	preTextSection := slack.NewSectionBlock(preTextField, nil, nil)

	msg := slack.MsgOptionBlocks(
		preTextSection,
		dividerSection1,
		jenkinsBuildDetailsSection,
	)

	_, _, _, err := api.SendMessage(
		"C028YGHJH6U",
		msg,
	)

	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
}
