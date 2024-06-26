package config

import (
	"errors"
	"os"
)

var (
	AuthToken, AppToken, ChannelID, UrlBase string
	ErrNotEnoughArgs                        = errors.New("not enough arguments provided")
)

func Init() {
	AuthToken = os.Getenv("AUTH_TOKEN")
	if AuthToken == "" {
		panic("no authToken provided")
	}
	AppToken = os.Getenv("APP_TOKEN")
	if AppToken == "" {
		panic("no appToken provided")
	}
	ChannelID = os.Getenv("CHANNEL_ID")
	if ChannelID == "" {
		panic("no channelid provided")
	}
	UrlBase = os.Getenv("URL")
	if UrlBase == "" {
		panic("no url provided")
	}
}