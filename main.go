package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool("debug") {
		fmt.Println("API runs on DEBUG mode")
	}
}

func initDbConnection() *sql.DB {
	var db *sql.DB
	var err error
	cs := fmt.Sprintf(`
	host=%s
	 port=%s
	 user=%s
	 password=%s
	 dbname=%s
	 sslmode=require`,
		viper.GetString("database.host"),
		viper.GetString("database.port"),
		viper.GetString("database.user"),
		viper.GetString("database.password"),
		viper.GetString("database.name"))

	db, err = sql.Open("postgres", cs)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database")
	return db
}

func main() {
	db := initDbConnection()
	defer db.Close()

	port := viper.GetString("server.port")
	fmt.Println("Start listening on port " + port)
	log.Fatal(http.ListenAndServe("localhost:"+port, nil))
}
