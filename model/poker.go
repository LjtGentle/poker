package model

// Poker 存放 文件中比较数据的结构体
// Bob和Alice都是手牌 是牌面加花色的组合
// 扑克牌52张，花色黑桃spades，红心hearts，方块diamonds，草花clubs各13张，2-10，J，Q，K，A
// Face：即2-10，J，Q，K，A，其中10用T来表示。
// Color：即S(spades)、H(hearts)、D(diamonds)、C(clubs)
// 用 Face字母+小写Color字母表示一张牌，比如As表示黑桃A，其中A为牌面，s为spades，即黑桃，Ah即红心A，以此类推。 。
type Poker struct {
	Alice  string `json:"alice"`
	Bob    string `json:"bob"`
	Result Result `json:"result"`
}

// Match 用于 存放读取文件的json格式数据
type Match struct {
	Matches []*Poker `json:"matches"`
}

// Result 对比结果
type Result int

const (
	Draw Result = iota //平局
	Win                //赢局
	Lose               //败局
)

// CardType 定义一个牌型
// 1.皇家同花顺
// 2.同花顺
// 3.四条
// 4.满堂彩（葫芦，三带二）
// 5.同花
// 6.顺子
// 7.三条
// 8.两对
// 9.一对
// 10.单张大牌
type CardType uint8

const (
	RoyalFlush CardType = iota + 1
	StraightFlush
	FourHouse
	FullHouse
	Flush
	Straight
	ThreeOfAKind
	TwoPair
	OnePair
	HighCard
)

// CardColor 花色 S(spades)、H(hearts)、D(diamonds)、C(clubs)
type CardColor byte

const (
	Spades   CardColor = 's'
	Hearts   CardColor = 'h'
	Diamonds CardColor = 'd'
	Clubs    CardColor = 'c'
)

type CardFace byte

const (
	Two   CardFace = '2'
	Three CardFace = '3'
	Four  CardFace = '4'
	Five  CardFace = '5'
	Six   CardFace = '6'
	Seven CardFace = '7'
	Eight CardFace = '8'
	Nine  CardFace = '9'
	Ten   CardFace = 'T'
	Jack  CardFace = 'J'
	Queen CardFace = 'Q'
	King  CardFace = 'K'
	Ace   CardFace = 'A'
	Joker CardFace = 'X'
)

type CardSize byte

const (
	CardSizeTwo   CardSize = 0x02
	CardSizeThree CardSize = 0x03
	CardSizeFour  CardSize = 0x04
	CardSizeFive  CardSize = 0x05
	CardSizeSix   CardSize = 0x06
	CardSizeSeven CardSize = 0x07
	CardSizeEight CardSize = 0x08
	CardSizeNine  CardSize = 0x09
	CardSizeTen   CardSize = 0x0A
	CardSizeJack  CardSize = 0x0B
	CardSizeQueen CardSize = 0x0C
	CardSizeKing  CardSize = 0x0D
	CardSizeAce   CardSize = 0x0E
	CardSizeJoker CardSize = 0x10
)

var CardFace2SizeSlice []CardSize

func init() {
	CardFace2SizeSlice = make([]CardSize, 100)
	CardFace2SizeSlice[Two] = CardSizeTwo
	CardFace2SizeSlice[Three] = CardSizeThree
	CardFace2SizeSlice[Four] = CardSizeFour
	CardFace2SizeSlice[Five] = CardSizeFive
	CardFace2SizeSlice[Six] = CardSizeSix
	CardFace2SizeSlice[Seven] = CardSizeSeven
	CardFace2SizeSlice[Eight] = CardSizeEight
	CardFace2SizeSlice[Nine] = CardSizeNine
	CardFace2SizeSlice[Ten] = CardSizeTen
	CardFace2SizeSlice[Jack] = CardSizeJack
	CardFace2SizeSlice[Queen] = CardSizeQueen
	CardFace2SizeSlice[King] = CardSizeKing
	CardFace2SizeSlice[Ace] = CardSizeAce
	CardFace2SizeSlice[Joker] = CardSizeJoker
}

const (
	FiveGameType uint = iota + 1
	SevenGameType
	SevenAndJokerGameType
)
