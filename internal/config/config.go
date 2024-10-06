package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUsername string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

func Read() (Config, error) {
	filePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	rawFile, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, err
	}
	var cfg Config
	err = json.Unmarshal(rawFile, &cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, nil
}

func (cfg *Config) SetUser(username string) error {
	filePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	cfg.CurrentUsername = username

	rawFile, err := json.Marshal(&cfg)
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, rawFile, 0644)
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	filePath := home + string(os.PathSeparator) + configFileName
	return filePath, nil
}
