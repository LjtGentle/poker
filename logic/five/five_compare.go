package five

import (
	"poker/logic"
	"poker/model"
)

type fiveCardCom struct {
	*logic.BaseCardCom
}

func (f fiveCardCom) JudgmentCardType(card string) (model.CardType, map[model.CardSize]int, model.CardSize, []model.CardSize) {
	//TODO implement me
	panic("implement me")
}

func (f fiveCardCom) IsStraight(cardsSize []model.CardSize, sizeMap map[model.CardSize]int) (shunZi bool, max model.CardSize) {
	//TODO implement me
	panic("implement me")
}

func (f fiveCardCom) HighCardCompareCom() model.Result {
	//TODO implement me
	panic("implement me")
}

func (f fiveCardCom) OnePairCom() model.Result {
	//TODO implement me
	panic("implement me")
}

func (f fiveCardCom) TwoPairCom() model.Result {
	//TODO implement me
	panic("implement me")
}

func (f fiveCardCom) ThreeOfAKindCom() model.Result {
	//TODO implement me
	panic("implement me")
}

func (f fiveCardCom) FlushCom() model.Result {
	//TODO implement me
	panic("implement me")
}

func (f fiveCardCom) StraightFlushCom() model.Result {
	//TODO implement me
	panic("implement me")
}

func (f fiveCardCom) FourHouseCom() model.Result {
	//TODO implement me
	panic("implement me")
}

func (f fiveCardCom) FullHouseCom() model.Result {
	//TODO implement me
	panic("implement me")
}

func init() {
	logic.RegisterPoker(model.FiveGameType, &fiveCardCom{})
}
