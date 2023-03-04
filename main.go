package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/fsnotify/fsnotify"
)

// init runs before main()
func init() {
	// we are making a list of folders here to check for the presence of
	// if they don't exist, we create them
	foldersToCreate := []string{inFolder, outFolder, templateFolder, pluginsFolder}
	createFoldersErr := createFolders(foldersToCreate)
	if createFoldersErr != nil {
		log.Println(createFoldersErr)
	}
	FindZips(pluginsFolder)
}

func checkFlags() {
	flag.StringVar(&inFolder, "input", inFolder, "the input folder for markdown files")
	flag.StringVar(&outFolder, "output", outFolder, "the output folder for html files")
	flag.StringVar(&templateFolder, "templates", templateFolder, "the templates folder for header and footer html files")
	flag.StringVar(&pluginsFolder, "plugins", pluginsFolder, "the plugins folder for plugins")

	flag.Parse()

	// there is some concern whether there is potential for infinite write loop
	// when using --watch and setting your folders to the same directory
	// the assumption is that if you're outputting your built html to a
	// 'watched' directory, fsnotify will trigger a rebuild each time a build
	// deposits files into the watched directory.
	// this should prevent that.
	if inFolder == outFolder || inFolder == templateFolder || inFolder == pluginsFolder || outFolder == templateFolder || outFolder == pluginsFolder || templateFolder == pluginsFolder {
		message := "Error: The input, output, templates, and plugins folders must be different directories"
		log.Panicln(message)
	}
}

func main() {

	// chek for directory variable overrides
	checkFlags()

	// check to see if the user ran with --watch
	watchFlag := flag.Bool("watch", false, "watch the current directory for changes")
	flag.Parse()

	// if they did...
	if *watchFlag {

		// make a list of folders to keep an eye on
		foldersToWatch := []string{templateFolder, inFolder}

		// the process to watch it goes in a goroutine
		go watchFoldersForChanges(foldersToWatch)

		// give the user some type of confirmation
		fmt.Println("Waiting for changes: ", foldersToWatch)

		// select {} is a blocking operation that keeps
		// the program from closing
		select {}
	}

	// now, if nothing has gone wrong, we process the html
	markdownToHTML(inFolder, outFolder, templateFolder)
	createPostList(inFolder, outFolder, templateFolder)
	createAboutPage(outFolder, templateFolder)
}

// recreateHeaderFooterFiles recreates the header and footer files
// if we're rebuilding the templates, we'll need these.
func recreateHeaderFooterFiles(templatesFolder string) error {
	headerFile, err := os.Create(templatesFolder + "/header.html")
	if err != nil {
		return err
	}
	defer headerFile.Close()
	_, err = headerFile.WriteString(headerHTML)
	if err != nil {
		return err
	}
	footerFile, err := os.Create(templatesFolder + "/footer.html")
	if err != nil {
		return err
	}
	defer footerFile.Close()
	_, err = footerFile.WriteString(footerHTML)
	if err != nil {
		return err
	}

	return nil
}

func watchFoldersForChanges(folders []string) {
	// range through the watched files
	for _, folder := range folders {

		// create a new watcher for each watched folder
		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			log.Fatal(err)
		}

		// don't forget to close it
		defer watcher.Close()

		// make a channel for goroutine messaging
		done := make(chan bool)
		go func() {
			// for { ev ... ver }
			for {
				select {
				case event := <-watcher.Events:

					// if there's an event, remark and rebuild
					log.Println("modified:", event.Name, " - rebuilding files..")
					markdownToHTML(inFolder, outFolder, templateFolder)
					createPostList(inFolder, outFolder, templateFolder)

					// if we get an error instead of event, announce it
				case err := <-watcher.Errors:
					log.Println("error:", err)

				}
			}
		}()

		err = watcher.Add(folder)
		if err != nil {
			log.Fatal(err)
		}
		<-done
	}
}

// createFolders takes a list of folders and checks for them to exist, and creates them if they don't exist.
func createFolders(folders []string) error {
	for _, folder := range folders {
		if _, err := os.Stat(folder); os.IsNotExist(err) {

			err = os.MkdirAll(folder, os.ModePerm)
			if err != nil {
				return err
			}

			if folder == "templates" {

				err = recreateHeaderFooterFiles(folder)
				if err != nil {
					return err
				}
			}

		}
	}
	return nil
}

