package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/spf13/viper"
)

var (
	Con *gorm.DB
)

func Connect() error {
	dbType := viper.GetString("database.type")
	switch dbType {
	case "mysql":
		return connectMySQL()
	case "postgres":
		return connectPostgres()
	case "sqlite":
		return connectSQLite()
	default:
		return InvalidDatabaseTypeError{dbType}
	}
}

func Close() error {
	return Con.Close()
}

func connectMySQL() error {
	dbHost := viper.GetString("database.host")
	dbPort := viper.GetInt("database.port")
	dbName := viper.GetString("database.name")
	dbUser := viper.GetString("database.user")
	dbPass := viper.GetString("database.password")

	conString := fmt.Sprintf(
		"%s:%s@%s:%s/%s?charset=utf8&parseTime=True&loc=Local",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName,
	)

	con, err := gorm.Open("mysql", conString)
	if err != nil {
		return err
	}

	Con = con
	return nil
}

func connectPostgres() error {
	dbHost := viper.GetString("database.host")
	dbPort := viper.GetInt("database.port")
	dbName := viper.GetString("database.name")
	dbUser := viper.GetString("database.user")
	dbPass := viper.GetString("database.password")

	conString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s",
		dbHost,
		dbPort,
		dbUser,
		dbName,
		dbPass,
	)
	con, err := gorm.Open("postgres", conString)
	if err != nil {
		return err
	}

	Con = con
	return nil
}

func connectSQLite() error {
	dbPath := viper.GetString("database.path")

	con, err := gorm.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}

	Con = con
	return nil
}

func Migrate() {
	Con.AutoMigrate(
		&User{},
		&Device{},
		&Password{},
	)
}
