package gamedata

import (
	"testing"

	"github.com/owen-crook/board-game-tracker-go-common/pkg/games"
)

func TestGetScoringCategoriesByGame(t *testing.T) {
	cats, err := GetScoringCategoriesByGame(games.Wyrmspan)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if len(cats) != len(WyrmspanScoringCategories) {
		t.Errorf("expected %d categories, got %d", len(WyrmspanScoringCategories), len(cats))
	}

	cats, err = GetScoringCategoriesByGame(games.Wingspan)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if len(cats) != len(WingspanScoringCategories) {
		t.Errorf("expected %d categories, got %d", len(WingspanScoringCategories), len(cats))
	}

	_, err = GetScoringCategoriesByGame(games.Game("unknown"))
	if err == nil {
		t.Errorf("expected error for unknown game")
	}
}

func TestGetScorecardGeometryByGame(t *testing.T) {
	geom, err := GetScorecardGeometryByGame(games.Wyrmspan)
	if err != nil || geom == "" {
		t.Errorf("expected geometry for Wyrmspan, got error: %v", err)
	}
	geom, err = GetScorecardGeometryByGame(games.Wingspan)
	if err != nil || geom == "" {
		t.Errorf("expected geometry for Wingspan, got error: %v", err)
	}
	_, err = GetScorecardGeometryByGame(games.Game("unknown"))
	if err == nil {
		t.Errorf("expected error for unknown game")
	}
}

func TestGetExampleJsonByGame(t *testing.T) {
	json, err := GetExampleJsonByGame(games.Wyrmspan)
	if err != nil || json == nil {
		t.Errorf("expected example json for Wyrmspan, got error: %v", err)
	}
	json, err = GetExampleJsonByGame(games.Wingspan)
	if err != nil || json == nil {
		t.Errorf("expected example json for Wingspan, got error: %v", err)
	}
	_, err = GetExampleJsonByGame(games.Game("unknown"))
	if err == nil {
		t.Errorf("expected error for unknown game")
	}
}
