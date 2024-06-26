package main

import (
	"context"
	"log"
	"golang.org/x/sync/errgroup"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
	"slackbot/config"
	"slackbot/models"
)

func init() {
	config.Init()
}

func main() {
	var eg errgroup.Group
	api := slack.New(config.AuthToken, slack.OptionDebug(true), slack.OptionAppLevelToken(config.AppToken))
	client := socketmode.New(api, socketmode.OptionDebug(true))

	c, cancel := context.WithCancel(context.Background())
	defer cancel()
}