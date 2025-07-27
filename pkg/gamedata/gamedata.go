package gamedata

import (
	"fmt"

	"github.com/owen-crook/board-game-tracker-go-common/pkg/games"
)

type ScoringCategory struct {
	ShortName string
	LongName  string
}

var (
	WyrmspanScoringCategories []ScoringCategory = []ScoringCategory{
		{ShortName: "guild", LongName: "markers on the dragon guild"},
		{ShortName: "tails-on-dragons", LongName: "printed on dragons"},
		{ShortName: "end-game-abilities", LongName: "from end-game abilities"},
		{ShortName: "eggs", LongName: "per egg*"},
		{ShortName: "cached-resources", LongName: "per cached resource*"},
		{ShortName: "public-objectives", LongName: "from public objectives (ties are friendly--see rulebook p.15)"},
		{ShortName: "remaining-coins-items", LongName: "from remaining coins & items 1 per coin 1 per 4 food, dragon card, cave card (in any combination) (round down)"},
	}

	WingspanScoringCategories []ScoringCategory = []ScoringCategory{
		{ShortName: "birds", LongName: "birds"},
		{ShortName: "bonus-cards", LongName: "bonus cards"},
		{ShortName: "end-of-round-goals", LongName: "end-of-round goals"},
		{ShortName: "eggs", LongName: "eggs"},
		{ShortName: "food-on-cards", LongName: "food on cards"},
		{ShortName: "tucked-cards", LongName: "tucked cards"},
	}
)

func GetScoringCategoriesByGame(game games.Game) ([]ScoringCategory, error) {
	if !games.IsSupportedGame(game) {
		return nil, fmt.Errorf("invalid game: %s", game)
	}
	switch game {
	case games.Wyrmspan:
		return WyrmspanScoringCategories, nil
	case games.Wingspan:
		return WingspanScoringCategories, nil
	default:
		return nil, fmt.Errorf("unable to find scoring categories for game: %s", game)
	}
}

func GetScorecardGeometryByGame(game games.Game) (string, error) {
	if !games.IsSupportedGame(game) {
		return "", fmt.Errorf("invalid game: %s", game)
	}
	switch game {
	case games.Wyrmspan, games.Wingspan:
		return "The rows of the scorecard represent score categories and the columns represent different players of that game. The values at each row and column represent the score a given player acheived for that category. ", nil
	default:
		return "", fmt.Errorf("unable to find scorecard geometry for game: %s", game)
	}
}

func GetExampleJsonByGame(game games.Game) (map[string]interface{}, error) {
	if !games.IsSupportedGame(game) {
		return nil, fmt.Errorf("invalid game: %s", game)
	}
	switch game {
	case games.Wyrmspan:
		return map[string]interface{}{
			"date":     "2025-06-22",
			"location": "minty",
			"players": []map[string]interface{}{
				{
					"name":                    "SM",
					"markers-on-dragon-guild": 3,
					"tails-on-dragons":        32,
					"end-game-abilities":      13,
					"eggs":                    12,
					"cached-resources":        1,
					"tucked-cards":            5,
					"public-objectives":       15,
					"remaining-coins-items":   0,
					"total":                   81,
				},
				{
					"name":                    "WS",
					"markers-on-dragon-guild": 7,
					"tails-on-dragons":        24,
					"end-game-abilities":      16,
					"eggs":                    12,
					"cached-resources":        7,
					"tucked-cards":            6,
					"public-objectives":       9,
					"remaining-coins-items":   1,
					"total":                   82,
				},
			},
		}, nil
	case games.Wingspan:
		return map[string]interface{}{
			"date":     "2025-06-23",
			"location": "crook nook",
			"players": []map[string]interface{}{
				{
					"name":               "OC",
					"birds":              31,
					"bonus-cards":        10,
					"end-of-round-goals": 18,
					"eggs":               18,
					"food-on-cards":      5,
					"tucked-cards":       2,
					"total":              84,
				},
				{
					"name":               "JB",
					"birds":              49,
					"bonus-cards":        4,
					"end-of-round-goals": 17,
					"eggs":               9,
					"food-on-cards":      4,
					"tucked-cards":       7,
					"total":              90,
				},
				{
					"name":               "MS",
					"birds":              34,
					"bonus-cards":        11,
					"end-of-round-goals": 13,
					"eggs":               6,
					"food-on-cards":      2,
					"tucked-cards":       1,
					"total":              67,
				},
			},
		}, nil
	default:
		return nil, fmt.Errorf("unable to find example json for game: %s", game)
	}
}
