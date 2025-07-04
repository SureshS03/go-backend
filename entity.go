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
	ID        string    `json:"id"`
	User      int       `json:"user_id"`
	URl       string    `json:"url"`
	Like      uint      `json:"likes"`
	CreatedAt time.Time `json:"created_at"`
}

type CreationPost struct {
	UserId 		int    `json:"user_id"`
	URL    		string `json:"url"`
}

type GetPost struct {
	Id        int       `json:"id"`
	UserID    int       `json:"user_id"`
	URL       string    `json:"url"`
	Like      uint      `json:"likes"`
	CreatedAt time.Time `json:"created_at"`
}
