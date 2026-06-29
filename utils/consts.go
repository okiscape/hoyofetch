package utils

type Game int

const (
	GameHI3RD Game = 1
	GameGI    Game = 2
	GameTOT   Game = 4
	GameHSR   Game = 6
	GameZZZ   Game = 8
	GameHNA   Game = 9
	GamePP    Game = 10
)

var GameAbbrs = map[string]Game{
	"zzz":   GameZZZ,
	"hsr":   GameHSR,
	"gi":    GameGI,
	"tot":   GameTOT,
	"hi3rd": GameHI3RD,
	"pp":    GamePP,
	"hna":   GameHNA,
}

func (g Game) String() string {
	switch g {
	case GameZZZ:
		return "Zenless Zone Zero"
	case GameHSR:
		return "Honkai: Star Rail"
	case GameGI:
		return "Genshin Impact"
	case GameTOT:
		return "Tears of Themis"
	case GameHI3RD:
		return "Honkai Impact 3rd"
	case GamePP:
		return "Petit Planet"
	case GameHNA:
		return "Honkai: New Anima"
	default:
		return "Unknown"
	}
}
