package five

import (
	"poker/dao"
	"testing"
)

func BenchmarkRead(b *testing.B) {
	file := dao.GetCurrentAbPathByCaller()
	file += "/resources/match_result.json"
	for i := 0; i < b.N; i++ {
		dao.ReadFile(file)
	}
}

func BenchmarkPokerMain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PokerMan()
	}
}

// windows
// 1249 ns/op
// 1278 ns/op
// 1170 ns/op
// 顺子 vs 双对
func BenchmarkCompare(b *testing.B) {
	c1 := "3h6c5dAsQd"
	c2 := "Qs7d7hQcTs"
	for i := 0; i < b.N; i++ {
		compare(c1, c2)
	}
}

// windows
// 565.0 ns/op
// 552.0 ns/op
// 548.1 ns/op
// mac
// 355.0 ns/op
// 353.1 ns/op
// 354.0 ns/op
func BenchmarkJudgmentGroupNew_Flush_C1(b *testing.B) {
	c1 := "3h6c5dAsQd"
	for i := 0; i < b.N; i++ {
		JudgmentGroupNew(c1)
	}
}

// windows
// 634.1 ns/
// 654.4 ns/
// 642.5 ns/
func BenchmarkJudgmentGroupNew_TwoPair_C2(b *testing.B) {
	c2 := "Qs7d7hQcTs"
	for i := 0; i < b.N; i++ {
		JudgmentGroupNew(c2)
	}
}

// mac
// 323.0 ns/op
// 321.2 ns/op
// 302.1 ns/op
// 320.3 ns/op
// 方法2
// 302.2 ns/op
// 323.3 ns/op
// 295.5 ns/op
// 300.8 ns/op

// windows
// 372.3 ns/op
// 380.4 ns/op
// 378.1 ns/op
func BenchmarkCardsSplitMapCount(b *testing.B) {
	cards := "AhKhQhJhTh"
	for i := 0; i < b.N; i++ {
		CardsSplitMapCount(cards)
	}
}

// windows
// 18.39 ns/op
// 19.80 ns/op
// 19.58 ns/op
//func BenchmarkSizeTranByte(b *testing.B) {
//	card := 'A'
//	for i := 0; i < b.N; i++ {
//		SizeTranByte(model.CardFace(card))
//	}
//}

// mac
// 25.91 ns/op
// 25.58 ns/op
// 25.50 ns/op

// windows
// 37.06 ns/op
// 37.94 ns/op
// 38.13 ns/op
// 方法2
// 32.22 ns/op
// 33.90 ns/op
// 33.29 ns/op
func BenchmarkIsShunZiNew_true(b *testing.B) {
	cards := "AhKhQhJhTh"
	sizeMap, _, cardsSize := CardsSplitMapCount(cards)
	for i := 0; i < b.N; i++ {
		IsShunZiNew(cardsSize, sizeMap)
	}
}

// 32.21 ns/op
// 31.63 ns/op
// 31.65 ns/op

// windows
// 46.40 ns/op
// 48.25 ns/op
// 46.58 ns/op
// 方法2
// 34.30 ns/op
// 36.09 ns/op
// 33.28 ns/op
func BenchmarkIsShunZiNew_false(b *testing.B) {
	cards := "6hKhQhJhTh"
	sizeMap, _, cardsSize := CardsSplitMapCount(cards)
	for i := 0; i < b.N; i++ {
		IsShunZiNew(cardsSize, sizeMap)
	}
}

// 1 皇家同花顺
// 359.8 ns/op
// 386.1 ns/op
// 372.9 ns/op
func BenchmarkJudgmentGroupNew_RoyalFlush(b *testing.B) {
	cards := "AhKhQhJhTh"

	for i := 0; i < b.N; i++ {
		JudgmentGroupNew(cards)
	}
}

// 2 同花顺
// 357.8 ns/op
// 359.4 ns/op
// 363.6 ns/op
func BenchmarkJudgmentGroupNew_StraightFlush(b *testing.B) {
	cards := "KhQhJhTh9h"

	for i := 0; i < b.N; i++ {
		JudgmentGroupNew(cards)
	}
}

// 3 四条
// 419.8 ns/op
// 427.3 ns/op
// 423.7 ns/op
func BenchmarkJudgmentGroupNew_FourHouse(b *testing.B) {
	cards := "KhKhKhKd9c"

	for i := 0; i < b.N; i++ {
		JudgmentGroupNew(cards)
	}
}

// 4 3带2
// 428.2 ns/op
// 430.3 ns/op
// 432.8 ns/op
func BenchmarkJudgmentGroupNew_FullHouse(b *testing.B) {
	cards := "KhKhKhQdQc"

	for i := 0; i < b.N; i++ {
		JudgmentGroupNew(cards)
	}
}

// 5 同花
// 377.1 ns/op
// 385.1 ns/op
// 379.5 ns/op
func BenchmarkJudgmentGroupNew_Flush(b *testing.B) {
	cards := "Ah3hKhQhQh"

	for i := 0; i < b.N; i++ {
		JudgmentGroupNew(cards)
	}
}
