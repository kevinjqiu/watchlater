package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/gorilla/feeds"
)

func Main() {
	feeds.Feed{}
	app := cli.NewApp{}
	app.Run(os.Args)
}
