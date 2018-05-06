package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
)

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
	for true {
		ClearScreen()
		Welcome()
		var command string
		fmt.Print(">> ")
		fmt.Scan(&command)
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
