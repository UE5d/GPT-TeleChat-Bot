
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