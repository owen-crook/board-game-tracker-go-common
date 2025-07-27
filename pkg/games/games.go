package games

type Game string

const (
	Wingspan Game = "wingspan"
	Wyrmspan Game = "wyrmspan"
)

func GetSupportedGames() []Game {
	return []Game{Wingspan, Wyrmspan}
}

func IsSupportedGame(game Game) bool {
	for _, supportedGame := range GetSupportedGames() {
		if supportedGame == game {
			return true
		}
	}
	return false
}
