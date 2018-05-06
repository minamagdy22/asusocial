package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Account struct

// User strcut
type User struct {
	CreatedAt  time.Time
	ID         int
	FirstName  string `form:"statement" json:"statement"`
	SecondName string `form:"answer" json:"answer"`
	Password   string `form:"correct_answer" json:"correct_answer"`
	Posts      []Post
}

// Post struct
type Post struct {
	CreatedAt time.Time
	ID        int
	UserID    int
	GroupID   int
	Content   string
}

// Group struct
type Group struct {
	CreatedAt time.Time
	ID        int
	AdminID   int
	Posts     []Post
}

var (
	db          *gorm.DB
	err         error
	currentUser User
)

func init() {
	// database connection
	db, err = gorm.Open("sqlite3", "database.db")
	db.LogMode(true)

	if err != nil {
		panic("failed to connect database")
	}
	// `defer` for setting a time-closing fn.
	// Migrate the schema
	db.AutoMigrate(&User{}, &Post{}, &Group{})
}
func main() {
	// create new user
	xmarcoied := User{FirstName: "mo", SecondName: "salah", Password: "sdflskdnfsdf"}
	AddNewUser(xmarcoied)
	Login(xmarcoied)
	po := Post{Content: "fooo"}
	AddPost(po)
	var p Post
	db.Model(&currentUser).Related(&p)
	fmt.Println(p)
	//Save()
}
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

func Load() {
	file, err := os.Open("foo.xml")
	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read %d bytes: %q\n", count, data[:count])
}

// AddNewUser function add new user to the database
func AddNewUser(u User) {
	db.Create(&u)
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

func Login(u User) {
	currentUser = u
}

func isLogged() bool {
	return currentUser.FirstName != ""
}

func Logout() {
	currentUser = User{}
}

func AddNewGroup(g Group) {
	g.AdminID = currentUser.ID
	db.Create(&g)
}

func AddGroupPost(g Group, p Post) {

}

