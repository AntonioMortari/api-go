package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)


func Connection() (*sql.DB, error) {
	err := godotenv.Load()
	if(err != nil){
		panic(err)
	}

	var (
		host     = os.Getenv("HOST")
		port     = os.Getenv("PORT")
		dbname   = os.Getenv("DBNAME")
		password = os.Getenv("PASSWORD")
		user     = os.Getenv("USER")
	)
	
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("ERRO AO SE CONECTAR COM BANCO DE DADOS: ", err)
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db, nil
}
