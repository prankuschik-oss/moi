package configs

import (
	_ "errors"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/joho/godotenv"
	"github.com/nicitapa/firstProgect/internal/models"
	"os"
)

var AppSettings models.Config

func ReadSettings() error {
	if err := godotenv.Load(".env"); err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}

	configFile, err := os.Open("internal/configs/configs.json")
	if err != nil {
		return fmt.Errorf("error while opening configs file: %w", err)
	}
	defer configFile.Close()

	if err = json.NewDecoder(configFile).Decode(&AppSettings); err != nil {
		return fmt.Errorf("error while parsing configs file: %w", err)
	}

	return nil
}
