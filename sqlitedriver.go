package ck_orm

import (
	"database/sql"
	"fmt"
	//_ "github.com/mattn/go-sqlite3"
)

var SqliteDrivers = make(map[string]*sql.DB)

func InitSqliteDb(conf *DBConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"file:%s?cache=shared",
		conf.DBName,
	)

	if db, ok := SqliteDrivers[dsn]; ok {
		return db, nil
	}

	sqliteDb, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}

	//sqliteDb.SetMaxOpenConns(conf.DBPoolSize)
	//sqliteDb.SetMaxIdleConns(conf.DBIdleSize)
	sqliteDb.SetMaxOpenConns(1)
	err = sqliteDb.Ping()
	if err != nil {
		return nil, err
	}

	SqliteDrivers[dsn] = sqliteDb

	return sqliteDb, nil
}
