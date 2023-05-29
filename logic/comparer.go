package logic

import "poker/model"

type IComparer interface {
	// JudgmentCardType 判断牌型
	JudgmentCardType(card string) (model.CardType, map[model.CardSize]int, model.CardSize, []model.CardSize)
	// IsStraight 判断是否是顺子
	IsStraight(cardsSize []model.CardSize, sizeMap map[model.CardSize]int) (shunZi bool, max model.CardSize)
	// HighCardCompareByLen 高排比较大小
	HighCardCompareByLen(comLen int, cardSizeSlice1, cardSizeSlice2 []model.CardSize) model.Result
	// HighCardCompareCom 高排比较大小
	HighCardCompareCom() model.Result
	// OnePairCom 一对比较
	OnePairCom() model.Result
	// TwoPairCom 两对比较
	TwoPairCom() model.Result
	// ThreeOfAKindCom 三条比较
	ThreeOfAKindCom() model.Result
	// StraightCom 顺子比较
	StraightCom() model.Result
	// FlushCom 同花比较
	FlushCom() model.Result
	// StraightFlushCom 同花顺比较
	StraightFlushCom() model.Result
	// FourHouseCom 四条对比
	FourHouseCom() model.Result
	// FullHouseCom  满堂彩（葫芦，三带二）比较
	FullHouseCom() model.Result
	// Compare 两幅牌比较
	Compare(alices, bobs string) model.Result
	// CardsSplitMapCount 面值转换、手牌面值和花色分隔，并且统计数量
	CardsSplitMapCount(cards string) (map[model.CardSize]int, map[model.CardColor]int, []model.CardSize)
}

var pokerFactory = make(map[uint]IComparer)

// RegisterPoker  RegisterPoker
func RegisterPoker(name uint, pool IComparer) {
	pokerFactory[name] = pool
}

// GetPoker GetPoker
func GetPoker(name uint) IComparer {
	return pokerFactory[name]
}
