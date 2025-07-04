package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/SureshS03/goconnect/internal/redis"
)

func SetUserCache(user User, min time.Duration) error {
	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("err in set cache", err)
		return err
	}
	id := strconv.Itoa(user.ID)
	return redis.SetCache("user:"+id, string(data), time.Minute*min)
}

func GetUserCache(id string) (*User, error) {
	data, err := redis.GetCache("user:" + id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var user User
	err = json.Unmarshal([]byte(data), &user)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &user, nil
}

func SetPostCache(post Post, min time.Duration) error {
	data, err := json.Marshal(post)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return redis.SetCache("post:"+post.ID, string(data), time.Minute*min)
}

func GetPostCache(id string) (*Post, error) {
	data, err := redis.GetCache("post:" + id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var post Post
	err = json.Unmarshal([]byte(data), &post)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &post, nil
}

func IncPostsInUser(id string) error {
	if !redis.CacheExists("user:" + id) {
		return nil
	}
	user, err := GetUserCache(id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	user.NoOfPost++
	err = SetUserCache(*user, time.Minute*5)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
