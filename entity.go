package main

type user struct{
	ID string
	UserName string
	NoOfPost uint
	Bio string
	Mail string
	Password string
}

type post struct{
	ID string
	User user
	URl string
	Like uint
}