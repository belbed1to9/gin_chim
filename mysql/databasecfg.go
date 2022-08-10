package mysql

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
)

var DBmap = initDb("private/databasecfg.json")

type DataBaseConfig struct {
	User     string `json:"USER"`
	Passwd   string `json:"PASSWD"`
	Host     string `json:"HOST"`
	Port     string `json:"PORT"`
	Database string `json:"DB"`
}

func (d DataBaseConfig) String() string {
	return fmt.Sprintf("%v:%v(%v:%v)/%v", d.User, d.Passwd, d.Host, d.Port, d.Database)
}

func processConfig(dbCfgPath string) (DataBaseConfig, error) {
	var c DataBaseConfig
	file, err := os.Open(dbCfgPath)
	if err != nil {
		return c, err
	}
	byteVal, _ := ioutil.ReadAll(file)

	err = json.Unmarshal(byteVal, &c)
	return c, err
}

func initDb(dbCfgPath string) *gorp.DbMap {
	d, err := processConfig(dbCfgPath)
	if err != nil {
		log.Fatal(err)
	}
	db, err := sql.Open("mysql", d.String())
	CheckErr(err, "sql.Open failed")
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	err = dbmap.CreateTablesIfNotExists()
	CheckErr(err, "Create tables failed")
	return dbmap
}
