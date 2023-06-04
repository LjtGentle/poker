package five

import (
	"fmt"
	"poker/dao"
	"poker/logic"
	"poker/model"
	"time"
)

func init() {
	logic.RegisterPoker(model.FiveGameType, &fiveCardCom{})
}

type fiveCardCom struct {
	logic.BaseCardCom
}

func (f *fiveCardCom) JudgmentCardType(card string) (model.CardType, map[model.CardSize]int, model.CardSize, []model.CardSize) {
	var resMax model.CardSize
	var judeCardType model.CardType
	// 扫描牌 分别放好大小，花色   --key放的是花色或是面值，--value放的是出现的次数
	sizeMap, colorsMap, cardsSize, _ := f.CardsSplitMapCount(card)
	// 获取map的长度
	sizeLen := len(sizeMap)
	colorLen := len(colorsMap)
	// 同花的时候，5个颜色一样，所以 colorLen = 1
	if colorLen > 1 {
		// 非同花
		switch sizeLen {

		case 2: // 3带2  或是 4带1
			// 遍历map value
			judeCardType = model.FullHouse
			for _, v := range sizeMap {
				if v == 4 {
					judeCardType = model.FourHouse
					return judeCardType, sizeMap, resMax, cardsSize
				}
			}
			return judeCardType, sizeMap, resMax, cardsSize
		case 3:
			// 3条 或是 两对
			judeCardType = model.TwoPair
			for _, v := range sizeMap {
				if v == 3 {
					judeCardType = model.ThreeOfAKind
					return judeCardType, sizeMap, resMax, cardsSize
				}
			}
			return judeCardType, sizeMap, resMax, cardsSize
		case 4:
			// 一对
			judeCardType = model.OnePair
			return judeCardType, sizeMap, resMax, cardsSize
		case 5:
			// 单牌或是顺子
			judeCardType = model.HighCard
			isShun, max := f.IsStraight(cardsSize, sizeMap)
			if isShun {
				resMax = max
				judeCardType = model.Straight
				return judeCardType, sizeMap, resMax, cardsSize
			}
			return judeCardType, sizeMap, resMax, cardsSize
		default:
			panic("card num is unknown !")
		}
	} else {
		// 同花 或是 同花顺 皇家同花顺
		judeCardType = model.Flush
		isShun, max := f.IsStraight(cardsSize, sizeMap)
		if isShun {
			resMax = max
			judeCardType = model.StraightFlush
			if max == model.CardSizeAce {
				judeCardType = model.RoyalFlush
				return judeCardType, sizeMap, resMax, cardsSize
			}
		}
		return judeCardType, sizeMap, resMax, cardsSize
	}

}

func (f *fiveCardCom) Compare(alices, bobs string) model.Result {
	// 分牌型
	val1, cardSizesMap1, max1, cardsSize1 := f.JudgmentCardType(alices)
	val2, cardSizesMap2, max2, cardsSize2 := f.JudgmentCardType(bobs)
	if val1 < val2 {
		return model.Win
	}
	if val1 > val2 {
		return model.Lose
	}
	// 牌型相同的处理情况

	f.CardsSize1 = cardsSize1
	f.CardsSize2 = cardsSize2
	f.Max1 = max1
	f.Max2 = max2
	f.CardSizeMap1 = cardSizesMap1
	f.CardSizeMap2 = cardSizesMap2
	switch val1 {
	case model.HighCard:
		// 同类型下的单张大牌比较
		return f.HighCardCompareCom()
	case model.OnePair:
		// 同类型的一对
		return f.OnePairCom()
	case model.TwoPair:
		// 同类型两对
		return f.TwoPairCom()
	case model.ThreeOfAKind:
		// 同类型三条
		return f.ThreeOfAKindCom()
	case model.Straight:
		// 同类型顺子
		return f.StraightCom()
	case model.Flush:
		// 同类型同花
		return f.FlushCom()
	case model.FullHouse:
		// 同类型3带2
		return f.FullHouseCom()
	case model.FourHouse:
		// 同类型四条
		return f.FourHouseCom()
	case model.StraightFlush: // 同类型同花顺
		return f.StraightFlushCom()
	case model.RoyalFlush:
		//皇家同花顺
		return model.Draw
	default:
		panic("unknown card type!")
	}
	// 最后比较结果
}

// PokerMan 5张遍历判断 文件扑克牌的函数
func (f *fiveCardCom) PokerMan() {
	file := dao.GetCurrentAbPathByCaller()
	file += "/resources/match_result.json"
	alices, bobs, results := dao.ReadFile(file)
	t1 := time.Now()
	k := 0
	// 遍历全部对比
	for i := 0; i < len(alices); i++ {
		result := f.Compare(alices[i], bobs[i])
		// 打印判断出错的信息
		if result != results[i] {
			k++
			fmt.Printf("[%#v]5张判断错误--->alice:%#v,bob:%#v<----- ===>文档的结果：%#v, 我的结果:%#v <==\n", k, alices[i], bobs[i], results[i], result)
		}
	}
	t2 := time.Now()
	fmt.Println("time--->", t2.Sub(t1))

}
