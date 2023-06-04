package util

import (
	"poker/model"
	"sort"
)

// QuickSort 系统排序函数 倒序
func QuickSort(cards []model.CardSize) {
	sort.Slice(cards, func(i, j int) bool {
		return cards[i] > cards[j]
	})
}
