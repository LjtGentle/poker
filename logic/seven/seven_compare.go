package seven

import (
	"poker/logic"
	"poker/model"
)

type sevenCom struct {
	*logic.BaseCardCom
}

func (s *sevenCom) JudgmentCardType(card string) (model.CardType, map[model.CardSize]int, model.CardSize, []model.CardSize) {
	//TODO implement me
	panic("implement me")
}

func (s *sevenCom) IsStraight(cardsSize []model.CardSize, sizeMap map[model.CardSize]int) (shunZi bool, max model.CardSize) {
	//TODO implement me
	panic("implement me")
}

func (s *sevenCom) HighCardCompareCom() model.Result {
	//TODO implement me
	panic("implement me")
}

func (s *sevenCom) OnePairCom() model.Result {
	//TODO implement me
	panic("implement me")
}

func (s *sevenCom) TwoPairCom() model.Result {
	//TODO implement me
	panic("implement me")
}

func (s *sevenCom) ThreeOfAKindCom() model.Result {
	//TODO implement me
	panic("implement me")
}

func (s *sevenCom) FlushCom() model.Result {
	//TODO implement me
	panic("implement me")
}

func (s *sevenCom) StraightFlushCom() model.Result {
	//TODO implement me
	panic("implement me")
}

func (s *sevenCom) FourHouseCom() model.Result {
	//TODO implement me
	panic("implement me")
}

func (s *sevenCom) FullHouseCom() model.Result {
	//TODO implement me
	panic("implement me")
}

func init() {
	logic.RegisterPoker(model.SevenGameType, &sevenCom{})
}
