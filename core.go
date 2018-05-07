package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"log"
)

func jsonPrettyPrint(in string) string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(in), "", "\t")
	if err != nil {
		return in
	}
	return out.String()
}
func Save() {
	xmlrespond, _ := xml.Marshal(GetAllUsers())
	err = ioutil.WriteFile("data.xml", []byte(xmlrespond), 0644)
	jsonresond, _ := json.Marshal(GetAllUsers())
	err = ioutil.WriteFile("data.json", []byte(jsonPrettyPrint(string(jsonresond))), 0644)
	if err != nil {
		log.Println(err)
	}
}

// ValidateLogin used to validate data
func ValidateLogin(id string, password string) bool {
	var u User
	db.Find(&u).Where("id = ?", id)
	return u.Password == password
}
func Login(u User) {
	currentUser = u
}

func isLogged() bool {
	return currentUser.FirstName != ""
}

func Logout() {
	currentUser = User{}
}

// AddNewUser function add new user to the database
func AddNewUser(u User) {
	db.Create(&u)
}

func GetUser(id string) User {
	var u User
	db.Find(&u).Where("id = ?", id)
	return u
}

// GetAllUsers get all users from the database
func GetAllUsers() []User {
	var users []User
	db.Find(&users)
	return users
}

func AddFriend(u User) {

}

func AddPost(p Post) {
	p.UserID = currentUser.ID
	db.Create(&p)
}

func AddNewGroup(g Group) {
	g.AdminID = currentUser.ID
	db.Create(&g)
}

func AddGroupPost(g Group, p Post) {

}
