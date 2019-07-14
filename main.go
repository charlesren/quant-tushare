package main

import (
	"fmt"
	"quant-tushare/src/config"
	"quant-tushare/src/tushare"
)

func main() {
	dbconn := config.GetDB()
	defer dbconn.Close()
	api := tushare.New("jklj;lj;j;j;jkl;j;jj")
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
	fmt.Println(resp)
}
