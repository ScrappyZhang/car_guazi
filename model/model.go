package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"log"
	"routinego/car_guazi/spiders"
)

// mysql配置信息
var (
	DB *gorm.DB

	username string = "root"
	password string = "hitzzy"
	dbName   string = "ershouche"
)

//初始化model就连接数据库
func init() {
	var err error
	DB, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", username, password, dbName))
	if err != nil {
		log.Fatalf(" gorm.Open.err: %v", err)
	}

	DB.SingularTable(true)
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "gz_" + defaultTableName
	}
}

//向数据表中添加数据
func AddCars(cars []spiders.QcCar){
	for index, car := range cars {
		if err := DB.Create(&car).Error; err != nil {
			log.Printf("db.Create index: %s, err : %v", index, err)
		}
	}
}