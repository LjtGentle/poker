package seven

import (
	"fmt"
	"poker/dao"
	"poker/logic"
	"poker/model"
	"poker/util"
	"time"
)

type sevenCom struct {
	logic.BaseCardCom
}

func (s *sevenCom) JudgmentCardType(card string) (model.CardType, map[model.CardSize]int, model.CardSize, []model.CardSize) {
	sizeMap, colorsMap, cardsSize, colorsSlice := s.CardsSplitMapCount(card)
	// 获取map的长度
	sizeLen := len(sizeMap)
	var resMax model.CardSize
	var color model.CardColor
	var cardType model.CardType
	for k, v := range colorsMap {
		if v >= 5 { // 5个花色一样的
			color = k
			break
		}
	}
	if color != 0x00 {
		// 重写cardSize
		newCardsSize := make([]model.CardSize, len(cardsSize))
		i := 0
		for index, value := range colorsSlice {
			if value == color {
				newCardsSize[i] = cardsSize[index]
				i++
			}
		}
		// 已经是同花
		// 然后判断是不是顺子
		cardType = model.Flush
		isStraight := false
		isStraight, resMax = s.IsStraight(newCardsSize, sizeMap)
		if isStraight {
			// 同花顺
			cardType = model.StraightFlush
			if resMax == model.CardSizeAce {
				cardType = model.RoyalFlush
			}
		}
		if i >= 5 {
			// 需要排序
			util.QuickSort(newCardsSize)
		}
		return cardType, sizeMap, resMax, newCardsSize[:5]
	}
	// 然后再根据存放面值的map长度来判断类型
	switch sizeLen {
	case 7: // 不是顺子 就是7个单牌
		// 判断是不是顺子
		isStraight := false
		if isStraight, resMax = s.IsStraight(cardsSize, sizeMap); isStraight {
			return model.Straight, sizeMap, resMax, cardsSize
		}
		return model.HighCard, sizeMap, resMax, cardsSize
	case 6: // 1对 或是 顺子
		// 判断是不是顺子
		isStraight := false
		if isStraight, resMax = s.IsStraight(cardsSize, sizeMap); isStraight {
			return model.Straight, sizeMap, resMax, cardsSize
		}
		return model.OnePair, sizeMap, resMax, cardsSize
	case 5: // 可以是顺子 两对 或是 3条
		// 顺子大先判断是不是顺子
		isStraight := false
		if isStraight, resMax = s.IsStraight(cardsSize, sizeMap); isStraight {
			return model.Straight, sizeMap, resMax, cardsSize
		}
		// cardSizeMap 的key是面值 value是该面值出现的次数
		// 然后判断是不是3条
		for _, v := range sizeMap {
			if v == 3 {
				return model.ThreeOfAKind, sizeMap, resMax, cardsSize
			}
		}
		// 2对
		return model.TwoPair, sizeMap, resMax, cardsSize
	case 4: // 可以是 4条 3带2  两对（3个对子）
		for _, v := range sizeMap {
			if v == 4 {
				return model.FourHouse, sizeMap, resMax, cardsSize
			} else if v == 3 {
				return model.FullHouse, sizeMap, resMax, cardsSize
			}
		}
		// 剩下两对
		return model.TwoPair, sizeMap, resMax, cardsSize
	case 3: // 4条（4条1对） 3带2（3条和3条， 3条和两对）
		for _, v := range sizeMap {
			if v == 4 {
				return model.FourHouse, sizeMap, resMax, cardsSize
			}
		}
		return model.FullHouse, sizeMap, resMax, cardsSize
	case 2: // 4条（4条和3条）
		return model.FourHouse, sizeMap, resMax, cardsSize
	default:
		panic(fmt.Sprintf("unknown card size,sizeLen=%d", sizeLen))

	}
}

func init() {
	logic.RegisterPoker(model.SevenGameType, &sevenCom{})
}

func (s *sevenCom) Compare(alices, bobs string) model.Result {
	// 分牌型
	val1, cardSizesMap1, max1, cardsSize1 := s.JudgmentCardType(alices)
	val2, cardSizesMap2, max2, cardsSize2 := s.JudgmentCardType(bobs)
	if val1 < val2 {
		return model.Win
	}
	if val1 > val2 {
		return model.Lose
	}
	// 牌型相同的处理情况

	s.CardsSize1 = cardsSize1
	s.CardsSize2 = cardsSize2
	s.Max1 = max1
	s.Max2 = max2
	s.CardSizeMap1 = cardSizesMap1
	s.CardSizeMap2 = cardSizesMap2
	switch val1 {
	case model.HighCard:
		// 同类型下的单张大牌比较
		return s.HighCardCompareCom()
	case model.OnePair:
		// 同类型的一对
		return s.OnePairCom()
	case model.TwoPair:
		// 同类型两对
		return s.TwoPairCom()
	case model.ThreeOfAKind:
		// 同类型三条
		return s.ThreeOfAKindCom()
	case model.Straight:
		// 同类型顺子
		return s.StraightCom()
	case model.Flush:
		// 同类型同花
		return s.FlushCom()
	case model.FullHouse:
		// 同类型3带2
		return s.FullHouseCom()
	case model.FourHouse:
		// 同类型四条
		return s.FourHouseCom()
	case model.StraightFlush: // 同类型同花顺
		return s.StraightFlushCom()
	case model.RoyalFlush:
		//皇家同花顺
		return model.Draw
	default:
		panic("unknown card type!")
	}
	// 最后比较结果
}

// PokerMan 5张遍历判断 文件扑克牌的函数
func (s *sevenCom) PokerMan() {
	file := dao.GetCurrentAbPathByCaller()
	file += "/resources/seven_cards_with_ghost.json"
	alices, bobs, results := dao.ReadFile(file)
	t1 := time.Now()
	k := 0
	// 遍历全部对比
	for i := 0; i < len(alices); i++ {
		result := s.Compare(alices[i], bobs[i])
		// 打印判断出错的信息
		if result != results[i] {
			k++
			fmt.Printf("[%#v]7张判断错误--->alice:%#v,bob:%#v<----- ===>文档的结果：%#v, 我的结果:%#v <==\n", k, alices[i], bobs[i], results[i], result)
		}
	}
	t2 := time.Now()
	fmt.Println("time--->", t2.Sub(t1))

}
