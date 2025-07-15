package config

import (
	"encoding/json"
	"os"
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
	return homeDir + "/" + configFileName, nil
}

func write(cfg Config) error {
	configPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	jsonBytes, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	configFile, err := os.OpenFile(configPath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer configFile.Close()

	_, err = configFile.Write(jsonBytes)
	if err != nil {
		return err
	}

	return nil
}
