package tushare

import (
	"regexp"
	"strings"
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
