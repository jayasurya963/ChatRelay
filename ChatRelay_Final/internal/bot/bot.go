// Handles Slack event listening and relay logic
package bot

import (
	"chatrelay/internal/backend"
	"chatrelay/internal/config"
	"context"
	"fmt"
	"github.com/slack-go/slack/socketmode"
	"github.com/slack-go/slack"
	"log"
	"strings"
)

func Start(ctx context.Context, cfg *config.Config, backendClient *backend.Client) error {
	api := slack.New(
		cfg.SlackBotToken,
		slack.OptionAppLevelToken(cfg.SlackAppToken),
	)
	socketClient := socketmode.New(api)

	go func() {
		for evt := range socketClient.Events {
			switch evt.Type {
			case socketmode.EventTypeInteractive:
				continue
			case socketmode.EventTypeEventsAPI:
				eventsAPIEvent, ok := evt.Data.(slack.EventsAPIEvent)
				if !ok {
					continue
				}
				if eventsAPIEvent.Type == slack.EventsAPITypeAppMention {
					event := eventsAPIEvent.InnerEvent.Data.(*slack.AppMentionEvent)
					go handleMessage(ctx, api, backendClient, event)
				}
				socketClient.Ack(*evt.Request)
			}
		}
	}()

	return socketClient.RunContext(ctx)
}

func handleMessage(ctx context.Context, api *slack.Client, backendClient *backend.Client, event *slack.AppMentionEvent) {
	channel := event.Channel
	user := event.User
	text := strings.TrimSpace(strings.Replace(event.Text, fmt.Sprintf("<@%s>", event.BotID), "", 1))

	msg, _, _ := api.PostMessage(channel, slack.MsgOptionText("Processing your query...", false))

	err := backendClient.StreamResponse(ctx, backend.ChatRequest{
		UserID: user,
		Query:  text,
	}, func(part string, done bool) {
		if done {
			return
		}
		api.UpdateMessage(channel, msg.Timestamp, slack.MsgOptionText(part, false))
	})

	if err != nil {
		api.PostMessage(channel, slack.MsgOptionText("❌ Error: "+err.Error(), false))
		log.Printf("error streaming response: %v", err)
	}
}
