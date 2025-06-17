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
	ID 		string `json:"id"`
	User      User   `json:"user"`
	URl       string `json:"url"`
	Like      uint   `json:"like"`
	CreatedAt time.Time `json:"created_at"`
}