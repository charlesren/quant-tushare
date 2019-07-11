package config

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
)


var DBConfig struct {
    Charset      string  `json:"charset"`
    MaxIdleConns int    `json:"maxidleconns"`
    MaxOpenConns int   `json:"maxopenconns"`
    ConnMaxLifetime int64 `json:"connmaxlifetime"`
    Sslmode         string `json:"sslmode"`
    Platform string `json:"platform"`
	Host     string `json:"host"`
	DbPort   int    `json:"dbport"`
	Username string `json:"username"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
	BindPort int    `json:"bindport"`
}
func initDBConfig() {
    dbcfg, err := ioutil.ReadFile("./config.json")
    if err != nil {
        panic(err) 
    }

    if err := json.Unmarshal(dbcfg, &DBConfig); err != nil {
        fmt.Println(err)
       return
    }
}


func init() {
    initDBConfig()
}