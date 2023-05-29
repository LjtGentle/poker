package sevenAndJoker

import (
	"poker/logic"
	"poker/model"
)

type sevenAndJoker struct {
	*logic.BaseCardCom
}

func (s *sevenAndJoker) JudgmentCardType(card string) (model.CardType, map[model.CardSize]int, model.CardSize, []model.CardSize) {
	//TODO implement me
	panic("implement me")
}

func (s *sevenAndJoker) IsStraight(cardsSize []model.CardSize, sizeMap map[model.CardSize]int) (shunZi bool, max model.CardSize) {
	//TODO implement me
	panic("implement me")
}

func (s *sevenAndJoker) HighCardCompareCom() model.Result {
	//TODO implement me
	panic("implement me")
}

func (s *sevenAndJoker) OnePairCom() model.Result {
	//TODO implement me
	panic("implement me")
}

func (s *sevenAndJoker) TwoPairCom() model.Result {
	//TODO implement me
	panic("implement me")
}

func (s *sevenAndJoker) ThreeOfAKindCom() model.Result {
	//TODO implement me
	panic("implement me")
}

func (s *sevenAndJoker) FlushCom() model.Result {
	//TODO implement me
	panic("implement me")
}

func (s *sevenAndJoker) StraightFlushCom() model.Result {
	//TODO implement me
	panic("implement me")
}

func (s *sevenAndJoker) FourHouseCom() model.Result {
	//TODO implement me
	panic("implement me")
}

func (s *sevenAndJoker) FullHouseCom() model.Result {
	//TODO implement me
	panic("implement me")
}

func init() {
	logic.RegisterPoker(model.SevenAndJokerGameType, &sevenAndJoker{})
}
