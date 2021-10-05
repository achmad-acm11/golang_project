package app

import (
	"database/sql"
	"golang-restful-api/helper"
)

func NewDB() *sql.DB {
	// db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/golang_database")
	db, err := sql.Open("mysql", "umawgr4x9dzi1uwj:BBJssHpsuq8s7S0Q2SZc@tcp(bozmzpyy01mewtqol8jj-mysql.services.clever-cloud.com:3306)/bozmzpyy01mewtqol8jj")
	helper.PanicIfError(err)

	// db.SetMaxIdleConns(5)
	// db.SetMaxOpenConns(20)
	// db.SetConnMaxLifetime(60 * time.Minute)
	// db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
