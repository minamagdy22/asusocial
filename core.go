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

func IsLogged() bool {
	return currentUser.FirstName != ""
}

func Logout() {
	currentUser = User{}
}

func Whoami() string {
	return currentUser.FirstName + " " + currentUser.SecondName
}

// AddNewUser function add new user to the database
func AddNewUser(u User) {
	db.Create(&u)
}

func GetUser(id string) User {
	var u User
	db.Where("id = ?", id).First(&u)
	return u
}

// GetAllUsers get all users from the database
func GetAllUsers() []User {
	var users []User
	db.Find(&users)
	return users
}

func GetUserPosts(u User) []Post {
	var posts []Post
	db.Model(&u).Related(&posts)
	return posts
}

func GetUserFriends(u User) []Friend {
	var friends []Friend
	db.Model(&u).Related(&friends)
	return friends
}
func AddFriend(u User) {
	var friend Friend
	friend.UserID = currentUser.ID
	friend.FriendID = u.ID
	db.Create(&friend)
	// add counter friendship
	friend.ID = 0
	friend.UserID = u.ID
	friend.FriendID = currentUser.ID
	db.Create(&friend)
}

func AddPost(p Post) {
	p.UserID = currentUser.ID
	db.Create(&p)
}

func GetPost(id string) Post {
	var p Post
	db.Where("id = ?", id).First(&p)
	return p
}

func AddNewGroup(g Group) {
	g.AdminID = currentUser.ID
	u := currentUser
	db.Model(&u).Association("Groups").Append(&g)
}

func JoinGroup(g Group) {
	u := currentUser
	db.Model(&u).Association("Groups").Append(&g)
}
func AddGroupPost(g Group, p Post) {

}

func GetAllGroups() []Group {
	var groups []Group
	db.Find(&groups)
	return groups
}

func GetUserGroup(u User) []Group {
	var groups []Group
	db.Model(&u).Related(&groups, "Groups")
	return groups
}
