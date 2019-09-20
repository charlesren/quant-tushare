package tushare

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/jinzhu/gorm"
)

// GetTushareData use http post methord to get data from https://tushare.pro
func (api *TuShare) GetTushareData(apiName string, params Params, fields Fields) (*APIResponse, error) {
	/*
		// Check params
		_, hasTsCode := params["ts_code"]
		_, hasTradeDate := params["trade_date"]

		// ts_code & trade_date required
		if (!hasTsCode && !hasTradeDate) || (hasTsCode && hasTradeDate) {
			return nil, fmt.Errorf("Need one argument ts_code or trade_date")
		}
	*/
	body := map[string]interface{}{
		"api_name": apiName,
		"token":    api.token,
		"fields":   fields,
		"params":   params,
	}
	return api.postData(body)
}

// ParsingTushareData save response of tushare  api  to slice
func ParsingTushareData(resp *APIResponse, dataTypeAddress interface{}, db *gorm.DB) {
	items := resp.Data.Items
	fields := resp.Data.Fields
	for i := 0; i < len(fields); i++ {
		fields[i] = SnakeToUpperCamel(fields[i])
	}

	dbdata := reflect.ValueOf(dataTypeAddress).Elem()
	iterType := reflect.TypeOf(dataTypeAddress).Elem().Elem()
	iterData := reflect.New(iterType)
	for _, value := range items {
		for i := 0; i < len(value); i++ {
			v := reflect.ValueOf(value[i])
			iterData.Elem().FieldByName(fields[i]).Set(v)
		}
		dbdata.Set(reflect.Append(dbdata, iterData.Elem()))
	}
	fmt.Println("Result of parsing tushare response data: ", dbdata)
}

//UpdateTradeCal function update trade calendar of SSE 、SZSE...
func UpdateTradeCal(db *gorm.DB, api *TuShare) {
	var checkPoint CheckPoint
	var stockExchange StockExchange
	stockExchange = SE
	for _, exchange := range stockExchange {
		checkPoint.Item = exchange
		if err := db.Select("day").Where("item = ?", exchange).Find(&checkPoint).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				checkPoint.Day = "19901219"
			}
		}
		startDate := checkPoint.Day
		endDate := time.Now().Format("20060102")
		params := make(Params)
		fields := Fields{}
		params["exchange"] = exchange
		params["start_date"] = startDate
		params["end_date"] = endDate
		if startDate == endDate {
			fmt.Printf("Trade calendar of %v is already up to date!!!\n", exchange)
		} else {
			if dateFormat := IsDateFormat(params["start_date"], params["end_date"]); !dateFormat {
				log.Fatal("please input right date format YYYYMMDD")
			}
			resp, err := api.GetTushareData("trade_cal", params, fields)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Response of tushare api: ", *resp)
			dataType := []TradeCal{}
			ParsingTushareData(resp, &dataType, db)
			// updata data
			for _, iterData := range dataType {
				if err := db.Find(&iterData).Error; err != nil {
					if err == gorm.ErrRecordNotFound {
						fmt.Printf("Updating %v\n", iterData)
						db.Create(&iterData)
					}
				}
			}
			// update checkPoint
			if err := db.Find(&checkPoint).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					checkPoint.Day = endDate
					db.Create(&checkPoint)
				}
			} else {
				db.Delete(&checkPoint)
				checkPoint.Day = endDate
				db.Create(&checkPoint)
			}
			fmt.Printf("Trade calendar of %v update successfully!!!\n", exchange)
		}
	}
}
