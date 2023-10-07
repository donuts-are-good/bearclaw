package main

import (
	"bufio"
	"encoding/xml"
	"io"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/russross/blackfriday"
)

type RSSItem struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Link        string `xml:"link"`
	GUID        string `xml:"guid"`
	PubDate     string `xml:"pubDate"`
}

type RSSChannel struct {
	Title       string    `xml:"title"`
	Link        string    `xml:"link"`
	Description string    `xml:"description"`
	Item        []RSSItem `xml:"item"`
}

type RSSFeed struct {
	XMLName xml.Name   `xml:"rss"`
	Version string     `xml:"version,attr"`
	Channel RSSChannel `xml:"channel"`
}

func CreateXMLRSSFeed(inFolder, outFolder string) {

	// get the posts we have
	files, _ := os.ReadDir(inFolder)

	// most recently modified file comes first
	sort.Slice(files, func(i, j int) bool {

		// get info for each file in the list
		fi, _ := os.Stat(filepath.Join(inFolder, files[i].Name()))
		fj, _ := os.Stat(filepath.Join(inFolder, files[j].Name()))
		return fi.ModTime().After(fj.ModTime())
	})

	rss := RSSFeed{
		XMLName: xml.Name{
			Space: "http://www.w3.org/2005/Atom",
			Local: "feed",
		},
		Version: "2.0",
		Channel: RSSChannel{
			Title:       "My Blog",
			Link:        "https://mycoolblog.com",
			Description: "A blog about cool stuff",
		},
	}

	// rss.channel.atomlink = atomlink{rel: "self", href: "https://mycoolblog.com/feed.xml"}
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".md" {
			fileInfo, err := os.Stat(filepath.Join(inFolder, file.Name()))
			if err != nil {
				continue
			}
			modTime := fileInfo.ModTime().Format(time.RFC1123)
			title := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
			markdownFile, err := os.Open(filepath.Join(inFolder, file.Name()))
			if err != nil {
				continue
			}

			reader := bufio.NewReader(markdownFile)
			markdown, err := io.ReadAll(reader)
			markdownFile.Close()
			if err != nil {
				continue
			}

			html := blackfriday.MarkdownCommon(markdown)

			var item RSSItem
			item.Title = title
			item.Description = string(html)
			item.Link = "https://mycoolblog.com/" + url.PathEscape(file.Name()) + ".html"
			item.GUID = item.Link
			item.PubDate = modTime
			rss.Channel.Item = append(rss.Channel.Item, item)
		}
	}

	// create the rss file
	rssFile, _ := os.Create(filepath.Join(outFolder, "feed.xml"))

	// don't forget to close it
	defer rssFile.Close()

	// write the rss feed to file
	enc := xml.NewEncoder(rssFile)
	enc.Indent("", "  ")
	if err := enc.Encode(rss); err != nil {
		log.Fatalf("Error marshaling XML: %v", err)
	}
}
