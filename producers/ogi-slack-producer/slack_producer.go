package main

import (
	"encoding/json"
	"fmt"
	"time"

	logger "github.com/OpenChaos/ogi/logger"

	"github.com/gol-gol/golenv"
	"github.com/slack-go/slack"
)

type Notif struct {
	Context string `json:"context"`
	Message string `json:"message"`
}

var (
	SlackApiToken  = golenv.OverrideIfEnv("OGI_SLACK_API_TOKEN", "")
	SlackChannelId = golenv.OverrideIfEnv("OGI_SLACK_CHANNEL_ID", "")
	TZ             = golenv.OverrideIfEnv("OGI_TIMEZONE", "Asia/Kolkata")
)

func Close() {
	return
}

func Produce(msgid string, msg []byte) ([]byte, error) {
	var notif Notif
	if errJson := json.Unmarshal(msg, &notif); errJson != nil {
		return []byte{}, errJson
	}
	client := slack.New(SlackApiToken, slack.OptionDebug(true))
	attachment := slack.Attachment{
		Pretext: notif.Context,
		Text:    notif.Message,
		Color:   "#f0256f",
		Fields: []slack.AttachmentField{
			{
				Title: "Date",
				Value: NowInTimeZone(),
			},
		},
	}
	if _, _, err := client.PostMessage(
		SlackChannelId,
		slack.MsgOptionAttachments(attachment),
	); err != nil {
		return []byte{}, err
	}
	logger.Infof("Slack notifcation sent for msgid[%s].", msgid)
	return []byte{}, nil
}

func NowInTimeZone() string {
	currentTime := time.Now()
	if location, err := time.LoadLocation(TZ); err == nil {
		return currentTime.In(location).String()
	}
	return fmt.Sprintf("%s (TZ: %s)", currentTime.String(), TZ)
}
