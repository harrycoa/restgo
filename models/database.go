package models

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"../config"
)

var db *sql.DB


func CreateConnection() {
	url := config.GetUrlDatabase()
	if connection, err := sql.Open("mysql", url); err != nil {
		panic(err)
	} else {
		db = connection
		fmt.Println("conexion exitosa")
	}
}

func CreateTables() {
	createTable("users", userSchema)
}

func createTable(tableName, schema string) {
	if !existTable(tableName) {
		Exec(schema)
	} else {
		truncateTable(tableName)
	}
}

func truncateTable(tableName string) {
	sql := fmt.Sprintf("TRUNCATE %s", tableName)
	Exec(sql)
}

func existTable(tableName string) bool {
	sql2 := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
	rows, _ := Query(sql2)
	return rows.Next()
}

func Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := db.Exec(query, args...)
	if err != nil {
		log.Println(err)
	}
	return result, err
}

func Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		log.Println(err)
	}
	return rows, err
}

func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}

func Closeconnection() {
	db.Close()
}


//<username>:<password>@tcp(<hot>:<port>)/<database>
