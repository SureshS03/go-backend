package main

import (
	"fmt"
	"net/http"
	"github.com/SureshS03/goconnect/backend/redis"
	//"os"
)

func main() {
	db := NewDB("postgres", "postgres://postgres:arya@postgres:5432/goconnect?sslmode=disable")
	if db == nil {
		fmt.Println("DB connection failed. Exiting...")
		return
	}
	redis.Init()
	defer db.Close()
	router := NewRouter(db)
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	fmt.Println(server.Addr)
	fmt.Println("server on 8080")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(" main err", err)
	}
}
