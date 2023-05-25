package logic

import "testing"

func BenchmarkPokerMain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PokerMan()
	}
}

// 306.3 ns/op
// 300.0 ns/op
// 314.4 ns/op
func BenchmarkCardsSplitMapCount(b *testing.B) {
	cards := "AhKhQhJhTh"
	for i := 0; i < b.N; i++ {
		CardsSplitMapCount(cards)
	}
}

// 25.91 ns/op
// 25.58 ns/op
// 25.50 ns/op
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
