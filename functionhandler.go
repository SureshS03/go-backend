package main

import (
	"database/sql"
	"fmt"
	"net/http"
)

type service struct{
	DB *sql.DB
}

func NewService(db *sql.DB) *service{
	return &service{DB: db}
}

func (s *service) AddUser(w http.ResponseWriter, req *http.Request) {
	
}

func (s *service) GetAllUser(w http.ResponseWriter, req *http.Request) {
	q := `SELECT id user_name no_of_post bio FROM "users"`
	rows, err := s.DB.Query(q)
	if err != nil {
		panic(err)
	}
	fmt.Println(rows)
	defer rows.Close()
}