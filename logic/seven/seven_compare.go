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

func (s *sevenCom) TwoPairCom() model.Result {
	// 存放对子的slice
	pairs1 := make([]model.CardSize, 2)
	pairs2 := make([]model.CardSize, 2)
	// 存放单牌的slice
	vals1 := make([]model.CardSize, 3)
	vals2 := make([]model.CardSize, 3)
	j := 0
	i := 0
	for k, v := range s.CardSizeMap1 {
		if v == 2 {
			pairs1[i] = k
			i++
		} else {
			vals1[j] = k
			j++
		}
	}
	i = 0
	j = 0
	for k, v := range s.CardSizeMap2 {
		if v == 2 {
			pairs2[i] = k
			i++
		} else {
			vals2[j] = k
			j++
		}
	}
	if result := s.HighCardCompareByLen(2, pairs1, pairs2); result != model.Draw {
		return result
	}
	return s.HighCardCompareByLen(3, vals1, vals2)
}

func init() {
	logic.RegisterPoker(model.SevenGameType, &sevenCom{})
}
