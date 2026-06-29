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

func (hc *HoyoClient) fetchGameData(roleID, server, gamePath string) ([]byte, error) {
	return hc.FetchJSON(gamePath, map[string]string{
		"role_id": roleID,
		"server":  server,
	})
}

func (hc *HoyoClient) FetchZZZ(roleID, server string) ([]byte, error) {
	cardfetch := hc.FetchGameRecordCard()

}

func (hc *HoyoClient) FetchHSR(roleID, server string) ([]byte, error) {
	return hc.fetchGameData(roleID, server, "/event/game_record/hkrpg/api/index")
}

func (hc *HoyoClient) FetchGI(roleID, server string) ([]byte, error) {
	return hc.fetchGameData(roleID, server, "/event/game_record/genshin/api/index")
}

func (hc *HoyoClient) FetchTOT(roleID, server string) ([]byte, error) {
	return hc.fetchGameData(roleID, server, "/event/game_record/tears_of_themis/api/index")
}

func (hc *HoyoClient) FetchHI3RD(roleID, server string) ([]byte, error) {
	return hc.fetchGameData(roleID, server, "/event/game_record/honkai3rd/api/index")
}
