package main

import (
	"database/sql"
	"fmt"
	"net/http"
)

type service struct {
	DB *sql.DB
}

func NewService(db *sql.DB) *service {
	return &service{DB: db}
}

func (s *service) AddUser(w http.ResponseWriter, req *http.Request) {
	err := AuthChecker(w, req)
	if err != nil {
		http.Error(w, "bad token", 600)
		return
	}
	user := User{}
	q := `INSERT INTO users (username, mail, password, bio) VALUES ($1, $2, $3, $4) RETURNING id`
	err = RequestReader(req, &user)
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
	q := `SELECT id, username, mail, no_of_post, bio, created_at FROM "users"`
	err := AuthChecker(w, req)
	if err != nil {
		http.Error(w, "bad token", 600)
		return
	}
	rows, err := s.DB.Query(q)
	if err != nil {
		http.Error(w, "Bad Request Err", http.StatusBadRequest)
	}
	defer rows.Close()
	user := []User{}
	for rows.Next() {
		temp := &User{}
		err = rows.Scan(&temp.ID, &temp.UserName, &temp.Mail, &temp.NoOfPost, &temp.Bio, &temp.CreatedAt)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "DATABASE Err", http.StatusInternalServerError)
		}
		user = append(user, *temp)
	}
	if err = rows.Err(); err != nil {
		fmt.Println(err)
		http.Error(w, "DATABASE Err", http.StatusInternalServerError)
	}
	GetResponseWriter(w, user)
	fmt.Println("done sending respones")
}

func (s *service) GetUser(w http.ResponseWriter, req *http.Request) {
	err := AuthChecker(w, req)
	if err != nil {
		http.Error(w, "bad token", 600)
		return
	}
	id := GetParam(req, "id")
	user := &User{}
	q := `SELECT id, username, mail, no_of_post, bio, created_at FROM "users" WHERE id = ($1)`
	err = s.DB.QueryRow(q, id).Scan(&user.ID, &user.UserName, &user.Mail, &user.NoOfPost, &user.Bio, &user.CreatedAt)
	if err != nil {
		http.Error(w, "DATABASE Err", http.StatusInternalServerError)
		return
	}
	GetResponseWriter(w, user)
}

func (s *service) Addpost(w http.ResponseWriter, req *http.Request) {
	err := AuthChecker(w, req)
	if err != nil {
		http.Error(w, "bad token", 600)
		return
	}
	CreationPost := &CreationPost{}
	post := &Post{}
	err = RequestReader(req, CreationPost)
	if err != nil {
		fmt.Println(err)
	}
	q := `INSERT INTO posts (user_id, url) VALUES ($1, $2) RETURNING id`
	err = s.DB.QueryRow(q, &CreationPost.UserId, &CreationPost.URL).Scan(&post.ID)
	if err != nil {
		fmt.Println(err)
	}
	tx, err := s.DB.Begin()
	if err != nil {
		fmt.Println(err)
	}
	incq := `UPDATE users SET no_of_post = no_of_post + 1 WHERE id = $1`
	_, err = tx.Exec(incq, &CreationPost.UserId)
	if err != nil {
		tx.Rollback()
		fmt.Println(err)
	}
	err = tx.Commit()
	if err != nil {
		fmt.Println(err)
	}
	PostResponseWriter(w, CreationPost)
}