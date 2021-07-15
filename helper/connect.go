package helper

import (
	"fmt"
	"log"
	"os"
	"gorm.io/driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func Connection() (*gorm.DB, error) {
	// Load env
	err := godotenv.Load()
	if err != nil {
		log.Println("There is something wrong")
		return nil, err
	}

	//Identification env on variable
	db_user := os.Getenv("DBUSER")
	db_pass := os.Getenv("DBPASS")
	db_name := os.Getenv("DBNAME")
	db_port := os.Getenv("DBPORT")
	db_host := os.Getenv("DBHOST")

	//set connection database
	conn_string := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", db_user, db_pass, db_host, db_port, db_name)
	conn, err := gorm.Open(mysql.Open(conn_string), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	//return the result
	return conn, nil
}

func CloseConnection(db *gorm.DB) error{
	dbSQL, err := db.DB();
	if err != nil{
		return err
	}
	dbSQL.Close()
	return nil
}
