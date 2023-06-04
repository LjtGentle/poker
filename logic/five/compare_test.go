package five

import (
	"github.com/stretchr/testify/assert"
	"poker/logic"
	"poker/model"
	"testing"
)

func TestIsShunZiNew(t *testing.T) {
	cards := "6s5h4c3s2c"
	sizeMap, _, cardsSize, _ := f.CardsSplitMapCount(cards)

	flag, max := f.IsStraight(cardsSize, sizeMap)
	assert.Equal(t, flag, true)
	assert.Equal(t, max, model.CardSizeSix)
}

func TestPokerMain(t *testing.T) {
	f := logic.GetPoker(model.FiveGameType)
	f.PokerMan()
}

func TestCompare(t *testing.T) {
	c1 := "3h6c5dAsQd"
	c2 := "Qs7d7hQcTs"
	f := logic.GetPoker(model.FiveGameType)

	res := f.Compare(c1, c2)
	assert.Equal(t, model.Lose, res)

}
