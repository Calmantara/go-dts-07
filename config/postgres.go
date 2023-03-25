package config

import (
	"database/sql"
	"fmt"
	"time"
)

type PostgresConfig struct {
	// minimum config untuk postgres
	Port     uint
	Host     string
	Username string
	Password string
	DBName   string
	// extended best practice
	// param ini memang digunakan
	// untuk membatasi
	// apps connect ke DB
	// kita batasi supaya apps kita
	// tetap berada di performa yang stable

	// kalau open > 7, akan kill connection
	// yang kelebihannya

	// kalau ada banyak request
	// untuk akses data, apakah
	// beberapa user tidak akan mendapatkan data?
	// 1. dia akan tetap mendapatkan data, tapi sedikit nunggu
	// 2. ketika kita sudah deploy (e.g Kubernetes), bisa dengan aman
	// melakukan AUTO SCALING service kita

	MaxIdleConnection int
	MaxOpenConnection int
	MaxIdleTime       int
}

func NewPostgresConn() (db *sql.DB) {
	pgConf := PostgresConfig{
		Port:              25432,
		Host:              "127.0.0.1",
		Username:          "postgres",
		Password:          "mysecretpassword",
		DBName:            "user_management",
		MaxOpenConnection: 7,
		MaxIdleConnection: 5,
		MaxIdleTime:       int(30 * time.Minute),
	}

	connString := fmt.Sprintf(`
		host=%v
		port=%v
		user=%v
		password=%v
		dbname=%v
		sslmode=disable
	`,
		pgConf.Host,
		pgConf.Port,
		pgConf.Username,
		pgConf.Password,
		pgConf.DBName,
	)
	db, err := sql.Open("postgres", connString)
	if err != nil {
		panic(err)
	}
	// set extended config
	db.SetMaxIdleConns(pgConf.MaxIdleConnection)
	db.SetMaxOpenConns(pgConf.MaxOpenConnection)
	db.SetConnMaxIdleTime(time.Duration(pgConf.MaxIdleTime))

	// defer close dilakukan
	// supaya apps menutup koneksi
	// ketika apps dimatikan
	// defer db.Close()

	// test connection
	if err := db.Ping(); err != nil {
		panic(err)
	}

	// tidak menyarankan
	// untuk melakukan DDL
	// di apps

	// DDL dilakukan
	// biasanya manual / di CI/CD pipeline
	return
}
