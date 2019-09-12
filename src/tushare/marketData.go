package tushare

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/jinzhu/gorm"
)

// GetDaily 获取股票行情数据, 日线行情
func (api *TuShare) GetDaily(params map[string]string, fields []string) (*APIResponse, error) {
	// Check params
	_, hasTsCode := params["ts_code"]
	_, hasTradeDate := params["trade_date"]

	// ts_code & trade_date required
	if (!hasTsCode && !hasTradeDate) || (hasTsCode && hasTradeDate) {
		return nil, fmt.Errorf("Need one argument ts_code or trade_date")
	}

	if dateFormat := IsDateFormat(params["trade_date"], params["start_date"], params["end_date"]); !dateFormat {
		return nil, fmt.Errorf("please input right date format YYYYMMDD")
	}

	body := map[string]interface{}{
		"api_name": "daily",
		"token":    api.token,
		"fields":   fields,
		"params":   params,
	}

	return api.postData(body)
}

// ParsingDaily  save response f tushare daily api  to []Daily slice
func (resp *APIResponse) ParsingDaily() []Daily {
	items := resp.Data.Items
	fields := resp.Data.Fields
	for i := 0; i < len(fields); i++ {
		fields[i] = SnakeToUpperCamel(fields[i])
	}
	var dbdata []Daily
	for _, value := range items {
		iterData := Daily{}
		for i := 0; i < len(value); i++ {
			v := reflect.ValueOf(value[i])
			reflect.ValueOf(&iterData).Elem().FieldByName(fields[i]).Set(v)
		}
		dbdata = append(dbdata, iterData)
	}
	return dbdata
}

// GetTradeCal get trade calendar of SSE or SZSE
func (api *TuShare) GetTradeCal(params map[string]string, fields []string) (*APIResponse, error) {
	// Check params

	if dateFormat := IsDateFormat(params["start_date"], params["end_date"]); !dateFormat {
		return nil, fmt.Errorf("please input right date format YYYYMMDD")
	}

	body := map[string]interface{}{
		"api_name": "trade_cal",
		"token":    api.token,
		"fields":   fields,
		"params":   params,
	}

	return api.postData(body)
}

// ParsingTradeCal save response f tushare trade_cal api  to []TradeCal slice
func (resp *APIResponse) ParsingTradeCal() []TradeCal {
	items := resp.Data.Items
	fields := resp.Data.Fields
	for i := 0; i < len(fields); i++ {
		fields[i] = SnakeToUpperCamel(fields[i])
	}
	var dbdata []TradeCal
	for _, value := range items {
		iterData := TradeCal{}
		for i := 0; i < len(value); i++ {
			v := reflect.ValueOf(value[i])
			reflect.ValueOf(&iterData).Elem().FieldByName(fields[i]).Set(v)
		}
		dbdata = append(dbdata, iterData)
	}
	return dbdata
}

//UpdateTradeCal function update trade calendar of SSE 、SZSE...
func UpdateTradeCal(db *gorm.DB, api *TuShare) {
	var checkPoint CheckPoint
	var StockExchange []string
	StockExchange = []string{"SSE", "SZSE"}
	for _, exchange := range StockExchange {
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
			resp, err := api.GetTradeCal(params, fields)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(*resp)
			for _, v := range resp.ParsingTradeCal() {
				if err := db.Find(&v).Error; err != nil {
					if err == gorm.ErrRecordNotFound {
						db.Create(&v)
					}
				}
			}
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
