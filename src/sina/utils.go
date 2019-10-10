package sina

import (
	"strings"
)

const sinaURLPrefix = "http://hq.sinajs.cn/list"

// GetData return real time stock data info
func GetData(sotck string) Data {
	var data Data
	return data
}

// TushareToSina convert Stock name from tushare format to Sina format
func TushareToSina(stock string) string {
	t := strings.Split(stock, ".")
	stock = strings.ToLower(t[1]) + t[0]
	return stock
}
