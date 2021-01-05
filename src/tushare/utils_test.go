package tushare

import (
	"fmt"
	"testing"
)

func TestNextDay(t *testing.T) {
	date := "20191001"
	nextDay := NextDay(date)
	if nextDay == "20191002" {
		fmt.Println("Good!!!")
	} else {
		t.Errorf("Wrong date!!!\n")
	}
}
func TestTushareModelFields(t *testing.T) {
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
}