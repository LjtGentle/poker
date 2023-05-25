package logic

import (
	"poker/model"
)

// SizeTranByte 对面值转译整对应的byte值  -- 方便大小比较
// 19.93
func SizeTranByte(card model.CardFace) model.CardSize {
	//res, ok := model.CardFaceMap[card]
	//if !ok {
	//	panic(fmt.Sprintf("unkown card face =%+v", card))
	//}
	//return res

	return model.CardFaceMap[card]

}
