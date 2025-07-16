package config

import (
	"encoding/json"
	"os"
	"path"
)

const configFileName string = ".gatorconfig.json"

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	configPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	configBytes, err := os.ReadFile(configPath)
	if err != nil {
		return Config{}, err
	}

	config := Config{}
	err = json.Unmarshal(configBytes, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

func (cfg *Config) SetUser(name string) error {
	cfg.CurrentUserName = name
	err := write(*cfg)
	if err != nil {
		return err
	}
	return nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return path.Join(homeDir, configFileName), nil
}

func write(cfg Config) error {
	configPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	configFile, err := os.Create(configPath)
	if err != nil {
		return err
	}
	defer configFile.Close()

	encoder := json.NewEncoder(configFile)
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}

	return nil
}
