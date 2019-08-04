package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DBConfig save database information
var DBConfig struct {
	Charset         string `json:"charset"`
	MaxIdleConns    int    `json:"maxidleconns"`
	MaxOpenConns    int    `json:"maxopenconns"`
	ConnMaxLifetime int64  `json:"connmaxlifetime"`
	Sslmode         string `json:"sslmode"`
	Platform        string `json:"platform"`
	Host            string `json:"host"`
	DbPort          int    `json:"dbport"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	Dbname          string `json:"dbname"`
	BindPort        int    `json:"bindport"`
}

var dbConn *gorm.DB

func initDB() {
	dbcfg, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(dbcfg, &DBConfig); err != nil {
		fmt.Println(err)
		return
	}
	dbURL := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=%s ", DBConfig.Host, DBConfig.DbPort, DBConfig.Dbname, DBConfig.Username, DBConfig.Password, DBConfig.Sslmode)
	//fmt.Println("Connecting: " + dbUri)
	dbConn, err = gorm.Open(DBConfig.Platform, dbURL)
	if err != nil {
		panic(err)
	}

	//    if config.ServerConfig.Env == DevelopmentMode { DB.LogMode(true) }
	dbConn.DB().SetMaxIdleConns(DBConfig.MaxIdleConns)
	dbConn.DB().SetMaxOpenConns(DBConfig.MaxOpenConns)
	//user singular table in case y change to ies
	dbConn.SingularTable(true)
	//dbConn.AutoMigrate(&tushare.Daily{}, &tushare.Weekly{}, &tushare.TradeCal{}, &tushare.CheckPoint{})
	//	dbConn.Model(&model.OrderLine{}).AddForeignKey("order_id", "orders(order_id)", "CASCADE", "CASCADE")
	//	dbConn.Model(&model.OrderLine{}).AddForeignKey("product_id", "products(product_id)", "CASCADE", "CASCADE")
	//	dbConn.Model(&model.Order{}).AddForeignKey("user_id", "users(user_id)", "CASCADE", "CASCADE")
}

// GetDB  return dbconn
func GetDB() *gorm.DB { return dbConn }
func init() {
	initDB()
}
