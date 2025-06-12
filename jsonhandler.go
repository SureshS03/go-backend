package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetResponseWriter(w http.ResponseWriter, res interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("By", "Suresh")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(res)
}

func PostResponseWriter(w http.ResponseWriter, res interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("By", "Suresh")
	w.WriteHeader(http.StatusCreated)
	encoder := json.NewEncoder(w)
	encoder.Encode(res)
}

func RequestReader(req *http.Request, res interface{}) error{
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(res)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return nil
}