package main

import (
	"log"
	"github.com/charlesren/quant-tushare/src/config"
	"github.com/charlesren/quant-tushare/src/tushare"
)

func main() {
	//init database
	db := config.GetDB()
	defer db.Close()
	//init tushare account
	api := tushare.New(tushare.TushareConfig.Token)
	//define api params
	params := make(tushare.Params)
	//define api response  field
	var fields tushare.Fields
	fields = tushare.APIFullFields["stock_basic"]
	//request date from tushare
	resp, err := api.GetTushareData("stock_basic", params, fields)
	if err != nil {
		log.Fatal(err)
	}
	//print response fields
	resp.TushareModelFields()
	//request data and save into db
	tushare.UpdateTradeCal(db, api)
	tushare.UpdateStockBasic(db, api)
	tushare.UpdateDaily(db, api)
}
