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
	user := User{}
	q :=`INSERT INTO users (user_name, mail, password, bio) VALUES ($1, $2, $3, $4) RETURNING id`
	err := RequestReader(req, &user)
	if err != nil {
		fmt.Println("add user err", err)
		http.Error(w, "Bad Request Err", http.StatusBadRequest)
	}
	err = s.DB.QueryRow(q, &user.UserName, &user.Mail, &user.Password, &user.Bio).Scan(&user.ID)
	if err != nil {
		fmt.Println("query row err", err)
		http.Error(w, "Bad Request Err", http.StatusBadRequest)
	}
	PostResponseWriter(w, user)
}

func (s *service) GetAllUser(w http.ResponseWriter, req *http.Request) {
	q := `SELECT id, user_name, mail, no_of_post, bio, created_at FROM "users"`
	rows, err := s.DB.Query(q)
	if err != nil {
		http.Error(w, "Bad Request Err", http.StatusBadRequest)
	}
	defer rows.Close()
	user := []User{}
	for rows.Next() {
		temp := User{}
		err = rows.Scan(&temp.ID, &temp.UserName, &temp.Mail, &temp.NoOfPost, &temp.Bio, &temp.CreatedAt)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "DATABASE Err", http.StatusInternalServerError)
		}
		user = append(user, temp)
	}
	if err = rows.Err(); err != nil {
		fmt.Println(err)
		http.Error(w, "DATABASE Err", http.StatusInternalServerError)
	}
	GetResponseWriter(w, user)
	fmt.Println("done sending respones")
}