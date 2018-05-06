package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

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
			fmt.Println("Hello from the web")
		case "cli":
			fmt.Println("Hello from the cli")
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

func ClearScreen() {
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}
