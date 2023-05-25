package logic

import "poker/model"

type iComparer interface {
	JudgmentCardType(card string) (model.CardType, map[model.CardFace]int, byte)
}
