package main

import (
	"github.com/announce/altogether/al2"
	"github.com/urfave/cli"
	"log"
	"os"
)

var Version string = "0.0.0"

func main() {
	err := CreateApp().Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}

func CreateApp() *cli.App {
	app := cli.NewApp()
	app.Usage = "A CLI tool to synchronize configuration files between Alfred and Albert"
	app.Description = `https://github.com/announce/altogether`
	app.Version = Version
	app.Author = "Kenta Yamamoto"
	app.Email = "ymkjp@jaist.ac.jp"
	app.Commands = al2.Commands
	return app
}
