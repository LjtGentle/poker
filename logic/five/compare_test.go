package five

import (
	"github.com/stretchr/testify/assert"
	"poker/model"
	"testing"
)

func TestIsShunZiNew(t *testing.T) {
	cards := "6s5h4c3s2c"
	sizeMap, _, cardsSize := CardsSplitMapCount(cards)

	flag, max := IsShunZiNew(cardsSize, sizeMap)
	assert.Equal(t, flag, true)
	assert.Equal(t, max, model.CardSizeSix)
}

func TestPokerMain(t *testing.T) {
	PokerMan()
}
