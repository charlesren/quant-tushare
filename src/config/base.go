package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var dbConn *gorm.DB

func initDB() {
	dbUri := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=%s ", DBConfig.Host, DBConfig.DbPort, DBConfig.Dbname, DBConfig.Username, DBConfig.Password, DBConfig.Sslmode)
	//	fmt.Println("Connecting: " + dbUri)
	var err error
	dbConn, err = gorm.Open(DBConfig.Platform, dbUri)
	if err != nil {
		panic(err)
	}

	//    if config.ServerConfig.Env == DevelopmentMode { DB.LogMode(true) }
	dbConn.DB().SetMaxIdleConns(DBConfig.MaxIdleConns)
	dbConn.DB().SetMaxOpenConns(DBConfig.MaxOpenConns)
	//user singular table in case y change to ies
	dbConn.SingularTable(true)
	//	dbConn.Debug().AutoMigrate(&model.User{}, &model.Product{}, &model.Order{}, &model.OrderLine{})
	//	dbConn.Model(&model.OrderLine{}).AddForeignKey("order_id", "orders(order_id)", "CASCADE", "CASCADE")
	//	dbConn.Model(&model.OrderLine{}).AddForeignKey("product_id", "products(product_id)", "CASCADE", "CASCADE")
	//	dbConn.Model(&model.Order{}).AddForeignKey("user_id", "users(user_id)", "CASCADE", "CASCADE")
}

//
func GetDB() *gorm.DB { return dbConn }
func init() {
	initDB()
}
