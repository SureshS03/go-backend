package main

import (
	"database/sql"
	"fmt"
	"time"
	_ "github.com/lib/pq"
)

func NewDB(dn string, ds string) *sql.DB {
	for i:=0; i <= 10; i++ {

		db, err := sql.Open(dn, ds)
		if err == nil {
			err = db.Ping()
			if err == nil {
				fmt.Println("Connected to DB successfully")
				return db
			}
		}
		fmt.Println(err)
		fmt.Println("Not Connecnted attempt",i)
		time.Sleep(2 * time.Second)
	}
	return nil
}
