package golang_database

import (
	"database/sql"
	"fmt"
	"time"
)

const (
	HOST     = "localhost"
	PORT     = 5432
	USER     = "postgres"
	PASSWORD = "postgres"
	DBNAME   = "belajar_golang_database"
)

func GetConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", HOST, PORT, USER, PASSWORD, DBNAME)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(10)                 // jumlah koneksi yang dibuat diawal
	db.SetMaxIdleConns(100)                // jumlah maksimal koneksi yang dibuat, jika ada request yang membutuhkan koneksi maka diharuskan untuk menunggu salah satu koneksi selesai
	db.SetConnMaxIdleTime(5 * time.Minute) // jika koneksi diam saja tidak melakukan apa apa selama waktu yang ditentukan maka akan ditutup
	db.SetConnMaxLifetime(1 * time.Hour)   // lama koneksi tersedia, jika telah sampai maka akan di close dan dibuat ulang koneksi yang baru
	return db
}
