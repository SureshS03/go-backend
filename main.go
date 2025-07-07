package main

//TODO write a docker file with backend, DB, redis all (error in docker installations)
//docker , hosting, and kubunets ,cicd in github action
import (
	"fmt"
	"github.com/SureshS03/goconnect/internal/redis"
	"net/http"
	//"os"
)

func main() {
	//s := fmt.Sprintf("user=%v dbname=%v password=%v sslmode=%v", os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("PASSWORD"), os.Getenv("SSL_MODE"))
	//fmt.Println(s)
	db := NewDB("postgres", "host=db user=postgres dbname=goconnect password=arya sslmode=disable")
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
