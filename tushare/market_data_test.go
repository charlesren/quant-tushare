package tushare

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDaily(t *testing.T) {
	ast := assert.New(t)
	params := make(map[string]string)
	var fields []string
	// Check params
	_, err := client.Daily(params, fields)
	if err != nil {
		ast.Equal(err.Error(), "Need one argument ts_code or trade_date")
	}

	params["trade_date"] = "20181101"
	resp, err := client.Daily(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestDailyInvalidDateArgs(t *testing.T) {
	ast := assert.New(t)
	params := make(map[string]string)
	params["trade_date"] = "2018-11-01"
	var fields []string
	_, err := client.Daily(params, fields)

	if err != nil {
		ast.Equal(err.Error(), "please input right date format YYYYMMDD")
	}
}

func TestWeekly(t *testing.T) {
	ast := assert.New(t)
	params := make(map[string]string)
	var fields []string
	// Check params
	_, err := client.Weekly(params, fields)
	if err != nil {
		ast.Equal(err.Error(), "Need one argument ts_code or trade_date")
	}

	params["trade_date"] = "20181101"
	resp, err := client.Weekly(params, fields)

	if err != nil {
		t.Errorf("Api should not return an error, got: %s", err)
	}
	if resp == nil {
		t.Errorf("Api should return data")
	}
}

func TestWeeklyInvalidDateArgs(t *testing.T) {
	ast := assert.New(t)
	params := make(map[string]string)
	params["trade_date"] = "2018-11-01"
	var fields []string
	_, err := client.Weekly(params, fields)

	if err != nil {
		ast.Equal(err.Error(), "please input right date format YYYYMMDD")
	}
}
