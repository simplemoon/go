package main

import (
	"fmt"
	"time"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/clickhouse"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"

	// "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Config struct {
	UserName string // 用户名
	Password string // 密码
	Addr string // 地址
	Port int // 端口号
	DBName string // DB 的名称
}

func (c *Config) GetConnectStr() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", 
		c.UserName, c.Password, c.Addr, c.Port, c.DBName)
}

func (c *Config) GetPostgreSQLStr() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		c.Addr, c.UserName, c.Password, c.DBName, c.Port)
}

func (c *Config) GetSqlServerStr() string {
	return fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s", c.UserName, c.Password, c.Addr, c.Port, c.DBName)
}

func (c *Config) GetClickHouseStr() string {
	return fmt.Sprintf("tcp://%s:%d?database=%s&username=%s&password=%s&read_timeout=10&write_timeout=20", 
	c.Addr, c.Port, c.DBName, c.UserName, c.Password)
}



func ConnectMysql() (*gorm.DB, error) {
	m := Config{
		UserName: "root",
		Password: "root123456",
		Addr:     "localhost",
		Port:     3306,
		DBName:   "test",
	}

	dsn := m.GetConnectStr()
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                           dsn,
		SkipInitializeWithVersion:     false,
		DefaultStringSize:             256,
		DisableDatetimePrecision:      true,
		DontSupportRenameIndex:        true,
		DontSupportRenameColumn:       true,
	}), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, err
}

func ConnectPostgres() (*gorm.DB, error) {
	m := Config{
		UserName: "root",
		Password: "root123456",
		Addr:     "localhost",
		Port:     3306,
		DBName:   "test",
	}
	dsn := m.GetPostgreSQLStr()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func ConnectSqlite() (*gorm.DB, error) {
	// db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	return db, err
}

func ConncetSqlServer() (*gorm.DB, error) {
	m := Config{
		UserName: "root",
		Password: "root123456",
		Addr:     "localhost",
		Port:     3306,
		DBName:   "test",
	}
	dsn := m.GetSqlServerStr()
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	return db, err
}

func ConnectClickHouse() (*gorm.DB, error) {
	m := Config{
		UserName: "root",
		Password: "root123456",
		Addr:     "localhost",
		Port:     3306,
		DBName:   "test",
	}
	dsn := m.GetClickHouseStr()

	db, err := gorm.Open(clickhouse.Open(dsn), &gorm.Config{})
	return db, err
}

func SetConnectPool(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		return
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxIdleTime(time.Hour)
}

