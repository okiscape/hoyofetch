package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type AccountConfig struct {
	Name  string      `json:"name,omitempty"`
	Auth  AccountAuth `json:"auth"`
	Games []string    `json:"games,omitempty"`
}

type AccountAuth struct {
	LtuidV2  string `json:"ltuid_v2,omitempty"`
	LtokenV2 string `json:"ltoken_v2,omitempty"`
	LtmidV2  string `json:"ltmid_v2,omitempty"`
}

type Config struct {
	Accounts []AccountConfig `json:"accounts"`
}

func getConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".config", "hoyofetch", "config.json"), nil
}

func LoadConfig() (*Config, error) {
	configPath, err := getConfigPath()
	if err != nil {
		return nil, fmt.Errorf("Cannot receive home directory: %w", err)
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		fmt.Println("Config was not found, using default")

		defaultConfig := Config{
			Accounts: []AccountConfig{
				{
					Name: "main",
					Auth: AccountAuth{
						LtuidV2:  "",
						LtokenV2: "",
						LtmidV2:  "",
					},
					Games: []string{"zzz", "gi"},
				},
			},
		}

		err := os.MkdirAll(filepath.Dir(configPath), 0755)
		if err != nil {
			return nil, fmt.Errorf("Cannot create config directory: %w", err)
		}

		file, err := os.Create(configPath)
		if err != nil {
			return nil, fmt.Errorf("Cannot create config file: %w", err)
		}
		defer file.Close()

		encoder := json.NewEncoder(file)
		encoder.SetIndent("", "  ")
		if err := encoder.Encode(&defaultConfig); err != nil {
			return nil, fmt.Errorf("Cannot create config file: %w", err)
		}

		return &defaultConfig, nil
	}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cfg Config
	if err := json.NewDecoder(file).Decode(&cfg); err != nil {
		return nil, fmt.Errorf("Error occured while JSON parsing: %w", err)
	}

	return &cfg, nil
}
