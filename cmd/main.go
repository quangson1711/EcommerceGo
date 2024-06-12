package main

import (
	"Ecommerce-Go/cmd/api"
	"Ecommerce-Go/config"
	db2 "Ecommerce-Go/db"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
)

func main() {

	//envs := config.InitConfig()
	config.SetupLogger()

	db, err := db2.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewApiServer(":8080", db)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully connected to database")
}
