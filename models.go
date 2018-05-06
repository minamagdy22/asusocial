package main

type UserData struct {
	ID         int
	FirstName  string
	SecondName string
	Password   string
}
type User struct {
	Info   UserData
	Friend UserData
}
