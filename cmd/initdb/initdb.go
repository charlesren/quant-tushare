package main

import (
	"fmt"
	"github.com/charlesren/quant-tushare/src/config"
	"github.com/charlesren/quant-tushare/src/tushare"
)

func main() {
	// connect to database
	db := config.GetDB()
	defer db.Close()
	// migrate schema
	db.AutoMigrate(&tushare.TradeCal{})
	db.AutoMigrate(&tushare.CheckPoint{})
	db.AutoMigrate(&tushare.StockBasic{})
	db.AutoMigrate(&tushare.Daily{})
	db.AutoMigrate(&tushare.Weekly{})
	db.AutoMigrate(&tushare.Monthly{})
	fmt.Println("migrate schema successfully!!!")
}