package main

import (
	"fmt"
	"net/http"
)

func main() {
	db := NewDB("postgres", "user=postgres dbname=goconnect password=arya sslmode=disable")
	defer db.Close()
	router := NewRouter(db)
	server := http.Server{
		Addr: ":8080",
		Handler: router,
	}
	fmt.Println("server on 8080")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}