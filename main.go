package main

import (
	"fmt"
	"quant-tushare/src/config"
	"quant-tushare/src/model"
	"quant-tushare/src/tushare"
)

func main() {
	db := config.GetDB()
	defer db.Close()
	api := tushare.New("jljlj;ljl;jl;jl;j;j;j;klj")
	params := make(map[string]string)
	//params["trade_date"] = "20190708"
	params["ts_code"] = "000002.SZ"
	params["start_date"] = "20190707"
	params["end_date"] = "20190708"
	var fields []string
	resp, err := api.Daily(params, fields)
	if err != nil {
		err.Error()
	}
	fmt.Println(resp.Data)
	fmt.Println(resp.Data.Items)
	fmt.Println(resp.Data.Fields)
	var daily model.Daily
	daily.TsCode = "000002.SZ"
	daily.TradeDate = "20190708"
	daily.Open = 29.5
	db.Create(&model.Daily{TsCode: daily.TsCode, TradeDate: daily.TradeDate, Open: daily.Open})
}
