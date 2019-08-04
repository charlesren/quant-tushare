package tushare

import (
	"fmt"
	"reflect"
)

// GetDaily 获取股票行情数据, 日线行情
func (api *TuShare) GetDaily(params map[string]string, fields []string) (*APIResponse, error) {
	// Check params
	_, hasTsCode := params["ts_code"]
	_, hasTradeDate := params["trade_date"]

	// ts_code & trade_date required
	if (!hasTsCode && !hasTradeDate) || (hasTsCode && hasTradeDate) {
		return nil, fmt.Errorf("Need one argument ts_code or trade_date")
	}

	if dateFormat := IsDateFormat(params["trade_date"], params["start_date"], params["end_date"]); !dateFormat {
		return nil, fmt.Errorf("please input right date format YYYYMMDD")
	}

	body := map[string]interface{}{
		"api_name": "daily",
		"token":    api.token,
		"fields":   fields,
		"params":   params,
	}

	return api.postData(body)
}

// ParsingDaily  save response f tushare daily api  to []Daily slice
func (resp *APIResponse) ParsingDaily() []Daily {
	items := resp.Data.Items
	fields := resp.Data.Fields
	for i := 0; i < len(fields); i++ {
		fields[i] = SnakeToUpperCamel(fields[i])
	}
	var dbdata []Daily
	for _, value := range items {
		iterData := Daily{}
		for i := 0; i < len(value); i++ {
			v := reflect.ValueOf(value[i])
			reflect.ValueOf(&iterData).Elem().FieldByName(fields[i]).Set(v)
		}
		dbdata = append(dbdata, iterData)
	}
	return dbdata
}

// GetTradeCal get trade calendar of SSE or SZSE
func (api *TuShare) GetTradeCal(params map[string]string, fields []string) (*APIResponse, error) {
	// Check params

	if dateFormat := IsDateFormat(params["start_date"], params["end_date"]); !dateFormat {
		return nil, fmt.Errorf("please input right date format YYYYMMDD")
	}

	body := map[string]interface{}{
		"api_name": "trade_cal",
		"token":    api.token,
		"fields":   fields,
		"params":   params,
	}

	return api.postData(body)
}

// ParsingTradeCal save response f tushare trade_cal api  to []TradeCal slice
func (resp *APIResponse) ParsingTradeCal() []TradeCal {
	items := resp.Data.Items
	fields := resp.Data.Fields
	for i := 0; i < len(fields); i++ {
		fields[i] = SnakeToUpperCamel(fields[i])
	}
	var dbdata []TradeCal
	for _, value := range items {
		iterData := TradeCal{}
		for i := 0; i < len(value); i++ {
			v := reflect.ValueOf(value[i])
			reflect.ValueOf(&iterData).Elem().FieldByName(fields[i]).Set(v)
		}
		dbdata = append(dbdata, iterData)
	}
	return dbdata
}
