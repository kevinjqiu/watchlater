package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/codegangsta/cli"
	"github.com/gorilla/feeds"
)

type Config struct {
	Name         string `toml:"name"`
	Author       string `toml:"author"`
	AssetFolder  string `toml:"asset_folder"`
	ServerPrefix string `toml:"server_prefix"`
}

func Run(c *cli.Context) {
	var config Config

	configFile := c.String("config")
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(fmt.Errorf("Cannot read file: %s", configFile))
	}

	_, err = toml.Decode(string(data), &config)
	if err != nil {
		panic(fmt.Errorf("Cannot decode toml from file: %s", configFile))
	}
	fmt.Println(config)
}

func main() {
	fmt.Println(feeds.Feed{})
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
