package main

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/kevinjqiu/feeds"

	"camlistore.org/pkg/magic"
)

type FeedGenerator struct {
	Root         string
	ServerPrefix string
	Title        string
	Description  string
	Author       string
	Email        string
	Link         string
}

func (fg FeedGenerator) Generate() (*feeds.Feed, error) {
	now := time.Now()

	feed := &feeds.Feed{
		Title:       fg.Title,
		Description: fg.Description,
		Link:        &feeds.Link{Href: fg.Link},
		Author:      &feeds.Author{fg.Author, fg.Email},
		Created:     now,
		Items:       []*feeds.Item{},
	}

	err := filepath.Walk(fg.Root, func(path string, fi os.FileInfo, e error) error {
		f, err := os.Open(path)
		if err != nil {
			log.Println("Cannot open %s", path)
		}

		mimeType, _ := magic.MIMETypeFromReader(bufio.NewReader(f))
		if strings.HasPrefix(mimeType, "audio/") {
			feed.Items = append(feed.Items, &feeds.Item{
				Title: path,
				Enclosure: &feeds.Enclosure{
					Url:    fmt.Sprintf("%s/%s", fg.ServerPrefix, url.QueryEscape(fi.Name())),
					Type:   mimeType,
					Length: strconv.FormatInt(fi.Size(), 10),
				},
				Description: path,
				Author:      &feeds.Author{"NA", "na@example.com"},
				Created:     fi.ModTime(),
			})
		}
		log.Println(mimeType)
		return nil

	})

	if err != nil {
		return nil, err
	} else {
		return feed, nil
	}
}
