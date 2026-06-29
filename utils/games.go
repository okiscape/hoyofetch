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

func (hc *HoyoClient) fetchByGame(gameID Game) (*GameRecordCard, error) {
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

func (hc *HoyoClient) FetchZZZ() (*GameRecordCard, error) {
	return hc.fetchByGame(GameZZZ)
}

func (hc *HoyoClient) FetchHSR() (*GameRecordCard, error) {
	return hc.fetchByGame(GameHSR)
}

func (hc *HoyoClient) FetchGI() (*GameRecordCard, error) {
	return hc.fetchByGame(GameGI)
}

func (hc *HoyoClient) FetchTOT() (*GameRecordCard, error) {
	return hc.fetchByGame(GameTOT)
}

func (hc *HoyoClient) FetchHI3RD() (*GameRecordCard, error) {
	return hc.fetchByGame(GameHI3RD)
}

func (hc *HoyoClient) FetchPP() (*GameRecordCard, error) {
	return hc.fetchByGame(GamePP)
}

func (hc *HoyoClient) FetchHNA() (*GameRecordCard, error) {
	return hc.fetchByGame(GameHNA)
}
