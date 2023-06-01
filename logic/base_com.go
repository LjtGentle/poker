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
	for i := len(saves) - 1; i > 0; i-- {
		// slice有值
		if saves[i] == 0x00 {
			sum = 0
		} else {
			sum++
			// 5个连续
			if sum >= 5 {
				shunZi = true
				max = saves[i+4] // 返回顺子的最大值
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
	return b.HighCardCompareByLen(5, b.CardsSize1, b.CardsSize2)
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
		return b.HighCardCompareByLen(3, cardSizeSlice1, cardSizeSlice2)
	}

}

func (b *BaseCardCom) TwoPairCom() model.Result {
	// 用于存放两对的牌子
	num := 4 //兼容7张牌会出现3对的情况
	pairs1 := make([]model.CardSize, num)
	pairs2 := make([]model.CardSize, num)
	// 用于存放单牌
	val1s := make([]model.CardSize, num)
	val2s := make([]model.CardSize, num)

	i := 0
	j := 0
	for k, v := range b.CardSizeMap1 {
		if v == 2 {
			pairs1[i] = k
			i++
			continue
		}
		val1s[j] = k
		j++
	}
	i = 0
	j = 0
	for k, v := range b.CardSizeMap2 {
		if v == 2 {
			pairs2[i] = k
			i++
			continue
		}
		val2s[j] = k
		j++

	}

	// 比较对子的大小
	var result model.Result
	result = b.HighCardCompareByLen(2, pairs1, pairs2)
	if result != 0 {
		return result
	}
	// 三个对子的场景
	val1s[3] = pairs1[2]
	val2s[3] = pairs2[2]

	// 再比较单牌的大小
	return b.HighCardCompareByLen(1, val1s, val2s)
}

func (b *BaseCardCom) ThreeOfAKindCom() model.Result {
	// 用于存放单牌的面值
	num := len(b.CardsSize1) - 2
	cardSizeSlice1 := make([]model.CardSize, num)
	cardSizeSlice2 := make([]model.CardSize, num)
	// 用于存放三条的面值
	var three1 model.CardSize
	var three2 model.CardSize
	i := 0
	for k, v := range b.CardSizeMap1 {
		if v == 3 {
			three1 = k
		} else {
			cardSizeSlice1[i] = k
			i++
		}
	}
	i = 0
	for k, v := range b.CardSizeMap2 {
		if v == 3 {
			three2 = k
		} else {
			cardSizeSlice2[i] = k
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
		return b.HighCardCompareByLen(2, cardSizeSlice1, cardSizeSlice2)
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
		return b.HighCardCompareByLen(1, val1s, val2s)
	}
}

func (b *BaseCardCom) FullHouseCom() model.Result {
	threes1 := make([]model.CardSize, 2)
	threes2 := make([]model.CardSize, 2)
	// 存放对子的面值
	twos1 := make([]model.CardSize, 2)
	twos2 := make([]model.CardSize, 2)
	i := 0
	j := 0
	for k, v := range b.CardSizeMap1 {
		if v == 3 {
			threes1[i] = k
			i++
			continue
		}
		if v == 2 {
			twos1[j] = k
			j++
		}
	}
	i = 0
	j = 0
	for k, v := range b.CardSizeMap2 {
		if v == 3 {
			threes2[i] = k
			i++
			continue
		}
		if v == 2 {
			twos2[j] = k
			j++
		}
	}
	// 如果有两个三对的情况
	if threes1[1] != 0x00 {
		twos1[1] = threes1[0]
		if threes1[1] < threes1[0] {
			twos1[1] = threes1[1]
		}
	}
	if threes2[1] != 0x00 {
		twos2[1] = threes2[0]
		if threes2[1] < threes2[0] {
			twos2[1] = threes2[1]
		}
	}
	// 先对比3条的面值
	if result := b.HighCardCompareByLen(1, threes1, threes2); result != model.Draw {
		return result
	} else {
		// 再对比对子的面值
		return b.HighCardCompareByLen(1, twos1, twos2)
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

func (b *BaseCardCom) CardsSplitMapCount(cards string) (
	map[model.CardSize]int, map[model.CardColor]int, []model.CardSize, []model.CardColor) {
	num := len(cards) >> 1
	sizeMap := make(map[model.CardSize]int, num)
	colorsMap := make(map[model.CardColor]int, num)
	cardsSize := make([]model.CardSize, num)
	colorsSlice := make([]model.CardColor, num)
	var size model.CardSize
	var colors model.CardColor
	for i := 0; i < num; i++ {
		size = model.CardFace2SizeSlice[cards[i<<1]]
		cardsSize[i] = size
		sizeMap[size]++
		colors = model.CardColor(cards[(i<<1)+1])
		colorsMap[colors]++
		colorsSlice[i] = colors
	}
	return sizeMap, colorsMap, cardsSize, colorsSlice
}

func (b *BaseCardCom) PokerMan() {
	//TODO implement me
	panic("implement me")
}
