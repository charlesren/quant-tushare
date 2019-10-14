package tushare

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"
)

// IsDateFormat Check date format YYYYMMDD
func IsDateFormat(dates ...string) bool {
	re := regexp.MustCompile(`^\d{4}\d{2}\d{2}$`)
	for _, date := range dates {
		if date == "" {
			continue
		}
		if ret := re.MatchString(date); !ret {
			return false
		}

	}
	return true
}

// SnakeToUpperCamel translate snake-case to upper camel-case
// For instance :  abc_de to AbcDe
func SnakeToUpperCamel(snake string) (upperCamel string) {
	isToUpper := true
	for _, runeValue := range snake {
		if isToUpper {
			upperCamel += strings.ToUpper(string(runeValue))
			isToUpper = false
		} else {
			if runeValue == '_' {
				isToUpper = true
			} else {
				upperCamel += string(runeValue)
			}
		}
	}
	return
}

// TushareModelFields print http response fileds name and type
// this func could be helpful when define  gorm model
func (resp *APIResponse) TushareModelFields() {
	fields := resp.Data.Fields
	for i := 0; i < len(fields); i++ {
		fields[i] = SnakeToUpperCamel(fields[i])
		fmt.Println(fields[i], reflect.TypeOf(fields[i]))
	}
}

// NextDay return date string in format "20160102"
// eg: input "20191001"  return "20191002"
func NextDay(data string) string {
	format := "20060102"
	t, _ := time.Parse(format, data)
	n := t.AddDate(0, 0, 1)
	nextDay := n.Format(format)
	return nextDay
}