// createPostList creates the page that has a list of all of your posts
func createPostList(inFolder, outFolder, templateFolder string) {

	// read the files in the directory
	files, _ := os.ReadDir(inFolder)

	// sort them by mod time
	sort.Slice(files, func(i, j int) bool {
		fi, _ := os.Stat(inFolder + "/" + files[i].Name())
		fj, _ := os.Stat(inFolder + "/" + files[j].Name())
		return fi.ModTime().After(fj.ModTime())
	})

	// unordered list
	postList := "<b class=\"info\">all posts</b><br><span class=\"text-muted\"><em><small>sorted by recently modified</small></em></span>"

	postList += "<ul>"

	// for all files...
	for _, file := range files {
		// if it is a markdown file...
		if filepath.Ext(file.Name()) == ".md" {

			// get the filename/title
			title := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))

			// put it on the list with the html
			postList += "<li><a href='" + url.PathEscape(file.Name()) + ".html'>" + title + "</a></li>"

		}
	}

	// if there are more than 0 posts make an RSS feed
	if len(files) > 0 {
		// generate the rss feed
		CreateXMLRSSFeed(inFolder, outFolder)
	}

	// end the list
	postList += "</ul>"

	// create the posts file
	htmlFile, _ := os.Create(outFolder + "/index.html")

	// don't forget to close it
	defer htmlFile.Close()

	// read the header/footer templates
	header, _ := os.ReadFile(templateFolder + "/header.html")
	footer, _ := os.ReadFile(templateFolder + "/footer.html")

	// combine them
	fmt.Fprintln(htmlFile, string(header)+postList+string(footer))
}

func createAboutPage(outFolder, templateFolder string) error {
	// create the about file
	aboutFile, err := os.Create(outFolder + "/about.html")
	if err != nil {
		return err
	}
	defer aboutFile.Close()

	// read the header/footer templates
	header, err := os.ReadFile(templateFolder + "/header.html")
	if err != nil {
		return err
	}
	footer, err := os.ReadFile(templateFolder + "/footer.html")
	if err != nil {
		return err
	}

	// explainer text for the about.html page
	// the way this entire function is structured could be a lot better
	// it's not that it's wrong, it's just messy and ugly
	siteExplainer := "<b class=\"info\">about this site</b><br>"

	// content vars
	siteName := "name:&ensp;" + site_name + "<br>"
	siteDesc := "bio:&ensp;" + site_description + "<br>"
	siteLink := "url:&ensp;<a href='" + site_link + "'>" + site_link + "</a><br>"
	siteLicense := "license:&ensp;" + site_license + "<br><br><br>"

	// author vars
	authorExplainer := "<b class=\"info\">author information</b><br>"
	authorName := "name:&ensp;" + author_name + "<br>"
	authorBio := "bio:&ensp;" + author_bio + "<br>"
	authorLinks := "author links:"
	for _, link := range author_links {
		authorLinks += "<br>👉&emsp;<a href='" + link + "'>" + link + "</a>"
	}
	authorLinks += "<br><br>"

	// plugin vars
	pluginsSection := ""
	plugins, err := os.ReadDir(pluginsFolder)
	if err != nil {
		return err
	}
	if len(plugins) > 0 {
		pluginsSection = "<b>plugin credits</b>"
		for _, plugin := range plugins {
			file, err := os.Open(pluginsFolder + "/" + plugin.Name() + "/plugin.json")
			if err != nil {
				return err
			}
			defer file.Close()

			var pluginData map[string]string
			err = json.NewDecoder(file).Decode(&pluginData)
			if err != nil {
				return err
			}

			pluginsSection += "<li>" + pluginData["plugin_name"] + " v" + pluginData["plugin_version"] + " by " + pluginData["plugin_author"] + " - " + pluginData["plugin_description"] + "<br>" + pluginData["plugin_license"] + "<br>" + pluginData["plugin_link"] + "</li>"
		}
		pluginsSection += "</ul>"
	}

	// combine the content and write to the about file
	fmt.Fprintln(aboutFile, string(header)+siteExplainer+siteName+siteDesc+siteLink+siteLicense+authorExplainer+authorName+authorBio+authorLinks+pluginsSection+string(footer))

	return nil
}
