package games

import (
	"testing"
)

func TestGetSupportedGames(t *testing.T) {
	games := GetSupportedGames()
	if len(games) != 2 {
		t.Errorf("expected 2 supported games, got %d", len(games))
	}
	if games[0] != Wingspan || games[1] != Wyrmspan {
		t.Errorf("expected [Wingspan, Wyrmspan], got %v", games)
	}
}

func TestIsSupportedGame(t *testing.T) {
	if !IsSupportedGame(Wingspan) {
		t.Errorf("Wingspan should be supported")
	}
	if !IsSupportedGame(Wyrmspan) {
		t.Errorf("Wyrmspan should be supported")
	}
	if IsSupportedGame(Game("unknown")) {
		t.Errorf("'unknown' should not be supported")
	}
}
