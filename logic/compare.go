package logic

import (
	"fmt"
	"path"
	"poker/dao"
	"poker/model"
	"runtime"
	"sort"
	"time"
)

// cardCom用于同类型比较的函数传递参数
type cardCom struct {
	cardSizeMap1           map[model.CardSize]int
	cardSizeMap2           map[model.CardSize]int
	max1, max2             model.CardSize
	cardsSize1, cardsSize2 []model.CardSize
}

// JudgmentGroupNew 判断牌的类型
func JudgmentGroupNew(card string) (model.CardType, map[model.CardSize]int, model.CardSize, []model.CardSize) {
	var resMax model.CardSize
	var judeCardType model.CardType
	// 扫描牌 分别放好大小，花色   --key放的是花色或是面值，--value放的是出现的次数
	sizeMap, colorsMap, cardsSize := CardsSplitMapCount(card)
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
			isShun, max := IsShunZiNew(cardsSize, sizeMap)
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
		isShun, max := IsShunZiNew(cardsSize, sizeMap)
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

// IsShunZiNew 判断是否是顺子 返回顺子的最大值和是否是顺子
func IsShunZiNew(cardsSize []model.CardSize, sizeMap map[model.CardSize]int) (shunZi bool, max model.CardSize) {
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

// MyQuickSortCardFace 快排 对字节 逆序
func MyQuickSortCardFace(cards []model.CardFace) []model.CardFace {
	if len(cards) <= 1 {
		return cards
	}
	splitdata := cards[0]                 // 第一个数据
	low := make([]model.CardFace, 0, 0)   // 比我小的数据
	hight := make([]model.CardFace, 0, 0) // 比我大的数据
	mid := make([]model.CardFace, 0, 0)   // 与我一样大的数据
	mid = append(mid, splitdata)          // 加入一个
	for i := 1; i < len(cards); i++ {
		if cards[i] > splitdata {
			low = append(low, cards[i])
		} else if cards[i] < splitdata {
			hight = append(hight, cards[i])
		} else {
			mid = append(mid, cards[i])
		}
	}
	low, hight = MyQuickSortCardFace(low), MyQuickSortCardFace(hight)
	myarr := append(append(low, mid...), hight...)
	return myarr
}

// QuickSort 系统排序函数 倒序
func QuickSort(cards []model.CardSize) {
	sort.Slice(cards, func(i, j int) bool {
		return cards[i] > cards[j]
	})
}

// SingleCardCompareSizeNew 同类型单牌比较 返回值是比较结果 0是平局 1是前面赢 2是后面赢
func (com *cardCom) SingleCardCompareSizeNew() model.Result {
	// 比较5张牌的面值
	return SingleCardSizeCom(5, com.cardsSize1, com.cardsSize2)
}

// SingleCardSizeCom 对比单牌 大小0是平局 1是前面赢 2是后面赢
func SingleCardSizeCom(comLen int, cardSizeSlice1, cardSizeSlice2 []model.CardSize) model.Result {
	// 对传进来的slice逆序排序
	QuickSort(cardSizeSlice1)
	QuickSort(cardSizeSlice2)

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

// aPairComNew 同类型一对比较 0是平局 1是前面赢 2是后面赢
func (com *cardCom) aPairComNew() model.Result {
	// 用于存放单牌的面值
	cardSizeSlice1 := make([]model.CardSize, len(com.cardSizeMap1))
	cardSizeSlice2 := make([]model.CardSize, len(com.cardSizeMap1))
	// 用于存放对子的面值
	var pair1 model.CardSize
	var pair2 model.CardSize
	i := 0
	for k, v := range com.cardSizeMap1 {
		if v == 2 {
			pair1 = k
			continue
		}
		cardSizeSlice1[i] = k
		i++
	}
	i = 0
	for k, v := range com.cardSizeMap2 {
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
		return SingleCardSizeCom(3, cardSizeSlice1, cardSizeSlice2)
	}

}

// twoPairComNew 同类型的两对比较 0是平局 1是前面赢 2是后面赢
func (com *cardCom) twoPairComNew() model.Result {
	// 用于存放两对的牌子
	num := 2
	pairs1 := make([]model.CardSize, num)
	pairs2 := make([]model.CardSize, num)
	// 用于存放单牌
	var val1 model.CardSize
	var val2 model.CardSize

	i := 0
	for k, v := range com.cardSizeMap1 {
		if v == 2 {
			pairs1[i] = k
			i++
			continue
		}
		val1 = k
	}
	i = 0
	for k, v := range com.cardSizeMap2 {
		if v == 2 {
			pairs2[i] = k
			i++
		} else {
			val2 = k
		}

	}
	// 比较对子的大小
	var result model.Result
	result = SingleCardSizeCom(2, pairs1, pairs2)
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

// onlyThreeComNew 同类型的三条比较 0是平局 1是前面赢 2是后面赢
func (com *cardCom) onlyThreeComNew() model.Result {
	// 用于存放单牌的面值
	cardSizeSlice1 := make([]model.CardSize, len(com.cardSizeMap1))
	cardSizeSlice2 := make([]model.CardSize, len(com.cardSizeMap1))
	// 用于存放三条的面值
	var three1 model.CardSize
	var three2 model.CardSize
	i := 0
	for k, v := range com.cardSizeMap1 {
		cardSizeSlice1[i] = k
		if v == 3 {
			three1 = k
		} else {
			i++
		}
	}
	i = 0
	for k, v := range com.cardSizeMap2 {
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
		return SingleCardSizeCom(2, cardSizeSlice1, cardSizeSlice2)
	}
}

// onlyShunZiNew 同类型顺子的比较 0是平局 1是前面赢 2是后面赢
func (com *cardCom) onlyShunZiNew() model.Result {
	// max 是顺子的最大的牌，只要比较这张牌就行了
	if com.max1 > com.max2 {
		return model.Win
	} else if com.max1 < com.max2 {
		return model.Lose
	}
	return model.Draw
}

// onlySameFlowerNew 是同类型同花的比较 0是平局 1是前面赢 2是后面赢
func (com *cardCom) onlySameFlowerNew() model.Result {
	// 同类型同花 只要比较牌面值最大的，可以看着是单牌比较面值大小
	return com.SingleCardCompareSizeNew()
}

// straightFlushNew 同类型同花顺比较 0是平局 1是前面赢 2是后面赢
func (com *cardCom) straightFlushNew() model.Result {
	// 同类型同花顺比较，可以看作顺子之间比较
	return com.onlyShunZiNew()
}

// fourComNew 同类型4条比较 0是平局 1是前面赢 2是后面赢
func (com *cardCom) fourComNew() model.Result {
	// 存放四条的面值
	var four1 model.CardSize
	var four2 model.CardSize
	// 存放单牌的面值
	var val1 model.CardSize
	var val2 model.CardSize

	for k, v := range com.cardSizeMap1 {
		if v == 4 {
			four1 = k
		} else {
			val1 = k
		}
	}
	for k, v := range com.cardSizeMap2 {
		if v == 4 {
			four2 = k
		} else {
			val2 = k
		}
	}
	// 先比较4条大小
	if four1 > four2 {
		return model.Win
	} else if four1 < four2 {
		return model.Lose
	} else {
		// 再比较单牌的大小
		if val1 > val2 {
			return model.Win
		} else if val1 < val2 {
			return model.Lose
		} else {
			return model.Draw
		}
	}
}

// threeAndTwoNew 同类型3带2比较  0是平局 1是前面赢 2是后面赢
func (com *cardCom) threeAndTwoNew() model.Result {
	// 存放3条的面值
	var three1 model.CardSize
	var three2 model.CardSize
	// 存放对子的面值
	var two1 model.CardSize
	var two2 model.CardSize
	for k, v := range com.cardSizeMap1 {
		if v == 3 {
			three1 = k
		} else {
			two1 = k
		}
	}
	for k, v := range com.cardSizeMap2 {
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

// GetCurrentAbPathByCaller 得到项目的路径
func GetCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
		abPath = path.Dir(abPath)
	}
	return abPath
}

func compare(alices, bobs string) model.Result {

	// 分牌型
	val1, cardSizesMap1, max1, cardsSize1 := JudgmentGroupNew(alices)
	val2, cardSizesMap2, max2, cardsSize2 := JudgmentGroupNew(bobs)
	if val1 < val2 {
		return model.Win
	}
	if val1 > val2 {
		return model.Lose
	}
	// 牌型相同的处理情况
	// ...
	cardCom := cardCom{
		cardSizeMap1: cardSizesMap1,
		cardSizeMap2: cardSizesMap2,
		max1:         max1,
		max2:         max2,
		cardsSize1:   cardsSize1,
		cardsSize2:   cardsSize2,
	}
	switch val1 {
	case model.HighCard:
		// 同类型下的单张大牌比较
		return cardCom.SingleCardCompareSizeNew()
	case model.OnePair:
		// 同类型的一对
		return cardCom.aPairComNew()
	case model.TwoPair:
		// 同类型两对
		return cardCom.twoPairComNew()
	case model.ThreeOfAKind:
		// 同类型三条
		return cardCom.onlyThreeComNew()
	case model.Straight:
		// 同类型顺子
		return cardCom.onlyShunZiNew()
	case model.Flush:
		// 同类型同花
		return cardCom.onlySameFlowerNew()
	case model.FullHouse:
		// 同类型3带2
		return cardCom.threeAndTwoNew()
	case model.FourHouse:
		// 同类型四条
		return cardCom.fourComNew()
	case model.StraightFlush: // 同类型同花顺
		return cardCom.straightFlushNew()
	case model.RoyalFlush:
		//皇家同花顺
		return model.Draw
	default:
		panic("unknown card type!")
	}
	// 最后比较结果

}

// PokerMan 5张遍历判断 文件扑克牌的函数
func PokerMan() {
	file := GetCurrentAbPathByCaller()
	file += "/resources/match_result.json"
	alices, bobs, results := dao.ReadFile(file)
	t1 := time.Now()
	k := 0
	// 遍历全部对比
	for i := 0; i < len(alices); i++ {
		result := compare(alices[i], bobs[i])
		// 打印判断出错的信息
		if result != results[i] {
			k++
			fmt.Printf("[%#v]5张判断错误--->alice:%#v,bob:%#v<----- ===>文档的结果：%#v, 我的结果:%#v <==\n", k, alices[i], bobs[i], results[i], result)
		}
	}
	t2 := time.Now()
	fmt.Println("time--->", t2.Sub(t1))

}

// CardsSplit 将手牌的面值和花色分拆开
func CardsSplit(cards string) ([]model.CardFace, []model.CardColor) {
	num := len(cards) / 2
	faces := make([]model.CardFace, 0, num)
	colors := make([]model.CardColor, 0, num)
	for index, value := range cards {
		if index%2 == 0 {
			faces = append(faces, model.CardFace(value))
		} else {
			colors = append(colors, model.CardColor(value))
		}
	}
	return faces, colors
}

// CardsSplitMapCount 统计手牌中面值和花色的数量 并且把面值转换为可比较大小的
// 724.7 ns/op -> 429.0 ns/op(直接map++而不取出来判断是否已经存在key) ->374.8 ns/op (用下标赋值而不用append,除2用位运算)
func CardsSplitMapCount(cards string) (map[model.CardSize]int, map[model.CardColor]int, []model.CardSize) {
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
