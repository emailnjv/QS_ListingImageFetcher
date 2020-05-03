package db

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)
// TargetDBClient is a instance of the database client
type TargetDBClient struct {
	db *gorm.DB
}

// NewDBClient returns a new instance of a database client
func NewTargetDBClient() (TargetDBClient, error) {
	var result TargetDBClient

	// Load ENVs
	err := godotenv.Load("../local.env")
	if err != nil {
		return result, err
	}

	mysqlDB := os.Getenv("MYSQL_DATABASE")
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlHost := os.Getenv("MYSQL_HOST")

	// Create connection string
	connectionString := mysqlUser + ":" + mysqlPassword + "@" + mysqlHost + "/" + mysqlDB + "?charset=utf8&parseTime=True&loc=Local"

	// Open connection
	result.db, err = gorm.Open("mysql",  connectionString)
	if err != nil {
		return result, err
	}

	// Migrate the schema
	result.db.AutoMigrate(&Product{})

	return result, err
}

// CloseConnection closes the client's connection
// It returns an error if one occurred
func (d TargetDBClient) CloseConnection() error {
	return d.db.Close()
}