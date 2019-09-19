package main

import (
	"log"
	"quant-tushare/src/config"
	"quant-tushare/src/tushare"
)

func main() {
	db := config.GetDB()
	defer db.Close()
	api := tushare.New(tushare.TushareConfig.Token)
	params := make(tushare.Params)
	var fields tushare.Fields
	fields = tushare.Fields{"ts_code", "symbol", "name", "area", "industry", "fullname", "enname", "market", "exchange", "curr_type", "list_status", "list_date", "delist_date", "is_hs"}
	resp, err := api.GetTushareData("stock_basic", params, fields)
	if err != nil {
		log.Fatal(err)
	}
	resp.TushareModelFields()
	// Update trade calendar
	tushare.UpdateTradeCal(db, api)
}
