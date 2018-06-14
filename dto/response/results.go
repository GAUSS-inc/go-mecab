// Copyright 2018 GAUSS All Rights Reserved.
//

package response

// MecabResults 形態素解析結果レスポンス
type MecabResults struct {
	Results []*MecabResult `json:"results" mapstructure:"results"` // 形態素解析結果
}

// MecabResult 形態素解析結果
// 形態素解析結果
type MecabResult struct {
	Surface        string `json:"surface" mapstructure:"surface"`                             // 表層形
	Pos            string `json:"pos" mapstructure:"pos" feature:"0"`                         // 品詞
	PosDetail1     string `json:"pos_detail1" mapstructure:"pos_detail1" feature:"1"`         // 品詞細分類1
	PosDetail2     string `json:"pos_detail2" mapstructure:"pos_detail2" feature:"2"`         // 品詞細分類2
	PosDetail3     string `json:"pos_detail3" mapstructure:"pos_detail3" feature:"3"`         // 品詞細分類3
	ConjugatedType string `json:"conjugated_type" mapstructure:"conjugated_type" feature:"4"` // 活用型
	ConjugatedForm string `json:"conjugated_form" mapstructure:"conjugated_form" feature:"5"` // 活用形
	Baseform       string `json:"baseform" mapstructure:"baseform" feature:"6"`               // 基本形
	Reading        string `json:"reading" mapstructure:"reading" feature:"7"`                 // 読み
	Pronunciation  string `json:"pronunciation" mapstructure:"pronunciation" feature:"8"`     // 発音
	Custom1        string `json:"custom1" mapstructure:"custom1" feature:"9"`                 // カスタム1
	Custom2        string `json:"custom2" mapstructure:"custom2" feature:"10"`                // カスタム2
	Custom3        string `json:"custom3" mapstructure:"custom3" feature:"11"`                // カスタム3
}
