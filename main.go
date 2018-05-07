package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/urfave/cli"
)

var (
	db          *gorm.DB
	err         error
	currentUser User
)

func init() {
	// database connection
	db, err = gorm.Open("sqlite3", "docs/database.db")
	db.LogMode(true)

	if err != nil {
		panic("failed to connect database")
	}
	// `defer` for setting a time-closing fn.
	// Migrate the schema
	db.AutoMigrate(&User{}, &Post{}, &Group{}, &Friend{})
}
func main() {
	app := cli.NewApp()
	app.Name = "asu social"
	app.Version = "19.99.0"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Mina Magdy",
			Email: "mina.tomas2000@gmail.com",
		},
		cli.Author{
			Name:  "Marco Younan",
			Email: "xmarcoied@gmail.com",
		},
	}
	app.Copyright = "(c) 2018 Serious Enterprise"
	app.HelpName = "contrive"
	app.Usage = "social network micro-application"
	app.UsageText = "contrive - demonstrating the available API"
	app.ArgsUsage = "[args and such]"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "port",
			Value: "main",
			Usage: "port to the application",
		},
	}

	app.Action = func(c *cli.Context) error {
		switch c.String("port") {
		case "main":
			Welcome()
		case "web":
			Welcome()
			GoWeb()
		case "cli":
			GoCli()
		default:
			Welcome()
			fmt.Println("Invalid port")

		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func GoWeb() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello from the web")

	})
	r.Run()
}

