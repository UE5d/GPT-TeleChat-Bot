
package main

import (
	"context"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/joho/godotenv"
	"github.com/rakyll/openai-go"
	"github.com/rakyll/openai-go/chat"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	retainHistory bool
	promptName    = "prompt.txt"
)

func main() {
	// setup logger
	log.Logger = log.With().Caller().Logger()
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// Environment files
	err := godotenv.Load()
	if err != nil {
		log.Debug().Msg(err.Error())
	}

	retainHistory = os.Getenv("RETAIN_HISTORY") == "true"

	if err := ConnectDB(); err != nil {
		log.Fatal().Msg(err.Error())
	}

	// start server
	StartServer()
}

// StartServer starts the telegram server
func StartServer() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
	}

	b, err := bot.New(os.Getenv("TELEGRAM_API_KEY"), opts...)
	if err != nil {
		panic(err)
	}

	log.Debug().Msg("Telegram bot started!")
	b.Start(ctx)
}

// SendToChatGPT send a message to chatgpt
func SendToChatGPT(chatId, textMsg string) []*chat.Choice {
	var (
		ctx = context.Background()
		s   = openai.NewSession(os.Getenv("OPENAI_TOKEN"))

		// messages that will be sent to chatgpt
		gptMsgs = make([]*chat.Message, 0)
	)

	// check if the user has a previous conversation
	prevMessages, err := FindMessages(chatId)
	if err != nil {
		log.Err(err)
	}

	// get the systems prompt
	prmptB, _ := os.ReadFile(promptName)

	// add system prompt if user is initially starting out the conversation
	if len(prevMessages) == 0 {
		// create & add the systems prompt first
		log.Debug().Msg("added system prompt because its a first time user")
		gptMsgs = append(gptMsgs, &chat.Message{
			Role:    "user", // "system"
			Content: string(prmptB),
		})

	} else {
		// if we're retaining history
		if retainHistory {
			// add the whole previous users conversation + current text message and send to chatgpt
			// this may include the previous prompt from the conversation
			for _, prevMsg := range prevMessages {
				gptMsgs = append(gptMsgs, &chat.Message{
					Role:    prevMsg.Role,
					Content: prevMsg.Content,
				})
			}
		} else {
			// add only the system prompt to gpt
			gptMsgs = append(gptMsgs, &chat.Message{
				Role:    "user", // "system"
				Content: string(prmptB),
			})
		}
	}

	// add this current message
	gptMsgs = append(gptMsgs, &chat.Message{
		Role:    "user",
		Content: textMsg,
	})

	// process request
	client := chat.NewClient(s, "gpt-3.5-turbo-0301")
	resp, err := client.CreateCompletion(ctx, &chat.CreateCompletionParams{
		Messages: gptMsgs,
	})
	if err != nil {
		log.Error().Msgf("Failed to complete: %v", err)
		return nil
	}

	// save the new prompt + current text to DB
	if len(prevMessages) == 0 {
		for _, gptMsg := range gptMsgs {
			_, err := CreateMessage(Message{
				ChatID:  chatId,
				Content: gptMsg.Content,
				Role:    gptMsg.Role,

				// metrics for this single chat session