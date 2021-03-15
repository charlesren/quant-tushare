package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
)

// CHConfig save clickhouse config information used for DSN
var CHConfig struct {
	Database   string  `json:"database"`
	Host  string `json:"host"`
	Port int `json:"port"`
	Username  string `json:"username"`
	Password  string  `json:"password"`
	NoDelay  bool `json:"no_delay"`
	ConnectionOpenStrategy  string `json:"connection_open_strategy"`
	BlockSize  int64 `json:"block_size"`
	PoolSize  int  `json:"pool_size"`
	Debug  bool   `json:"debug"`
	Secure   bool  `json:"secure"`
	Skip_verify bool `json:"skip_verify"`
	Tls_config  string  `json:"tls_config"`
	AltHosts string   `json:"alt_hosts"`
	Compress int  `json:"compress"`
	ReadTimeout int  `json:"read_timeout"`
	WriteTimeout  int `json:"write_timeout"`
}

var dbConn *gorm.DB

func initDB() {
	dbcfg, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(dbcfg, &CHConfig); err != nil {
		fmt.Println(err)
		return
	}
	dsn := fmt.Sprintf("tcp://%s:%d?database=%s&username=%s&password=%s&read_timeout=%d&write_timeout=%d", CHConfig.Host, CHConfig.Port, CHConfig.Database, CHConfig.Username, CHConfig.Password, CHConfig.ReadTimeout, CHConfig.WriteTimeout)
	//fmt.Println("Connecting: " + dsn)
	dbConn, err = gorm.Open(clickhouse.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println(dbConn)
}

// GetDB  return dbconn
func GetDB() *gorm.DB { return dbConn }
func init() {
	initDB()
}
