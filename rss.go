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
		fi, _ := os.Stat(inFolder + "/" + files[i].Name())
		fj, _ := os.Stat(inFolder + "/" + files[j].Name())
		return fi.ModTime().After(fj.ModTime())
	})

	var rss RSSFeed
	rss.XMLName.Space = "http://www.w3.org/2005/Atom"
	rss.XMLName.Local = "feed"
	rss.Version = "2.0"
	rss.Channel.Title = "My Blog"
	rss.Channel.Link = "https://mycoolblog.com"
	rss.Channel.Description = "A blog about cool stuff"
	// rss.Channel.AtomLink = AtomLink{Rel: "self", Href: "https://mycoolblog.com/feed.xml"}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".md" {
			fileInfo, err := os.Stat(inFolder + "/" + file.Name())
			if err != nil {
				continue
			}
			modTime := fileInfo.ModTime().Format(time.RFC1123)
			title := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
			markdownFile, err := os.Open(inFolder + "/" + file.Name())
			if err != nil {
				continue
			}
			defer markdownFile.Close()

			reader := bufio.NewReader(markdownFile)
			markdown, err := io.ReadAll(reader)
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
	rssFile, _ := os.Create(outFolder + "/feed.xml")

	// don't forget to close it
	defer rssFile.Close()

	// write the RSS feed to file
	enc := xml.NewEncoder(rssFile)
	enc.Indent("", "  ")
	if err := enc.Encode(rss); err != nil {
		log.Fatalf("Error marshaling XML: %v", err)
	}
}
