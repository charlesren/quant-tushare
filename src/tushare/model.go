package tushare

// Daily struct store return data of tushare daily api
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

// TradeCal struct store return data of tushare trade_cal api
type TradeCal struct {
	Exchange     string `gorm:"primary_key"`
	CalDate      string `gorm:"primary_key"`
	IsOpen       string
	PretradeDate string
}

// Weekly struct store return data of tushare weekly api
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

// Monthly struct store return data of tushare monthly api
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

// CheckPoint struct store checkpoint like stock daily trade data etc
type CheckPoint struct {
	Item string `gorm:"primary_key"`
	Day  string `gorm:"primary_key"`
}
