package bot

import "ai-girlfriend-gin/models"

type Responder interface {
	GetName() string
	Respond(msg models.Message) models.Message
}