package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/carefree/project/common/db"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/yaml.v2"
)

func main() {
	var cfg db.Config
	yamlFile, err := ioutil.ReadFile("project/home/cfg.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, &cfg)
	if err != nil {
		fmt.Println(err.Error())
	}
	str := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", cfg.Username, cfg.Password, cfg.DBName)
	dbs, err := gorm.Open("mysql", str)
	if err != nil {
		log.Fatal(err)
	}
	defer dbs.Close()

	dbs = dbs.Table(cfg.DefaultTable)
	dbs.DropTable(&db.Row{})
	dbs.CreateTable(&db.Row{})
}
