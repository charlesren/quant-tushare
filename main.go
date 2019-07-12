package main

import (
	"fmt"
	"quant-tushare/src/config"
	"quant-tushare/src/tushare"
)

func main() {
	dbconn := config.GetDB()
	fmt.Print(dbconn)
	params := make(map[string]string)
	params["trade_date"] = "20190708"
	var fields []string
	var api *tushare.TuShare
	var resp tushare.APIResponse
	resp, err := api.Daily(params, fields)
	if err != nil {
		err.Error()
	}

	resp.r2j
}

func (r tushare.APIResponse) r2j() (string, error) {
	d = r.Data
	f = d.Fields
	fmt.Printf(f)
}
