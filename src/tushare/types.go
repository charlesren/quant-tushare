package tushare

// APIResponse tushare api response
type APIResponse struct {
	RequestID string      `json:"request_id"`
	Code      int         `json:"code"`
	Msg       interface{} `json:"msg"`
	Data      struct {
		Fields []string        `json:"fields"`
		Items  [][]interface{} `json:"items"`
	} `json:"data"`
}

// Params store input params used by tushare http api
type Params map[string]string

// Fields define return fields of tushare http api
type Fields []string

// StockExchange define stock exchange such as SSE、SZSE
type StockExchange []string

// SE is dedault stock exchange
var SE = StockExchange{"SSE", "SZSE"}
