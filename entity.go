package main

import "time"

type User struct {
	ID        int       `json:"id"`
	UserName  string    `json:"username"`
	Mail      string    `json:"mail"`
	Password  string    `json:"password"`
	NoOfPost  uint      `json:"no_of_post"`
	Bio       string    `json:"bio"`
	CreatedAt time.Time `json:"created_at"`
}

type Post struct {
	ID 		string 
	User      User
	URl       string
	Like      uint
	CreatedAt time.Time
}

type CreationPost struct {
	UserId int `json:"user_id"`
	URL string `json:"url"`
}

type GetPost struct {
	Id int `json:"id"`
	UserID int `json:"user_id"`
	URL string `json:"url"`
	Like uint `json:"likes"`
	CreatedAt time.Time `json:"created_at"`
}