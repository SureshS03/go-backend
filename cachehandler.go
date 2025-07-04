package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/SureshS03/goconnect/internal/redis"
)

func SetUserCache(user User) error {
	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("err in set cache",err)
		return err
	}
	id := strconv.Itoa(user.ID)
	return redis.SetCache("user:"+id, string(data), time.Minute * 8)
}

func GetUserCache(id string) (*User, error) {
	data, err := redis.GetCache("user:"+id)
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