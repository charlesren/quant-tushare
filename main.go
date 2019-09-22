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
	fields = tushare.APIFullFields["stock_basic"]
	resp, err := api.GetTushareData("stock_basic", params, fields)
	if err != nil {
		log.Fatal(err)
	}
	resp.TushareModelFields()
	tushare.UpdateTradeCal(db, api)
	tushare.UpdateStockBasic(db, api)
	tushare.UpdateDaily(db, api)
}
