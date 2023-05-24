package logic

import (
	"fmt"
	"poker/model"
)

// SizeTranByte 对面值转译整对应的byte值  -- 方便大小比较
func SizeTranByte(card model.CardFace) byte {
	res, ok := model.CardFaceMap[card]
	if !ok {
		panic(fmt.Sprintf("unkown card face =%+v", card))
	}
	return res
}
