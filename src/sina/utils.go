package sina

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const sinaURLPrefix = "http://hq.sinajs.cn/list="

// GetData return real time stock data info
func GetData(stock string) Data {
	stock = TushareToSina(stock)
	URL := sinaURLPrefix + stock
	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	t := strings.TrimSuffix(string(body), ";")
	d := strings.Split(t, "\"")
	var data *Data
	err = json.Unmarshal([]byte(d[1]), &data)
	return *data
}

// TushareToSina convert Stock name from tushare format to Sina format
func TushareToSina(stock string) string {
	t := strings.Split(stock, ".")
	stock = strings.ToLower(t[1]) + t[0]
	return stock
}
