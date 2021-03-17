package main

import (
	"fmt"
	"github.com/charlesren/quant-tushare/src/config"
	"github.com/charlesren/quant-tushare/src/tushare"
)

func main() {
	//init database
	db := config.GetDB()
	//init tushare account
	api := tushare.New(tushare.TushareConfig.Token)
	//update trade calendar of stock exchange
	fmt.Println("Start to update trade calendar")
	tushare.UpdateTradeCal(db, api)
	fmt.Println("Update trade calendar finished")
	//update stock list of stock exchange (sse/szse)
	fmt.Println("Start to update stock list")
	tushare.UpdateStockBasic(db, api)
	fmt.Println("Update stock list finished")
	//update daily exchange data
	fmt.Println("Start to update stock daily data")
	tushare.UpdateDaily(db, api)
	fmt.Println("Update daily data finished")

}
