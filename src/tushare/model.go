package tushare

type Daily struct {
	TsCode    string `gorm:"primary_key"`
	TradeDate string `gorm:"primary_key"`
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
	TsCode    string `gorm:"primary_key"`
	TradeDate string `gorm:"primary_key"`
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
	TsCode    string `gorm:"primary_key"`
	TradeDate string `gorm:"primary_key"`
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
