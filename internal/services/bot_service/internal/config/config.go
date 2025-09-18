package config

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/alexey-dobry/tech-support-platform/internal/pkg/logger/zap"
	"github.com/alexey-dobry/tech-support-platform/internal/pkg/validator"
	"github.com/alexey-dobry/tech-support-platform/internal/services/bot_service/internal/bot"
	"github.com/alexey-dobry/tech-support-platform/internal/services/bot_service/internal/session"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Logger  zap.Config     `yaml:"logger" validate:"required"`
	Bot     bot.Config     `yaml:"bot" validate:"required"`
	Session session.Config `yaml:"session" validate:"required"`
}

func MustLoad() Config {
	var cfg Config
	configPath := ParseFlag(cfg)

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("Failed to read config on path(%s): %s", configPath, err)
	}

	if err := validator.V.Struct(&cfg); err != nil {
		log.Fatalf("Failed to validate config: %s", err)
	}

	return cfg
}

func ParseFlag(cfg Config) string {
	configPath := flag.String("config", "./config/config.yaml", "config file path")
	configHelp := flag.Bool("help", false, "show configuration help")

	if *configHelp {
		headerText := "Configuration options:"
		help, err := cleanenv.GetDescription(&cfg, &headerText)
		if err != nil {
			log.Fatalf("error getting configuration description: %s", err)
		}
		fmt.Println(help)
		os.Exit(0)
	}

	return *configPath
}
