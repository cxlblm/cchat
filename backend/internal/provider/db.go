package provider

import (
	"cchart/internal/kernel"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB(config *kernel.Config) *gorm.DB {
	c := config.Database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", c.User, c.Password, c.Host, c.Port, c.DB)

	queryString := ""
	if c.Charset != "" {
		queryString += "charset=" + c.Charset
	} else {
		queryString += "charset=utf8mb4"
	}
	if c.ParseTime {
		queryString += "&parseTime=True"
	} else {
		queryString += "&parseTime=False"
	}
	if c.Loc != "" {
		queryString += "&loc=" + c.Loc
	} else {
		queryString += "&loc=Local"
	}

	dsn += "?" + queryString
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db
}
