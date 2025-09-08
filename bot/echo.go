package bot

import (
	"ai-girlfriend-gin/models"
	"time"
)

type EchoBot struct{}

func (b EchoBot) GetName() string {
	return "EchoBot"
}

func (b EchoBot) Respond(msg models.Message) models.Message {
	return models.Message{
		From: b.GetName(),
		Text: "You said: " + msg.Text,
		Time: time.Now(),
	}
}