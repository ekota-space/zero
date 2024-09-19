package ql // Query Layer

import "database/sql"

var ql *queryLayer

type queryLayer struct {
	DB *sql.DB
}

func InitLayer(db *sql.DB) {
	if ql != nil {
		panic("QueryLayer is already initialized")
	}

	ql = &queryLayer{
		DB: db,
	}
}

func GetDB() *sql.DB {
	if ql == nil {
		panic("QueryLayer is not initialized")
	}
	return ql.DB
}
