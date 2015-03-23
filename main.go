package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/codegangsta/cli"
	"github.com/gorilla/feeds"
)

type Config struct {
	Author       string `toml:"author"`
	Email        string `toml:"email"`
	Title        string `toml:"title"`
	Link         string `toml:"link"`
	Description  string `toml:"description"`
	AssetFolder  string `toml:"asset_folder"`
	ServerPrefix string `toml:"server_prefix"`
}

func Run(c *cli.Context) {
	var config Config
	var err error
	var data []byte
	var feed *feeds.Feed
	var atom string

	configFile := c.String("config")
	data, err = ioutil.ReadFile(configFile)
	if err != nil {
		panic(fmt.Errorf("Cannot read file: %s", configFile))
	}

	_, err = toml.Decode(string(data), &config)
	if err != nil {
		panic(fmt.Errorf("Cannot decode toml from file: %s", configFile))
	}

	feedGenerator := FeedGenerator{
		Root:   config.AssetFolder,
		Title:  config.Title,
		Author: config.Author,
		Email:  config.Email,
		Link:   config.Link,
	}

	feed, err = feedGenerator.Generate()
	if err != nil {
		log.Panicln(err)
	}

	atom, err = feed.ToAtom()
	fmt.Println(atom)
}

func main() {
	app := cli.NewApp()
	app.Name = "watchlater"
	app.Usage = "Generate an atom feed for audio/video files you want to watch/listen to later"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "c, config",
			Usage: "path to config file",
		},
	}

	app.Action = Run
	app.Run(os.Args)
}
