package conn

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func connectMysql(user, pwd, addr, dbName string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		pwd,
		addr,
		dbName,
	)
	fmt.Println("dsn:", dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDb, err := db.DB()
	if err != nil {
		return nil, err
	}

	err = sqlDb.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func getConfig() (*gorm.DB, error) {
	viper.SetDefault("mysql.user", "testuser")
	viper.SetDefault("mysql.password", "testpwd")
	viper.SetDefault("mysql.addr", "localhost:3306")
	viper.SetDefault("mysql.database", "testdb")

	return connectMysql(
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.addr"),
		viper.GetString("mysql.database"),
	)
}

func CheckConnect() *gorm.DB {
	db, err := getConfig()
	if err != nil {
		return nil
	}
	return db
}
