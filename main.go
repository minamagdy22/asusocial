package main

import (
	"fmt"
	"time"
)

type User struct {
	Id       int
	Name     string
	Email    string
	Password string
}

type Post struct {
	Id        int
	Content   string
	CreatedAt time.Time
	PostUser  User
}

func main() {
	//demo()
	Postdemo()
}
func CreateUser(userdata User){

}

func CreatePost(userdata User, postdata Post){
	// Putting the current time into the Post
	postdate.CreatedAt = time.Now()
}

func CurrentLoggedUser(){

}
func Postdemo() {

}

// func demo() {
// 	// Hello world page
// 	fmt.Println("Hello world")
// 	// Login/Registration
// 	fmt.Println("(1) Login", "(2) Register")
// 	var inputstr int
// 	fmt.Scanf("%d", &inputstr)
// 	if inputstr != 1 && inputstr != 2 {
// 		fmt.Println("Enter valid input")
// 	}

// 	if inputstr == 1 {
// 		// Login page
// 		var appuser User
// 		fmt.Print("Enter email:")
// 		fmt.Scanf("%s", appuser.Email)
// 		fmt.Print("Enter password:")
// 		fmt.Scanf("%s", appuser.Password)
// 		// Verification

// 	}
// 	if inputstr == 2 {
// 		// Reigstration page

// 	}
// }
