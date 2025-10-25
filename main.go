package main

import (
	"ai-girlfriend-gin/bot"
	"ai-girlfriend-gin/models"
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// next: implement other bots

func generateMessage(text string) models.Message {
	message := models.Message{
		From: "User",
		Text: text,
		Time: time.Now(),
	}

	return message
}

func incomingChan(ch chan <- models.Message, bot bot.Responder){
	reader := bufio.NewReader(os.Stdin)

	for {
		var message models.Message
		text, err := reader.ReadString('\n')
		if (err != nil){
			fmt.Println("Error occured: " + err.Error())
			return
		}
		text = strings.TrimSpace(text)

		if (text == "quit") {
			close(ch)
			return
		}

		message = generateMessage(text)
		ch <- bot.Respond(message)
	}
}

func outgoingChan(ch <- chan models.Message){
	for msg := range ch {
		fmt.Println(msg.From + ": " + msg.Text)
	}
}

func main() {
	ch := make(chan models.Message)
	bot := bot.EchoBot{}

	go incomingChan(ch, bot)
	outgoingChan(ch)
}
