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
	shunZi = false
	saves := make([]model.CardSize, 16)
	// 把扑克牌放如slice中
	for _, card := range cardsSize {
		saves[card] = card
	}

	_, ok := sizeMap[model.CardSizeAce]
	if ok {
		saves[1] = model.CardSizeAce
	}
	// 判断数组是否连续 倒序遍历
	sum := 0
	for i := 0; i < len(saves); i++ {
		// slice有值
		if saves[i] == 0x00 {
			sum = 0
		} else {
			sum++
			// 5个连续
			if sum >= 5 {
				shunZi = true
				max = saves[i] // 返回顺子的最大值
				return
			}
		}

	}
	return
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
	return b.HighCardCompareByLen(len(b.CardsSize1), b.CardsSize1, b.CardsSize2)
}

func (b *BaseCardCom) OnePairCom() model.Result {
	// 用于存放单牌的面值
	cardSizeSlice1 := make([]model.CardSize, len(b.CardSizeMap1))
	cardSizeSlice2 := make([]model.CardSize, len(b.CardSizeMap1))
	// 用于存放对子的面值
	var pair1 model.CardSize
	var pair2 model.CardSize
	i := 0
	for k, v := range b.CardSizeMap1 {
		if v == 2 {
			pair1 = k
			continue
		}
		cardSizeSlice1[i] = k
		i++
	}
	i = 0
	for k, v := range b.CardSizeMap2 {
		if v == 2 {
			pair2 = k
			continue
		}
		cardSizeSlice2[i] = k
		i++
	}
	// 先比较对子的大小
	if pair1 > pair2 {
		return model.Win
	} else if pair1 < pair2 {
		return model.Lose
	} else {
		// 再单牌大小
		return b.HighCardCompareByLen(len(b.CardsSize1)-2, cardSizeSlice1, cardSizeSlice2)
	}

}

func (b *BaseCardCom) TwoPairCom() model.Result {
	// 用于存放两对的牌子
	num := 2
	pairs1 := make([]model.CardSize, num)
	pairs2 := make([]model.CardSize, num)
	// 用于存放单牌
	num = len(b.CardsSize1) - 4
	val1s := make([]model.CardSize, num)
	val2s := make([]model.CardSize, num)
	var val2 model.CardSize

	i := 0
	j := 0
	for k, v := range b.CardSizeMap1 {
		if v == 2 {
			pairs1[i] = k
			i++
			continue
		}
		val1 = k
	}
	i = 0
	for k, v := range b.CardSizeMap2 {
		if v == 2 {
			pairs2[i] = k
			i++
		} else {
			val2 = k
		}

	}
	// 比较对子的大小
	var result model.Result
	result = b.HighCardCompareByLen(2, pairs1, pairs2)
	if result != 0 {
		return result
	}

	// 再比较单牌的大小
	if val1 > val2 {
		return model.Win
	} else if val1 < val2 {
		return model.Lose
	} else {
		return model.Draw
	}
}

func (b *BaseCardCom) ThreeOfAKindCom() model.Result {
	// 用于存放单牌的面值
	cardSizeSlice1 := make([]model.CardSize, len(b.CardSizeMap1))
	cardSizeSlice2 := make([]model.CardSize, len(b.CardSizeMap1))
	// 用于存放三条的面值
	var three1 model.CardSize
	var three2 model.CardSize
	i := 0
	for k, v := range b.CardSizeMap1 {
		cardSizeSlice1[i] = k
		if v == 3 {
			three1 = k
		} else {
			i++
		}
	}
	i = 0
	for k, v := range b.CardSizeMap2 {
		cardSizeSlice2[i] = k
		if v == 3 {
			three2 = k
		} else {
			i++
		}
	}
	// 先比较三条的面值
	if three1 > three2 {
		return model.Win
	} else if three1 < three2 {
		return model.Lose
	} else {
		// 再比较单牌的
		return b.HighCardCompareByLen(len(cardSizeSlice1)-3, cardSizeSlice1, cardSizeSlice2)
	}
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
	return b.StraightCom()
}

func (b *BaseCardCom) FourHouseCom() model.Result {
	// 存放四条的面值
	var four1 model.CardSize
	var four2 model.CardSize
	// 存放单牌的面值
	num := len(b.CardsSize1) - 4
	val1s := make([]model.CardSize, 0, num)
	val2s := make([]model.CardSize, 0, num)

	for k, v := range b.CardSizeMap1 {
		if v == 4 {
			four1 = k
		} else {
			val1s = append(val1s, k)
		}
	}
	for k, v := range b.CardSizeMap2 {
		if v == 4 {
			four2 = k
		} else {
			val2s = append(val2s, k)
		}
	}
	// 先比较4条大小
	if four1 > four2 {
		return model.Win
	} else if four1 < four2 {
		return model.Lose
	} else {
		return b.HighCardCompareByLen(num, val1s, val2s)
	}
}

func (b *BaseCardCom) FullHouseCom() model.Result {
	var three1 model.CardSize
	var three2 model.CardSize
	// 存放对子的面值
	var two1 model.CardSize
	var two2 model.CardSize
	for k, v := range b.CardSizeMap1 {
		if v == 3 {
			three1 = k
		} else {
			two1 = k
		}
	}
	for k, v := range b.CardSizeMap2 {
		if v == 3 {
			three2 = k
		} else {
			two2 = k
		}
	}
	// 先对比3条的面值
	if three1 > three2 {
		return model.Win
	} else if three1 < three2 {
		return model.Lose
	} else {
		// 再对比对子的面值
		if two1 > two2 {
			return model.Win
		} else if two1 < two2 {
			return model.Lose
		} else {
			return model.Draw
		}
	}
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
