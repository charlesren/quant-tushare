package main

import (
	"fmt"
	"quant-tushare/src/config"
	"quant-tushare/src/tushare"

	"github.com/jinzhu/gorm"
)

func main() {
	db := config.GetDB()
	defer db.Close()
	api := tushare.New(tushare.TushareConfig.Token)
	params := make(map[string]string)
	var fields []string
	//params["trade_date"] = "20190708"
	params["ts_code"] = "000002.SZ"
	params["start_date"] = "20190707"
	params["end_date"] = "20190711"
	resp, err := api.GetDaily(params, fields)
	if err != nil {
		err.Error()
	}
	fmt.Println(*resp)
	for _, v := range resp.ParsingDaily() {
		fmt.Println(v)
		if err := db.Find(&v).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				db.Create(&v)
			} else {
				fmt.Println(err)
			}
		}
	}
	var daily tushare.Daily
	//db.Select("trade_date, open").Find(&daily)
	db.Where("trade_date = ?", "20190708").Find(&daily)
	//db.Select("ts_code").Where("trade_date = ?", "20190708").Find(&daily)
	fmt.Printf("%v\n", daily)
	//tushare.UpdateTradeCal(db)
}
