package db

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mysql struct {
	DB_Host     string
	DB_Port     string
	DB_Name     string
	DB_User     string
	DB_Password string
}

func (m *Mysql) Dsn() string {
	return m.DB_User + ":" + m.DB_Password + "@tcp(" + m.DB_Host + ":" + m.DB_Port + ")/" + m.DB_Name + "?charset=utf8mb4&parseTime=True&loc=Local"
}

func (m *Mysql) Dialector() gorm.Dialector {
	return mysql.Open(m.Dsn())
}

func (m *Mysql) NewConnInfo() IDb {
	m = &Mysql{
		DB_Host:     os.Getenv("MYSQL_HOST"),
		DB_Port:     os.Getenv("MYSQL_PORT"),
		DB_Name:     os.Getenv("MYSQL_NAME"),
		DB_User:     os.Getenv("MYSQL_USER"),
		DB_Password: os.Getenv("MYSQL_PASSWORD"),
	}

	return m
}
