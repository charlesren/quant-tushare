package main

import (
	"fmt"
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
	//today := time.Now().Format("20060102")
	//params["ts_code"] = "000002.SZ"
	//params["start_date"] = "20190707"
	//params["end_date"] = today
	fields = tushare.Fields{"ts_code", "symbol", "name", "area", "industry", "fullname", "enname", "market", "exchange", "curr_type", "list_status", "list_date", "delist_date", "is_hs"}
	//
	resp, err := api.GetTushareData("stock_basic", params, fields)
	if err != nil {
		log.Fatal(err)
	}
	resp.TushareModelFields()
	var daily tushare.Daily
	//db.Select("trade_date, open").Find(&daily)
	db.Where("trade_date = ?", "20190708").Find(&daily)
	//db.Select("ts_code").Where("trade_date = ?", "20190708").Find(&daily)
	fmt.Printf("%v\n", daily)
	// Update trade calendar
	tushare.UpdateTradeCal(db, api)
}
