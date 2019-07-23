package tushare

import (
	"github.com/jinzhu/gorm"
)

type Daily struct {
	gorm.Model
	TsCode    string
	TradeDate string
	Open      float64
	High      float64
	Low       float64
	Close     float64
	PreClose  float64
	Change    float64
	PctChg    float64
	Vol       float64
	Amount    float64
}

type Weekly struct {
	gorm.Model
	TsCode    string
	TradeDate string
	Close     float64
	Open      float64
	High      float64
	Low       float64
	PreClose  float64
	Change    float64
	PctChg    float64
	Vol       float64
	Amount    float64
}

type Monthly struct {
	gorm.Model
	TsCode    string
	TradeDate string
	Close     float64
	Open      float64
	High      float64
	Low       float64
	PreClose  float64
	Change    float64
	PctChg    float64
	Vol       float64
	Amount    float64
}
