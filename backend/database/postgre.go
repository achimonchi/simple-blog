package database

import (
	"database/sql"
	"fmt"
	"simple-blog/config"
	"time"

	_ "github.com/lib/pq"
)

func GetConnection(conf *config.ConfigDB) *sql.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=%s", conf.Host, conf.Port, conf.User, conf.Pass, conf.SSLMode)
	fmt.Println(dsn)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(conf.MaxIdleConns)
	db.SetMaxOpenConns(conf.MaxOpenConns)
	db.SetConnMaxLifetime(time.Duration(conf.MaxLifeTime) * time.Second)
	db.SetConnMaxIdleTime(time.Duration(conf.MaxIdleConns) * time.Second)

	return db
}
