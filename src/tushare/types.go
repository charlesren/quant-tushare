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

// APIFullFields store tushare api name and full fields of api response
var APIFullFields = map[string]Fields{
	"trade_cal":   {"exchange", "cal_date", "is_open", "pretrade_date"},
	"stock_basic": {"ts_code", "symbol", "name", "area", "industry", "fullname", "enname", "market", "exchange", "curr_type", "list_status", "list_date", "delist_date", "is_hs"},
	"daily":       {"ts_code", "trade_date", "open", "high", "low", "close", "pre_close", "change", "pct_chg", "vol", "amount"},
	"weekly":      {"ts_code", "trade_date", "close", "open", "high", "low", "pre_close", "change", "pct_chg", "vol", "amount"},
	"monthly":     {"ts_code", "trade_date", "close", "open", "high", "low", "pre_close", "change", "pct_chg", "vol", "amount"},
	"daily_basic": {"ts_code", "trade_date", "close", "turnover_rate", "turnover_rate_f", "volume_ratio", "pe", "pe_ttm", "pb", "ps", "ps_ttm", "total_share", "float_share", "free_share", "total_mv", "circ_mv"},
	"stk_limit":   {"trade_date", "ts_code", "pre_close", "up_limit", "down_limit"},
	"stk_account": {"date", "weekly_new", "total", "weekly_hold", "weekly_trade"},
}
