package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"okiscape/hoyofetch/utils"
)

func main() {
	args := os.Args[1:]
	defaultOutput := `hoyofetch - fetch for your hoyoverse games account

hoyofetch get [zzz|hsr|gi|hi3rd|tot|hna|pp] - get fetch for your account in hoyo game

Sources: https://github.com/okiscape/hoyofetch`

	if len(args) == 0 {
		fmt.Println(defaultOutput)
		return
	}

	showVersion := flag.Bool("version", false, "Display tool version")
	shortShowVersion := flag.Bool("v", false, "Display tool version")
	flag.Parse()

	if *showVersion || *shortShowVersion {
		utils.PrintVersion()
		os.Exit(0)
	}

	switch args[0] {
	case "get":
		runGet(args[1:])
	default:
		fmt.Println(defaultOutput)
	}
}

func runGet(argv []string) {
	raw := false
	gameArg := ""
	for _, a := range argv {
		if a == "--raw" {
			raw = true
		} else if !strings.HasPrefix(a, "-") {
			gameArg = a
		}
	}

	if gameArg == "" {
		fmt.Println("Usage: hoyofetch get [--raw] [zzz|hsr|gi|hi3rd|tot|hna|pp]")
		return
	}

	game, ok := utils.GameAbbrs[gameArg]
	if !ok {
		fmt.Printf("Unknown game: %s\n", gameArg)
		fmt.Println("Available games: zzz, hsr, gi, tot, hi3rd, pp, hna")
		return
	}

	cfg, err := utils.LoadConfig()
	if err != nil {
		fmt.Printf("Config error: %v\n", err)
		return
	}

	if len(cfg.Accounts) == 0 {
		fmt.Println("No accounts in config")
		return
	}

	acct := cfg.Accounts[0]
	client := utils.NewHoyoClient(acct.Auth)

	if raw {
		cardJSON, err := client.FetchGameRecordCard()
		if err != nil {
			fmt.Printf("Fetch error: %v\n", err)
			return
		}
		fmt.Println(string(cardJSON))
		return
	}

	var card *utils.GameRecordCard
	switch game {
	case utils.GameZZZ:
		card, err = client.FetchZZZ()
	case utils.GameHSR:
		card, err = client.FetchHSR()
	case utils.GameGI:
		card, err = client.FetchGI()
	case utils.GameTOT:
		card, err = client.FetchTOT()
	case utils.GameHI3RD:
		card, err = client.FetchHI3RD()
	case utils.GamePP:
		card, err = client.FetchPP()
	case utils.GameHNA:
		card, err = client.FetchHNA()
	}
	if err != nil {
		fmt.Printf("Fetch error: %v\n", err)
		return
	}

	fmt.Printf("Account: %s (Lv.%d, %s)\n\n", card.Nickname, card.Level, card.Region)
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(card)
}
