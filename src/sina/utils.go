package sina

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const sinaURLPrefix = "http://hq.sinajs.cn/list="

// GetData return real time stock data info
func GetData(stock string) Data {
	stock = TushareToSina(stock)
	URL := sinaURLPrefix + stock
	fmt.Println(URL)
	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	t := strings.Split(string(body), "\"")
	d := strings.Split(t[1], ",")
	var data Data
	data.Name = d[0]
	data.Open, _ = strconv.ParseFloat(d[1], 64)
	data.Close, _ = strconv.ParseFloat(d[2], 64)
	data.Now, _ = strconv.ParseFloat(d[3], 64)
	data.High, _ = strconv.ParseFloat(d[4], 64)
	data.Low, _ = strconv.ParseFloat(d[5], 64)
	data.Buy, _ = strconv.ParseFloat(d[6], 64)
	data.Sell, _ = strconv.ParseFloat(d[7], 64)
	data.Turnover, _ = strconv.Atoi(d[8])
	data.Volume, _ = strconv.ParseFloat(d[9], 64)
	data.Bid1Volume, _ = strconv.Atoi(d[10])
	data.Bid1, _ = strconv.ParseFloat(d[11], 64)
	data.Bid2Volume, _ = strconv.Atoi(d[12])
	data.Bid2, _ = strconv.ParseFloat(d[13], 64)
	data.Bid3Volume, _ = strconv.Atoi(d[14])
	data.Bid3, _ = strconv.ParseFloat(d[15], 64)
	data.Bid4Volume, _ = strconv.Atoi(d[16])
	data.Bid4, _ = strconv.ParseFloat(d[17], 64)
	data.Bid5Volume, _ = strconv.Atoi(d[18])
	data.Bid5, _ = strconv.ParseFloat(d[19], 64)
	data.Ask1Volume, _ = strconv.Atoi(d[20])
	data.Ask1, _ = strconv.ParseFloat(d[21], 64)
	data.Ask2Volume, _ = strconv.Atoi(d[22])
	data.Ask2, _ = strconv.ParseFloat(d[23], 64)
	data.Ask3Volume, _ = strconv.Atoi(d[24])
	data.Ask3, _ = strconv.ParseFloat(d[25], 64)
	data.Ask4Volume, _ = strconv.Atoi(d[26])
	data.Ask4, _ = strconv.ParseFloat(d[27], 64)
	data.Ask5Volume, _ = strconv.Atoi(d[28])
	data.Ask5, _ = strconv.ParseFloat(d[29], 64)
	data.Date = d[30]
	data.Time = d[31]
	return data
}

// TushareToSina convert Stock name from tushare format to Sina format
func TushareToSina(stock string) string {
	t := strings.Split(stock, ".")
	stock = strings.ToLower(t[1]) + t[0]
	return stock
}
