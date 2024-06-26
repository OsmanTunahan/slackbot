package handlers

import (
	"context"
	"fmt"
	"strings"
	"net/http"
	"encoding/json"
	"bytes"
	"io"
	"os"
	"time"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
	"slackbot/config"
	"slackbot/models"
)

func HandleEvents(ctx context.Context, client *socketmode.Client) {
	for {
		select {
		case <-ctx.Done():
			return
		case event := <-client.Events:
			switch event.Type {
			case socketmode.EventTypeSlashCommand:
				command, ok := event.Data.(slack.SlashCommand)
				if !ok {
					continue
				}
				client.Ack(*event.Request)
				HandleCommand(client, command)
			}
		}
	}
}

func HandleCommand(client *socketmode.Client, command slack.SlashCommand) {
	var (
		err          error
		slackRequest *models.SlackRequest
	)

	switch command.Command {
	case models.EmailsCommand:
		slackRequest, err = ProcessEmailCommand(command)
	case models.SmsCommand:
		slackRequest, err = ProcessSmsCommand(command)
	case models.UrlCommand:
		slackRequest, err = ProcessUrlCommand(command)
	default:
		slackRequest = &models.SlackRequest{
			StatusCode: http.StatusNotFound,
			Content:    "command not found",
		}
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "error processing command: %s\n", err)
	}

	if err := SendSlackResponse(client, slackRequest); err != nil {
		fmt.Fprintf(os.Stderr, "error sending slack response: %s\n", err)
	}
}

func ProcessEmailCommand(command slack.SlashCommand) (*models.SlackRequest, error) {
	params := strings.Fields(command.Text)
	if len(params) < 4 {
		return &models.SlackRequest{StatusCode: http.StatusBadRequest, Content: config.ErrNotEnoughArgs.Error()}, config.ErrNotEnoughArgs
	}

	emailUrl := fmt.Sprintf("%s/sendgrid-email/sample/emails", config.UrlBase)
	payload := models.Email{
		From:    params[0],
		To:      params[1],
		Subject: params[2],
		Content: strings.Join(params[3:], " "),
	}

	return MakeHTTPRequest(emailUrl, payload)
}

func ProcessSmsCommand(command slack.SlashCommand) (*models.SlackRequest, error) {
	params := strings.Fields(command.Text)
	if len(params) < 3 {
		return &models.SlackRequest{StatusCode: http.StatusBadRequest, Content: config.ErrNotEnoughArgs.Error()}, config.ErrNotEnoughArgs
	}

	smsUrl := fmt.Sprintf("%s/twilio-sms/sample/sms", config.UrlBase)
	payload := models.Sms{
		From:    params[0],
		Number:  params[1],
		Message: strings.Join(params[2:], " "),
	}

	return MakeHTTPRequest(smsUrl, payload)
}

func ProcessUrlCommand(command slack.SlashCommand) (*models.SlackRequest, error) {
	params := strings.Fields(command.Text)
	if len(params) < 3 {
		return &models.SlackRequest{StatusCode: http.StatusBadRequest, Content: config.ErrNotEnoughArgs.Error()}, config.ErrNotEnoughArgs
	}

	preUrl := fmt.Sprintf("%s/presigned-url/presign/url", config.UrlBase)
	payload := models.Url{
		Filename: params[0],
		Type:     params[1],
		Duration: params[2],
	}

	return MakeHTTPRequest(preUrl, payload)
}

func MakeHTTPRequest(url string, payload interface{}) (*models.SlackRequest, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return &models.SlackRequest{StatusCode: http.StatusInternalServerError}, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return &models.SlackRequest{StatusCode: http.StatusInternalServerError}, err
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return &models.SlackRequest{StatusCode: http.StatusInternalServerError}, err
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return &models.SlackRequest{StatusCode: http.StatusInternalServerError}, err
	}

	return &models.SlackRequest{StatusCode: res.StatusCode, Content: string(resBody)}, nil
}

func SendSlackResponse(client *socketmode.Client, req *models.SlackRequest) error {
	api := client.Client
	attachment := slack.Attachment{
		Color: "#0069ff",
		Fields: []slack.AttachmentField{
			{
				Title: fmt.Sprintf("Response: %d", req.StatusCode),
				Value: req.Content,
			},
		},
		Footer: "OsmanTunahan SlackBot | " + time.Now().Format("01-02-2006 3:4:5 MST"),
	}
	_, _, err := api.PostMessage(
		config.ChannelID,
		slack.MsgOptionAttachments(attachment),
		slack.MsgOptionAsUser(true),
	)
	return err
}