package logic

import (
	"poker/model"
	"poker/util"
)

type BaseCardCom struct {
	CardSizeMap1           map[model.CardSize]int
	CardSizeMap2           map[model.CardSize]int
	Max1, Max2             model.CardSize
	CardsSize1, CardsSize2 []model.CardSize
}

var _ IComparer = &BaseCardCom{}

func (b *BaseCardCom) JudgmentCardType(card string) (model.CardType, map[model.CardSize]int, model.CardSize, []model.CardSize) {
	//TODO implement me
	panic("implement me")
}

func (b *BaseCardCom) IsStraight(cardsSize []model.CardSize, sizeMap map[model.CardSize]int) (shunZi bool, max model.CardSize) {
	//TODO implement me
	panic("implement me")
}

func (b *BaseCardCom) HighCardCompareByLen(comLen int, cardSizeSlice1, cardSizeSlice2 []model.CardSize) model.Result {
	// 对传进来的slice逆序排序
	util.QuickSort(cardSizeSlice1)
	util.QuickSort(cardSizeSlice2)

	// 一个个对比
	for i := 0; i < comLen; i++ {
		if cardSizeSlice1[i] > cardSizeSlice2[i] {
			return model.Win
		} else if cardSizeSlice1[i] < cardSizeSlice2[i] {
			return model.Lose
		}
	}
	return model.Draw
}

func (b *BaseCardCom) HighCardCompareCom() model.Result {
	//TODO implement me
	panic("implement me")
}

func (b *BaseCardCom) OnePairCom() model.Result {
	//TODO implement me
	panic("implement me")
}

func (b *BaseCardCom) TwoPairCom() model.Result {
	//TODO implement me
	panic("implement me")
}

func (b *BaseCardCom) ThreeOfAKindCom() model.Result {
	//TODO implement me
	panic("implement me")
}

func (b *BaseCardCom) StraightCom() model.Result {
	// max 是顺子的最大的牌，只要比较这张牌就行了
	if b.Max1 > b.Max2 {
		return model.Win
	} else if b.Max1 < b.Max2 {
		return model.Lose
	}
	return model.Draw
}

func (b *BaseCardCom) FlushCom() model.Result {
	return b.HighCardCompareCom()
}

func (b *BaseCardCom) StraightFlushCom() model.Result {
	//TODO implement me
	panic("implement me")
}

func (b *BaseCardCom) FourHouseCom() model.Result {
	//TODO implement me
	panic("implement me")
}

func (b *BaseCardCom) FullHouseCom() model.Result {
	//TODO implement me
	panic("implement me")
}

func (b *BaseCardCom) Compare(alices, bobs string) model.Result {
	// 分牌型
	val1, cardSizesMap1, max1, cardsSize1 := b.JudgmentCardType(alices)
	val2, cardSizesMap2, max2, cardsSize2 := b.JudgmentCardType(bobs)
	if val1 < val2 {
		return model.Win
	}
	if val1 > val2 {
		return model.Lose
	}
	// 牌型相同的处理情况

	b.CardsSize1 = cardsSize1
	b.CardsSize2 = cardsSize2
	b.Max1 = max1
	b.Max2 = max2
	b.CardSizeMap1 = cardSizesMap1
	b.CardSizeMap2 = cardSizesMap2
	switch val1 {
	case model.HighCard:
		// 同类型下的单张大牌比较
		return b.HighCardCompareCom()
	case model.OnePair:
		// 同类型的一对
		return b.OnePairCom()
	case model.TwoPair:
		// 同类型两对
		return b.TwoPairCom()
	case model.ThreeOfAKind:
		// 同类型三条
		return b.ThreeOfAKindCom()
	case model.Straight:
		// 同类型顺子
		return b.StraightCom()
	case model.Flush:
		// 同类型同花
		return b.FlushCom()
	case model.FullHouse:
		// 同类型3带2
		return b.FullHouseCom()
	case model.FourHouse:
		// 同类型四条
		return b.FourHouseCom()
	case model.StraightFlush: // 同类型同花顺
		return b.StraightFlushCom()
	case model.RoyalFlush:
		//皇家同花顺
		return model.Draw
	default:
		panic("unknown card type!")
	}
	// 最后比较结果
}

func (b *BaseCardCom) CardsSplitMapCount(cards string) (map[model.CardSize]int, map[model.CardColor]int, []model.CardSize) {
	num := len(cards) >> 1
	sizeMap := make(map[model.CardSize]int, num)
	colorsMap := make(map[model.CardColor]int, num)
	cardsSize := make([]model.CardSize, num)
	var size model.CardSize
	for i := 0; i < num; i++ {
		size = model.CardFace2SizeSlice[cards[i<<1]]
		cardsSize[i] = size
		sizeMap[size]++
		colorsMap[model.CardColor(cards[(i<<1)+1])]++
	}
	return sizeMap, colorsMap, cardsSize
}
