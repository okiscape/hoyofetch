package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type AccountConfig struct {
	Name string      `json:"name,omitempty"`
	Auth AccountAuth `json:"auth"`
}

type AccountAuth struct {
	LtuidV2  string `json:"ltuid_v2,omitempty"`
	LtokenV2 string `json:"ltoken_v2,omitempty"`
	LtmidV2  string `json:"ltmid_v2,omitempty"`
}

type ConfigModule struct {
	Type        string `json:"type"`
	Format      string `json:"format"`
	DisplayArgs string `json:"display"`
}

type Config struct {
	Accounts []AccountConfig `json:"accounts"`
	Modules  []ConfigModule  `json:"modules"`
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
				},
			},
			Modules: []ConfigModule{
				{
					Type:        "game",
					Format:      "hoyofetch - {}",
					DisplayArgs: "lower",
				},
				{
					Type:   "username",
					Format: "  username   {}",
				},
				{
					Type:   "userId",
					Format: "  user id    {}",
				},
				{
					Type:   "level",
					Format: "  level      {}",
				},
				{
					Type:   "server",
					Format: "  server     {}",
				},
				{
					Type:   "text",
					Format: "  achievements",
				},
				{
					Type:   "achievements",
					Format: "    {name}: {value}",
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

func applyDisplayArgs(s, args string) string {
	if args == "" {
		return s
	}
	for _, arg := range strings.FieldsFunc(args, func(r rune) bool {
		return r == ',' || r == ' ' || r == '|'
	}) {
		switch strings.ToLower(strings.TrimSpace(arg)) {
		case "lower":
			s = strings.ToLower(s)
		case "upper":
			s = strings.ToUpper(s)
		case "capitalize":
			s = strings.Title(s)
		case "trim":
			s = strings.TrimSpace(s)
		}
	}
	return s
}

func ParseModules(modules []ConfigModule, game *GameRecordCard) {
	for _, module := range modules {
		var line string
		switch module.Type {
		case "text":
			line = module.Format
		case "username":
			line = strings.ReplaceAll(module.Format, "{}", game.Nickname)
		case "level":
			line = strings.ReplaceAll(module.Format, "{}", strconv.Itoa(game.Level))
		case "achievements":
			for _, a := range game.Achievements {
				line := module.Format
				line = strings.Replace(line, "{name}", a.Name, 1)
				line = strings.Replace(line, "{value}", a.Value, 1)
				fmt.Println(applyDisplayArgs(line, module.DisplayArgs))
			}
			continue
		case "server":
			line = strings.ReplaceAll(module.Format, "{}", game.Region)
		case "game":
			line = strings.ReplaceAll(module.Format, "{}", Game(game.GameId).String())
		case "userId":
			line = strings.ReplaceAll(module.Format, "{}", game.RoleID)
		default:
			continue
		}
		fmt.Println(applyDisplayArgs(line, module.DisplayArgs))
	}
}
