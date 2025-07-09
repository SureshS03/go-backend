package main

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
	"github.com/SureshS03/goconnect/backend/redis"
)

type service struct {
	DB *sql.DB
}

func NewService(db *sql.DB) *service {
	return &service{DB: db}
}

func (s *service) GenToken(id string) (string, error) {
	t := make([]byte, 20)
	if _, err := rand.Read(t); err != nil {
		fmt.Println("err in get token", err)
		return "", err
	}
	tEncode := hex.EncodeToString(t)
	err := redis.SetCache("token:"+tEncode, id, time.Minute*10)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return tEncode, nil
}

func (s *service) UserLogin(w http.ResponseWriter, req *http.Request) {
	user := &LoginUser{}
	err := RequestReader(req, user)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	q := `SELECT id FROM users WHERE username = ($1) AND password = ($2)`
	var id *int
	err = s.DB.QueryRow(q, &user.UserName, &user.Password).Scan(&id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Bad Input", http.StatusBadRequest)
		return
	}
	strid := strconv.Itoa(*id)
	token, err := s.GenToken(strid)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	PostResponseWriter(w, token)
}

func (s *service) AddUser(w http.ResponseWriter, req *http.Request) {
	user := &User{}
	q := `INSERT INTO users (username, mail, password, bio) VALUES ($1, $2, $3, $4) RETURNING id`
	err := RequestReader(req, user)
	if err != nil {
		fmt.Println("add user err", err)
		http.Error(w, "Bad Request Err", http.StatusBadRequest)
		return
	}
	if strings.Contains(user.Mail, "@") && strings.Contains(user.Mail, ".") {
		fmt.Println("invaild mail")
		http.Error(w, "Invalid Mail", http.StatusNonAuthoritativeInfo)
		return
	}
	err = s.DB.QueryRow(q, &user.UserName, &user.Mail, &user.Password, &user.Bio).Scan(&user.ID)
	if err != nil {
		fmt.Println("query row err", err)
		http.Error(w, "Bad Request Err", http.StatusBadRequest)
		return
	}
	err = SetUserCache(*user, time.Minute*8)
	if err != nil {
		fmt.Println("error in set cache")
		return
	}
	PostResponseWriter(w, user)
}

func (s *service) GetAllUser(w http.ResponseWriter, req *http.Request) {
	q := `SELECT id, username, mail, no_of_post, bio, created_at FROM "users"`
	rows, err := s.DB.Query(q)
	if err != nil {
		http.Error(w, "Bad Request Err", http.StatusBadRequest)
		return
	}
	defer rows.Close()
	user := []User{}
	for rows.Next() {
		temp := &User{}
		err = rows.Scan(&temp.ID, &temp.UserName, &temp.Mail, &temp.NoOfPost, &temp.Bio, &temp.CreatedAt)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "DATABASE Err", http.StatusInternalServerError)
			return
		}
		user = append(user, *temp)
	}
	if err = rows.Err(); err != nil {
		fmt.Println(err)
		http.Error(w, "DATABASE Err", http.StatusInternalServerError)
		return
	}
	GetResponseWriter(w, user)
}

func (s *service) GetUser(w http.ResponseWriter, req *http.Request) {
	id := GetParam(req, "id")
	if id == "" {
		http.Error(w, "Invalid Request", http.StatusNotAcceptable)
		return
	}
	userRD, err := GetUserCache(id)
	if err != nil {
		user := &User{}
		q := `SELECT id, username, mail, no_of_post, bio, created_at FROM "users" WHERE id = ($1)`
		err = s.DB.QueryRow(q, id).Scan(&user.ID, &user.UserName, &user.Mail, &user.NoOfPost, &user.Bio, &user.CreatedAt)
		if err != nil {
			http.Error(w, "DATABASE Err", http.StatusInternalServerError)
			return
		}
		err = SetUserCache(*user, time.Minute*3)
		if err != nil {
			fmt.Println(err)
			return
		}
		GetResponseWriter(w, user)
		return
	}
	GetResponseWriter(w, userRD)
}

func (s *service) DelUser(w http.ResponseWriter, req *http.Request) {
	bearerToken := req.Header.Get("Authorization")
	token := strings.Split(bearerToken, " ")[1]
	user_id, err := redis.GetCache("token:"+token)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "UnAuthorization", 600)
		return
	}
	if user_id == "" {
		http.Error(w, "UnAuthorization", 600)
	}
	reqid := GetParam(req, "id")
	if reqid != user_id {
		http.Error(w, "UnAuthorization", 600)
		return
	}

	q := `DELETE FROM users WHERE id = ($1)`
	tx, err := s.DB.Begin()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	id, err := strconv.Atoi(user_id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	_, err = tx.Exec(q, &id)
	if err != nil {
		tx.Rollback()
		fmt.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		fmt.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	GetResponseWriter(w, "User Deleted!")
}

func (s *service) Addpost(w http.ResponseWriter, req *http.Request) {
	CreationPost := &CreationPost{}
	post := &Post{}
	err := RequestReader(req, CreationPost)
	if err != nil {
		fmt.Println(err)
		return
	}
	q := `INSERT INTO posts (user_id, url) VALUES ($1, $2) RETURNING id, user_id, url, likes, created_at`
	err = s.DB.QueryRow(q, &CreationPost.UserId, &CreationPost.URL).Scan(&post.ID, &post.User, &post.URl, &post.Like, &post.CreatedAt)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = SetPostCache(*post, time.Minute*8)
	if err != nil {
		fmt.Println("err in set post acache", err)
		return
	}
	tx, err := s.DB.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}
	incq := `UPDATE users SET no_of_post = no_of_post + 1 WHERE id = $1`
	_, err = tx.Exec(incq, &CreationPost.UserId)
	if err != nil {
		tx.Rollback()
		fmt.Println(err)
		return
	}
	err = tx.Commit()
	if err != nil {
		fmt.Println(err)
		return
	}
	var user_id string = strconv.Itoa(post.User)
	err = IncPostsInUser(user_id)
	if err != nil {
		fmt.Println(err)
		return
	}
	PostResponseWriter(w, post)
}

func (s *service) GetUserPost(w http.ResponseWriter, req *http.Request) {
	id := GetParam(req, "user_id")
	var posts []Post
	q := `SELECT id, user_id, url, likes, created_at FROM "posts" WHERE user_id = ($1)`
	rows, err := s.DB.Query(q, id)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		temp := &Post{}
		err = rows.Scan(&temp.ID, &temp.User, &temp.URl, &temp.Like, &temp.CreatedAt)
		if err != nil {
			fmt.Println(err)
		}
		posts = append(posts, *temp)
	}
	if err = rows.Err(); err != nil {
		fmt.Println(err)
		http.Error(w, "DATABASE Err", http.StatusInternalServerError)
	}
	GetResponseWriter(w, posts)
}

func (s *service) GetPost(w http.ResponseWriter, req *http.Request) {
	id := GetParam(req, "id")
	post, err := GetPostCache(id)
	if err != nil {
		q := `SELECT id, user_id, url, likes, created_at FROM "posts" WHERE id = ($1)`
		post := &Post{}

		err = s.DB.QueryRow(q, id).Scan(&post.ID, &post.User, &post.URl, &post.Like, &post.CreatedAt)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = SetPostCache(*post, time.Minute*3)
		if err != nil {
			fmt.Println(err)
			return
		}
		GetResponseWriter(w, post)
		return
	}

	GetResponseWriter(w, post)
}
