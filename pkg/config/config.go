package config

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	LogLevel    string
	LogFormat   string
	DatabaseUrl string
	ServerAddr  string
)

func init() {
	if LogLevel = os.Getenv("LOG_LEVEL"); LogLevel != "" {
		level, err := zerolog.ParseLevel(LogLevel)
		if err != nil {
			panic(err)
		}

		zerolog.SetGlobalLevel(level)
	}

	if LogFormat = os.Getenv("LOG_FORMAT"); LogFormat != "" {
		switch LogFormat {
		case "json":
			log.Logger = zerolog.New(os.Stderr)
		case "text":
			log.Logger = zerolog.New(zerolog.NewConsoleWriter())
		default:
			panic(fmt.Errorf("invalid log format %s: must be 'json' or 'text'", LogFormat))
		}

		log.Logger = log.Logger.With().Timestamp().Logger()
	}

	zerolog.DefaultContextLogger = &log.Logger

	if DatabaseUrl = os.Getenv("DATABASE_URL"); DatabaseUrl == "" {
		DatabaseUrl = "file:local.db?mode=memory"
	}

	if ServerAddr = os.Getenv("SERVER_ADDR"); ServerAddr == "" {
		ServerAddr = ":8000"
	}
}
