package main

import (
	"flag"
	"fmt"
	"os"

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
	if len(argv) == 0 {
		fmt.Println("Usage: hoyofetch get [zzz|hsr|gi|hi3rd|tot|hna|pp]")
		return
	}

	game, ok := utils.GameAbbrs[argv[0]]
	if !ok {
		fmt.Printf("Unknown game: %s\n", argv[0])
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

	card, err := client.FetchGameRecordCard()
	if err != nil {
		fmt.Printf("Failed to get game card: %v\n", err)
		return
	}

	role, err := utils.FindGameRole(card, game)
	if err != nil {
		fmt.Printf("Game not found on this account: %v\n", err)
		return
	}

	fmt.Printf("Account: %s (Lv.%d, %s)\n", role.Nickname, role.Level, role.Region)

	var data []byte
	switch game {
	case utils.GameZZZ:
		data, err = client.FetchZZZ(role.RoleID, role.Region)
	case utils.GameHSR:
		data, err = client.FetchHSR(role.RoleID, role.Region)
	case utils.GameGI:
		data, err = client.FetchGI(role.RoleID, role.Region)
	case utils.GameTOT:
		data, err = client.FetchTOT(role.RoleID, role.Region)
	case utils.GameHI3RD:
		data, err = client.FetchHI3RD(role.RoleID, role.Region)
	case utils.GamePP, utils.GameHNA:
		fmt.Println("Not yet available on Hoyolab API")
		return
	}
	if err != nil {
		fmt.Printf("Fetch error: %v\n", err)
		return
	}

	fmt.Println(string(data))
}