func GoCli() {
	scanner := bufio.NewScanner(os.Stdin)
	ClearScreen()
	Welcome()
	for true {
		var command string
		fmt.Print(">> ")
		scanner.Scan()
		command = scanner.Text()
		commands := strings.Split(command, " ")

		if commands[0] == "add" && commands[1] == "user" && len(commands) == 6 {
			// add user functionality
			u := User{
				FirstName:  commands[2],
				SecondName: commands[3],
				Email:      commands[4],
				Password:   commands[5],
			}
			AddNewUser(u)
		} else if commands[0] == "get" && commands[1] == "groups" && len(commands) == 2 {
			// get groups functionality
			groups := GetAllGroups()
			for _, k := range groups {
				fmt.Printf("(%d) %s\n", k.ID, k.Name)
			}
		} else if commands[0] == "get" && commands[1] == "users" && len(commands) == 2 {
			//get users functionality
			users := GetAllUsers()
			for _, k := range users {
				fmt.Printf("(%d) %s %s\n", k.ID, k.FirstName, k.SecondName)
			}
		} else if commands[0] == "get" && commands[1] == "post" && len(commands) == 3 {
			//get post functionality
			p := GetPost(commands[2])
			u := GetUser(strconv.Itoa(p.UserID))

			fmt.Println("Post: ", p.Content)
			fmt.Printf("By: %s %s (%d)\n", u.FirstName, u.SecondName, p.UserID)
			fmt.Println("Created at :", p.CreatedAt)
		} else if commands[0] == "get" && commands[1] == "group" && len(commands) == 3 {
			// get group functionality
			g := GetGroup(commands[2])
			fmt.Println("Name:", g.Name)
			fmt.Println("Created at:", g.CreatedAt)
			admin := GetUser(strconv.Itoa(g.AdminID))
			fmt.Printf("Admin: (%d) %s %s", admin.ID, admin.FirstName, admin.SecondName)
			posts := GetGroupPosts(g)
			fmt.Printf("Posts (%d post)\n", len(posts))
			for _, k := range posts {
				fmt.Printf("\t(%d) %s\n", k.ID, k.Content)
			}
		} else if commands[0] == "get" && commands[1] == "user" && len(commands) == 3 {
			//get user functionality
			u := GetUser(commands[2])
			fmt.Println("Name:", u.FirstName+" "+u.SecondName)
			fmt.Println("Email:", u.Email)
			fmt.Println("Password:", u.Password)
			fmt.Println("Created at:", u.CreatedAt)
			posts := GetUserPosts(u)
			fmt.Printf("Posts(%d post):\n", len(posts))
			for _, k := range posts {
				fmt.Printf("\t(%d) %s\n", k.ID, k.Content)
			}

			groups := GetUserGroup(u)
			fmt.Printf("Groups(%d group):\n", len(groups))
			for _, k := range groups {
				admin := "Member"
				if k.AdminID == u.ID {
					admin = "Admin"
				}
				fmt.Printf("\t(%s)%d -%s\n", k.Name, k.ID, admin)
			}
			friends := GetUserFriends(u)
			fmt.Printf("Friends(%d friend):\n", len(friends))
			for _, k := range friends {
				user := GetUser(strconv.Itoa(k.UserID))
				fmt.Printf("\t(%d)%s %s\n", k.UserID, user.FirstName, user.SecondName)
			}
		} else if commands[0] == "add" && commands[1] == "friend" && len(commands) == 3 {
			// add friend functionaliy
			if !IsLogged() {
				fmt.Println("Sorry man, you should login first before post")
				continue
			}
			u := GetUser(commands[2])
			AddFriend(u)
		} else if commands[0] == "add" && commands[1] == "post" {
			// add post functionality
			var p Post
			if len(commands) == 3 {
				i, _ := strconv.Atoi(commands[2])
				p.GroupID = i
			}
			if !IsLogged() {
				fmt.Println("Sorry man, you should login first before post")
				continue
			}
			fmt.Println("Enter the post you want:")
			scanner.Scan()
			content := scanner.Text()
			p.Content = content
			AddPost(p)
		} else if commands[0] == "login" && len(commands) == 3 {
			// login functionality
			if ValidateLogin(commands[1], commands[2]) {
				fmt.Println("succesfully logged in")
				u := GetUser(commands[1])
				Login(u)
			} else {
				fmt.Println("wrong password")
			}

		} else if commands[0] == "join" && commands[1] == "group" && len(commands) == 3 {
			if !IsLogged() {
				fmt.Println("Sorry man, you should login first before joining a group")
				continue
			}
			group_id, _ := strconv.Atoi(commands[2])
			g := Group{ID: group_id}
			JoinGroup(g)
		} else if commands[0] == "add" && commands[1] == "group" && len(commands) == 3 {
			if !IsLogged() {
				fmt.Println("Sorry man, you should login first before creating a group")
				continue
			}
			g := Group{Name: commands[2]}
			AddNewGroup(g)
		} else if commands[0] == "whoami" {
			if Whoami() == " " {
				fmt.Println("You aren't logged in yet")
			} else {
				fmt.Println(Whoami())
			}
		} else if commands[0] == "logout" {
			Logout()
		} else if commands[0] == "save" {
			// Save functionality , export to xml and json
			Save()
		} else if commands[0] == "exit" {
			// exit functionality
			os.Exit(3)
		} else if commands[0] == "clear" {
			// clearscreen functionality
			ClearScreen()
		} else if commands[0] == "deactivate" {
			db.LogMode(false)
		} else if commands[0] == "activate" {
			db.LogMode(true)
		} else if commands[0] == "ls" {
			// List function
			var listCommands = []string{
				"add user <first name> <second name> <email> <password>",
				"get users",
				"save",
				"exit",
				"clear",
				"ls",
				"login <id> <password>",
				"whoami",
				"logout",
				"add post ::optional <group_id>",
				"add group <name>",
				"join group <group_id>",
				"get groups",
				"get user <user_id>",
				"activate",
				"deactivate",
				"get post <post_id>",
				"get group <group_id>",
			}
			sort.Strings(listCommands)
			for _, val := range listCommands {
				fmt.Println("$", val)
			}
		} else {
			fmt.Println("Invalid")
		}

	}
}

// Welcome banner http://patorjk.com/software/taag - font Stop
func Welcome() {
	var str string = `

                                                                   /$$           /$$
                                                                  |__/          | $$
  /$$$$$$   /$$$$$$$ /$$   /$$        /$$$$$$$  /$$$$$$   /$$$$$$$ /$$  /$$$$$$ | $$
 |____  $$ /$$_____/| $$  | $$       /$$_____/ /$$__  $$ /$$_____/| $$ |____  $$| $$
  /$$$$$$$|  $$$$$$ | $$  | $$      |  $$$$$$ | $$  \ $$| $$      | $$  /$$$$$$$| $$
 /$$__  $$ \____  $$| $$  | $$       \____  $$| $$  | $$| $$      | $$ /$$__  $$| $$
|  $$$$$$$ /$$$$$$$/|  $$$$$$/       /$$$$$$$/|  $$$$$$/|  $$$$$$$| $$|  $$$$$$$| $$
 \_______/|_______/  \______/       |_______/  \______/  \_______/|__/ \_______/|__/
                                                                                
(c) asu social 2018 - All Rights Reserved

`
	fmt.Println(str)
}

//ClearScreen is cmd function to clean the screen
func ClearScreen() {
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}
