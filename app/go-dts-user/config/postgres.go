package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PostgresConfig struct {
	// minimum config untuk postgres
	Port     uint   `mapstructure:"port"`
	Host     string `mapstructure:"host"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbName"`

	// extended best practice
	MaxIdleConnection int `mapstructure:"maxIdleConnection"`
	MaxOpenConnection int `mapstructure:"maxOpenConnection"`
	MaxIdleTime       int `mapstructure:"maxIdleTime"`
}

func NewPostgresConn() (db *sql.DB) {
	db, err := sql.Open("postgres", postgresDSN())
	if err != nil {
		panic(err)
	}
	// set extended config
	postgresPoolConf(db)

	// defer close dilakukan
	// supaya apps menutup koneksi
	// ketika apps dimatikan
	// defer db.Close()

	// test connection
	if err := db.Ping(); err != nil {
		panic(err)
	}
	return
}

func NewPostgresGormConn() (db *gorm.DB) {
	// connect ke database menggunakan lib gorm
	// https://gorm.io/docs/query.html

	// connect to db
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	db, err := gorm.Open(postgres.Open(postgresDSN()), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	dbSQL, err := db.DB()
	if err != nil {
		panic(err)
	}
	postgresPoolConf(dbSQL)

	if err := dbSQL.Ping(); err != nil {
		panic(err)
	}
	log.Println("successfully connect to Postgres")
	return db
}

func postgresDSN() string {
	return fmt.Sprintf(`
		host=%v
		port=%v
		user=%v
		password=%v
		dbname=%v
		sslmode=disable
		application_name=%v
	`,
		Load.DataSource.Postgres.Master.Host,
		Load.DataSource.Postgres.Master.Port,
		Load.DataSource.Postgres.Master.Username,
		Load.DataSource.Postgres.Master.Password,
		Load.DataSource.Postgres.Master.DBName,
		Load.Server.Name,
	)
}

func postgresPoolConf(dbSQL *sql.DB) {
	// set extended config
	dbSQL.SetMaxIdleConns(Load.DataSource.Postgres.Master.MaxIdleConnection)
	dbSQL.SetMaxOpenConns(Load.DataSource.Postgres.Master.MaxOpenConnection)
	dbSQL.SetConnMaxIdleTime(time.Duration(Load.DataSource.Postgres.Master.MaxIdleTime))
}
