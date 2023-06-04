package seven

import (
	"github.com/stretchr/testify/assert"
	"poker/logic"
	"poker/model"
	"testing"
)

func BenchmarkPokerMan(b *testing.B) {
	s := logic.GetPoker(model.SevenGameType)
	for i := 0; i < b.N; i++ {
		s.PokerMan()
	}
}

func BenchmarkSevenCom_JudgmentCardType(b *testing.B) {
	s := sevenCom{}
	for i := 0; i < b.N; i++ {
		s.JudgmentCardType("5dQh4h3h7c6d8c")
	}
}

// 381.1 ns/op
// 385.7 ns/op
// 385.1 ns/op
func BenchmarkSplit(b *testing.B) {
	s := &logic.BaseCardCom{}
	card := "5dQh4h3h7c6d8c"
	for i := 0; i < b.N; i++ {
		s.CardsSplitMapCount(card)
	}
}

func TestPokerMan(t *testing.T) {
	s := logic.GetPoker(model.SevenGameType)
	s.PokerMan()
}

func TestCompare(t *testing.T) {
	s := logic.GetPoker(model.SevenGameType)
	// 55k3tq4
	// hsdshcc
	alice := "5h5sKd3sThQc4c"
	// q855k3t
	// shhsdsh
	bob := "Qs8h5h5sKd3sTh"
	res := s.Compare(alice, bob)
	assert.Equal(t, res, model.Draw)
	// 同为顺子比较
	// Ah6c4d2s5h3c7h
	// a642537->765432a
	// 3d9dAh6c4d2s5h
	// 39a6425->965432a
	alice = "Ah6c4d2s5h3c7h"
	bob = "3d9dAh6c4d2s5h"
	res = s.Compare(alice, bob)
	assert.Equal(t, res, model.Win)

	// 同为顺子比较
	// 5dQh4h3h7c6d8c
	// 5q43768
	// dhhhcdc
	// 6sKc5dQh4h3h7c
	// 65kQ437

	// file:1 my:0
	// Qd7cAhKsKd7sKh
	// Q7AKK7K -> KKK77 QA
	// 5hKcQd7cAhKsKd
	// 5KQ7AKK -> KKK QA57
	alice = "Qd7cAhKsKd7sKh"
	bob = "5hKcQd7cAhKsKd"
	res = s.Compare(alice, bob)
	assert.Equal(t, res, model.Win)

	// file:0 my:2
	// 5h5sKd3sThQc4c
	// 55K3TQ4 55 KTQ34
	// hsdshcc
	// Qs8h5h5sKd3sTh
	// Q855K3T 55 KTQ83
	// shhsdsh
	alice = "5h5sKd3sThQc4c"
	bob = "Qs8h5h5sKd3sTh"
	res = s.Compare(alice, bob)
	assert.Equal(t, res, model.Draw)

	// file:2 my:1
	// 8hThJh6hQhAs3d
	// hhhhh sd -> QJT86
	// 8TJ6QA3 ->AQJT863
	// Kh3c8hThJh6hQh
	// hhhhh hc 3 KQJT8
	// K38TJ6Q  ->KQJT863
	alice = "8hThJh6hQhAs3d"
	bob = "Kh3c8hThJh6hQh"
	res = s.Compare(alice, bob)
	assert.Equal(t, res, model.Lose)

	// file:0 my:1
	// Qd8d2d7cKc4h5h
	// Q827K45 -> KQ875 42
	// dddcchh
	// 5s3sQd8d2d7cKc
	// 53Q827K ->KQ875 32
	alice = "Qd8d2d7cKc4h5h"
	bob = "5s3sQd8d2d7cKc"
	res = s.Compare(alice, bob)
	assert.Equal(t, res, model.Draw)
	// file:0 my:2
	// 9cTdTcTh4s3dQd
	// 9TTT43Q-> TTT Q9 43
	// Qh6d9cTdTcTh4s
	// Q69TTT4  TTT Q9 64
	alice = "9cTdTcTh4s3dQd"
	bob = "Qh6d9cTdTcTh4s"
	res = s.Compare(alice, bob)
	assert.Equal(t, res, model.Draw)

	// file:0 my:2
	// QsQdKdKcTh2hKh
	// QQKKT2K KKKQQ T2
	// AsKsQsQdKdKcTh
	// AKQQKKT KKKQQ AT

	// file:2 my:1
	// JhJd6cAdJs3h2c
	// JJ6AJ32 JJJ A6 32
	// 7hTdJhJd6cAdJs
	// 7TJJ6AJ JJJ AT 67

	// file:2 my:1
	// Ts9s2cKsJs7d8s
	// T92KJ78 -> T9KJ8
	// sscssds
	// 6dAsTs9s2cKsJs
	// 6AT92KJ
	// dssscss -> AT9KJ
	alice = "Ts9s2cKsJs7d8s"
	bob = "6dAsTs9s2cKsJs"
	res = s.Compare(alice, bob)
	assert.Equal(t, res, model.Lose)

	// file:0 my:2
	// 6dQdQc2hQh6s2s
	// 6QQ2Q62 QQQ 66 22
	// 7d6h6dQdQc2hQh
	// 766QQ2Q QQQ 66 27
	alice = "6dQdQc2hQh6s2s"
	bob = "7d6h6dQdQc2hQh"
	res = s.Compare(alice, bob)
	assert.Equal(t, model.Draw, res)

	// file:0 my:1
	// TsTc4d4cTdAh3h
	// TT44TA3  TTT44 A3
	// 2c4hTsTc4d4cTd
	// 24TT44T 	T444T T2
	alice = "TsTc4d4cTdAh3h"
	bob = "2c4hTsTc4d4cTd"
	res = s.Compare(alice, bob)
	assert.Equal(t, model.Draw, res)
	// file:0 my:1
	// QdQc6d8d8c5h6c
	// QQ68856 QQ88 665
	// 6s2cQdQc6d8d8c
	// 62QQ688 QQ88 6662

	alice = "QdQc6d8d8c5h6c"
	bob = "6s2cQdQc6d8d8c"
	res = s.Compare(alice, bob)
	assert.Equal(t, model.Draw, res)
}

func TestFullHouseCom(t *testing.T) {
	b := logic.BaseCardCom{
		CardSizeMap1: map[model.CardSize]int{
			model.CardSizeTen:   3,
			model.CardSizeFour:  2,
			model.CardSizeAce:   1,
			model.CardSizeThree: 1,
		},
		CardSizeMap2: map[model.CardSize]int{
			model.CardSizeTen:  3,
			model.CardSizeFour: 3,
			model.CardSizeTwo:  1,
		},
		Max1:       0,
		Max2:       0,
		CardsSize1: nil,
		CardsSize2: nil,
	}
	res := b.FullHouseCom()
	assert.Equal(t, model.Draw, res)
}
