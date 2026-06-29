package utils

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	hoyoBaseURL    = "https://bbs-api-os.hoyolab.com"
	hoyoAppVersion = "2.50.0"
	hoyoClientType = "5"
	hoyoUserAgent  = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"
	hoyoReferer    = "https://www.hoyolab.com/"
)

type HoyoClient struct {
	auth AccountAuth
	c    *http.Client
}

type GameRecordCardAchievement struct {
	Name  string
	Type  int
	Value string
}

type GameRecordCard struct {
	GameId       int
	GameRoleId   int
	Nickname     string
	Level        int
	Achievements []GameRecordCardAchievement
	Region       string // regionname
}

func NewHoyoClient(auth AccountAuth) *HoyoClient {
	return &HoyoClient{
		auth: auth,
		c:    &http.Client{Timeout: 10 * time.Second},
	}
}

func (hc *HoyoClient) FetchJSON(path string, params map[string]string) ([]byte, error) {
	u, err := url.Parse(hoyoBaseURL + path)
	if err != nil {
		return nil, fmt.Errorf("invalid path: %w", err)
	}

	q := u.Query()
	for k, v := range params {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("request creation failed: %w", err)
	}

	req.Header.Set("Cookie", fmt.Sprintf("ltuid_v2=%s; ltoken_v2=%s; ltmid_v2=%s", hc.auth.LtuidV2, hc.auth.LtokenV2, hc.auth.LtmidV2))
	req.Header.Set("x-rpc-app_version", hoyoAppVersion)
	req.Header.Set("x-rpc-client_type", hoyoClientType)
	req.Header.Set("x-rpc-language", "en-us")
	req.Header.Set("User-Agent", hoyoUserAgent)
	req.Header.Set("Referer", hoyoReferer)
	req.Header.Set("Accept", "application/json, text/plain, */*")

	resp, err := hc.c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("API request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("response read failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned %s: %s", resp.Status, string(body))
	}

	return body, nil
}
