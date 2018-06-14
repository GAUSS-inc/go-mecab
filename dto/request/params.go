// Copyright 2018 GAUSS All Rights Reserved.
//

package request

// MecabParam Mecab実行パラメータ
type MecabParam struct {
	Sentence string `json:"sentence" query:"sentence" validate:"required"` // 対象文字列
}
