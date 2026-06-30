package utils

import (
	"encoding/json"
	"fmt"
)

func (hc *HoyoClient) FetchGameRecordCard() ([]byte, error) {
	return hc.FetchJSON(`/game_record/card/wapi/getGameRecordCard?uid=`+hc.auth.LtuidV2, nil)
}

func FindGameRole(data []byte, gameID Game) (*GameRecordCard, error) {
	var resp struct {
		Retcode int    `json:"retcode"`
		Message string `json:"message"`
		Data    struct {
			List []GameRecordCard `json:"list"`
		} `json:"data"`
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, fmt.Errorf("failed to parse card: %w", err)
	}
	if resp.Retcode != 0 {
		return nil, fmt.Errorf("API error (retcode %d): %s", resp.Retcode, resp.Message)
	}
	for _, role := range resp.Data.List {
		if Game(role.GameId) == gameID {
			return &role, nil
		}
	}
	return nil, fmt.Errorf("game %d not found on this account", gameID)
}

func (hc *HoyoClient) FetchByGame(gameID Game) (*GameRecordCard, error) {
	card, err := hc.FetchGameRecordCard()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch game card: %w", err)
	}
	gamecard, err := FindGameRole(card, gameID)
	if err != nil {
		return nil, err
	}
	return gamecard, nil
}
