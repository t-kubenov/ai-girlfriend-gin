package main

import (
	"ai-girlfriend-gin/bot"
	"ai-girlfriend-gin/models"
	"fmt"
	"time"
	// "context"
)

// func inputRoutine(ctx context.Context){
// 	select {
// 		case 
// 	}
// }

func generateMessage(text string) models.Message {
	message := models.Message{
		From: "User",
		Text: text,
		Time: time.Now(),
	}

	return message
}

func incomingChan(ch chan <- models.Message, bot bot.Responder){
	var text string
	var message models.Message
	_, err := fmt.Scanln(&text)

	if (err != nil){
		fmt.Println("Error occured: " + err.Error())
		return
	}

	message = generateMessage(text)
	ch <- bot.Respond(message)
}

func outgoingChan(ch <- chan models.Message){
	fmt.Println(<-ch)
}

func main() {
	ch := make(chan models.Message)
	bot := bot.EchoBot{}

	go incomingChan(ch, bot)
	outgoingChan(ch)
}
