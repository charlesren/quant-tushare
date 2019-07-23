package main

import (
	"fmt"
	"quant-tushare/src/tushare"
)

func main() {
<<<<<<< Updated upstream
	//db := config.GetDB()
	//defer db.Close()
=======
	//	db := config.GetDB()
	//	defer db.Close()
>>>>>>> Stashed changes
	api := tushare.New("ef010210f12832692df4bd8e01f9353e794515a8d79bc0e7692026e7")
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
	fmt.Println(resp.Data.Items[0])
	fmt.Println(resp.Data.Items[0][0])
	fmt.Println(resp.Data.Fields)
	fmt.Println(resp.Data.Fields[0])
	var daily tushare.Daily
	daily.TsCode = "000002.SZ"
	daily.TradeDate = "20190708"
	daily.Open = 29.5
	//db.Create(&model.Daily{TsCode: daily.TsCode, TradeDate: daily.TradeDate, Open: daily.Open})
}
