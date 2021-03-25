package bot_test

import (
	"testing"

	. "github.com/Heart-plus-N/habitica-multi-bot/bot"
)

func TestBuildInterest(t *testing.T) {
	res := BuildInterest(QuestEvent, UserEvent)
	if res != 3 {
		t.Errorf("BuildInterest: %d", res)
	}
}
