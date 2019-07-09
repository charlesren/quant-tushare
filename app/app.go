package main

import (
	"fmt"
	"quant/config"
)

func main() {
	dbconn := config.GetDB()
	fmt.Print(dbconn)
	params := make(map[string]string)
	var fields []string
	_, err := tushare.Daily(params, fields)
	if err != nil {
	ast.Equal(err.Error(), "Need one argument ts_code or trade_date")
	}

	params["trade_date"] = "20181101"
	resp, err := tushare.Daily(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
	resp.r2j
}

func (r *APIResponse) r2j() (string, error){
   d = r.Data
   i = d.Items
   f = d.Fields
   fmt.Printf (i)
   fmt.Printf (f)
}
