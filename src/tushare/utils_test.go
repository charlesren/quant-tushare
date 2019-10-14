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
