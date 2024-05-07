package store

import (
	go_ora "github.com/sijms/go-ora/v2"
	"log"
	"testKafka/config"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
)

var RegDB *sqlx.DB

func GetDBConnection(cfg *config.Config) *sqlx.DB {
	server := cfg.DbServer
	port, err := strconv.Atoi(cfg.DbPort)
	if err != nil {
		log.Printf("error parsing DbPort: %v\n", err)
		return nil
	}
	service := cfg.DbService
	user := cfg.DbUser
	password := cfg.DbPassword
	urlOptions := map[string]string{}

	dbUrl := go_ora.BuildUrl(server, port, service, user, password, urlOptions)
	db, err := sqlx.Connect("oracle", dbUrl)
	if err != nil {
		log.Print(err)
		return nil
	}

	if db != nil {
		db.SetMaxOpenConns(20)
		db.SetMaxIdleConns(10)
		db.SetConnMaxLifetime(time.Duration(120 * time.Second))
	}

	log.Print("connect to SU DB!")
	return db
}

func GetRegDB(cfg *config.Config) *sqlx.DB {
	if RegDB != nil {
		err := RegDB.Ping()
		if err != nil {
			RegDB = GetDBConnection(cfg)
		}
	} else {
		RegDB = GetDBConnection(cfg)
	}
	return RegDB
}
