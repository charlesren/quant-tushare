package tushare

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"gorm.io/gorm"
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
			if v.IsValid() {
				iterData.Elem().FieldByName(fields[i]).Set(v)
			}
		}
		dbdata.Set(reflect.Append(dbdata, iterData.Elem()))
	}
}

//UpdateTradeCal function update trade calendar of SSE 、SZSE...
func UpdateTradeCal(db *gorm.DB, api *TuShare) {
	var stockExchange StockExchange
	stockExchange = SE
	for _, exchange := range stockExchange {
		var tradeCal TradeCal
		fmt.Printf("Get last trade calendar day saved for %v !!!\n", exchange)
		if err := db.Table("trade_cals").Where("exchange = ?", exchange).Order("cal_date desc").Limit(1).Find(&tradeCal).Error; err == nil {
			if tradeCal.CalDate == "" {
				fmt.Printf("No trade calendar record found in db for %v !!!\n", exchange)
				fmt.Printf("Set last trade calendar day for %v to default !!!\n", exchange)
				tradeCal.CalDate = "19901219"
			}
			fmt.Printf("Last trade calendar day for %v is %v !!!\n", exchange, tradeCal.CalDate)
		} else {
			fmt.Printf("Get last trade calendar day saved for %v error !!!\n", exchange)
		}
		var startDate string
		endDate := time.Now().Format("20060102")
		params := make(Params)
		fields := APIFullFields["trade_cal"]
		params["exchange"] = exchange
		params["end_date"] = endDate
		if tradeCal.CalDate == endDate {
			fmt.Printf("Trade calendar of %v is already up to date!!!\n", exchange)
			continue
		} else {
			startDate = NextDay(tradeCal.CalDate)
		}
		params["start_date"] = startDate
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
			fmt.Printf("Start to update trade calendar of %v .\n", exchange)
			db.Create(dataType)
			fmt.Printf("Trade calendar of %v update successfully!!!\n", exchange)
		}
	}
}

// UpdateDaily update stock daily
func UpdateDaily(db *gorm.DB, api *TuShare) {
	params := make(Params)
	endDate := time.Now().Format("20060102")
	params["end_date"] = endDate
	//fields := APIFullFields["daily"]
	stockList := []StockBasic{}
	if err := db.Table("stock_basics").Find(&stockList).Error; err != nil {
		log.Fatal("Get stock basic data failed : ", err)
	}
	fmt.Printf("stockList is: %v !!!\n", stockList)
	var checkPoints []CheckPoint
	if err := db.Table("check_points").Find(&checkPoints).Error; err != nil {
		fmt.Println("Get checkpoint data failed !!!")
	}
	fmt.Printf("checkPoints is: %v !!!\n", checkPoints)
	/*
	for _, stock := range stockList {
		params["start_date"] = "19901219" //reset params["start_date"] to default start date
		var checkPoint CheckPoint
		params["ts_code"] = stock.TsCode
		flag := 0
		for i := 0; i < len(checkPoints); i++ {
			if stock.TsCode == checkPoints[i].Item {
				checkPoint = checkPoints[i]
				checkPointDay := checkPoints[i].Day
				if checkPointDay == params["end_date"] {
					params["start_date"] = checkPointDay
				} else {
					params["start_date"] = NextDay(checkPointDay)
				}
				checkPoints = append(checkPoints[:i], checkPoints[i+1:]...)
				flag = 1
				break
			}
		}
		if params["start_date"] == params["end_date"] {
			fmt.Printf("Daily data of %v is already up to date!!!\n", stock.TsCode)
		} else {
			resp, err := api.GetTushareData("daily", params, fields)
			if err != nil {
				log.Fatal(err)
			}
			respData := []Daily{}
			ParsingTushareData(resp, &respData, db)
			fmt.Printf("Response data for %v is : %v\n", stock.TsCode, respData) // updata data
			for _, iterData := range respData {
				fmt.Printf("Updating %v\n", iterData)
				db.Create(&iterData)
			}
			// update checkPoint
			lastDay := Daily{}
			fmt.Printf("Get last daily data for %v !!!\n", stock.TsCode)
			if err := db.Table("daily").Limit(1).Where("ts_code = ?", stock.TsCode).Order("trade_date desc").Find(&lastDay).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					fmt.Println("No checkpoint data found in db!!!")
					checkPoint.Day = "19901219"
				}
			}
			fmt.Printf("Last daily data for %v is : %v\n", stock.TsCode, lastDay)
			if flag == 1 {
				db.Delete(&checkPoint)
				checkPoint.Day = lastDay.TradeDate
				db.Create(&checkPoint)
				log.Printf("Checkpoint: %v update successfully!!!\n", checkPoint)
			} else {
				checkPoint.Item = stock.TsCode
				db.Create(&checkPoint)
				log.Printf("Checkpoint: %v create successfully!!!\n", checkPoint)
			}
		}
	}
	*/
}

// UpdateStockBasic update stock list
func UpdateStockBasic(db *gorm.DB, api *TuShare) {
	params := make(Params)
	fields := APIFullFields["stock_basic"]
	resp, err := api.GetTushareData("stock_basic", params, fields)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println("Response of tushare api: ", *resp)
	respData := []StockBasic{}
	ParsingTushareData(resp, &respData, db)
	existData := []StockBasic{}
	if err := db.Find(&existData).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("No data found in db!!!")
		}
	}
	for _, data := range respData {
		flag := 0
		for i := 0; i < len(existData); i++ {
			if existData[i].TsCode == data.TsCode {
				if existData[i] == data {
					existData = append(existData[:i], existData[i+1:]...)
					flag = 1
					break
				} else {
					db.Delete(existData[i])
				}
			}
		}
		if flag == 1 {
			fmt.Printf("%v already exist!!!\n", data)
		} else {
			fmt.Printf("Updating %v\n", data)
			db.Create(&data)
		}
	}
}
