package component

import (
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
  "fmt"
  config "github.com/spf13/viper"
)

type DBHandler struct{}

func GetDBHandler() *DBHandler {
	return new(DBHandler)
}

func (this *DBHandler) DB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(this.getDbConnectionString()), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func (db *DBHandler) getDbConnectionString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", 
		config.GetString("database.mysql.user"), 
		config.GetString("database.mysql.password"),
		config.GetString("database.mysql.host"), 
		config.GetString("database.mysql.port"), 
		config.GetString("database.mysql.name"),
	)
}
