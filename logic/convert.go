package logic

import (
	"poker/model"
)

// SizeTranByte 对面值转译整对应的byte值  -- 方便大小比较
// 19.93
func SizeTranByte(card model.CardFace) model.CardSize {
	return model.CardFaceMap[card]
}
