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

	// scan for or make a config file
	loadConfig()

	// scaffold out the folders we need to operate
	scaffold()

	// look for plugin zips
	FindZips(pluginsFolder)

}

func main() {

	// chek for directory variable overrides
	checkFlags()

}

// recreateHeaderFooterFiles recreates the header and footer files
// if we're rebuilding the templates, we'll need these.
func recreateHeaderFooterFiles(templatesFolder string) error {
	headerFile, err := os.Create(filepath.Join(templatesFolder, "header.html"))
	if err != nil {
		return err
	}
	defer headerFile.Close()
	_, err = headerFile.WriteString(headerHTML)
	if err != nil {
		return err
	}
	footerFile, err := os.Create(filepath.Join(templatesFolder, "footer.html"))
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

// watchFolderForChange will watch an individual folder for
// any type of change, then trigger a rebuild
func watchFolderForChange(folder string) {

	// make a watcher with fsnotify
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalf("unable to watch %s : %v", folder, err)
	}
	defer watcher.Close()

	err = watcher.Add(folder)
	if err != nil {
		log.Fatalf("couldn't add a watcher: %v", err)
	}

	for {
		select {
		case event := <-watcher.Events:
			log.Println("modified:", event.Name, " - rebuilding files..")
			markdownToHTML(inFolder, outFolder, templateFolder)
			createPostList(inFolder, outFolder, templateFolder)
		case err := <-watcher.Errors:
			log.Println("error:", err)
		}
	}
}

// watchFoldersForChanges loops through a list of folders and
// passes them to watchFolderForChange
func watchFoldersForChanges(folders []string) {

	// range through the watched files
	for _, folder := range folders {

		// watch the individual folder
		go watchFolderForChange(folder)
	}

}

// createPostList creates the page that has a list of all of your posts
func createPostList(inFolder, outFolder, templateFolder string) {

	// read the files in the directory
	files, filesErr := os.ReadDir(inFolder)
	if filesErr != nil {
		log.Fatalf("unable to read posts directory: %v", filesErr)
	}

	// sort them by mod time
	sort.Slice(files, func(i, j int) bool {
		fi, _ := os.Stat(filepath.Join(inFolder, files[i].Name()))
		fj, _ := os.Stat(filepath.Join(inFolder, files[j].Name()))
		return fi.ModTime().After(fj.ModTime())
	})

	// unordered list
	var postList strings.Builder
	postList.WriteString("<b class=\"info\">all posts</b><br><span class=\"text-muted\"><em><small>sorted by recently modified</small></em></span>")
	postList.WriteString("<ul>")

	// for all files...
	for _, file := range files {
		// if it is a markdown file...
		if filepath.Ext(file.Name()) == ".md" {

			// get the filename/title
			title := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))

			// put it on the list with the html
			toWrite := [...]string{
				"<li><a href='",
				url.PathEscape(file.Name()),
				".html'>",
				title,
				"</a></li>",
			}
			for _, v := range toWrite {
				postList.WriteString(v)
			}
		}
	}

	// if there are more than 0 posts make an RSS feed
	if len(files) > 0 {
		// generate the rss feed
		CreateXMLRSSFeed(inFolder, outFolder)
	}

	// end the list
	postList.WriteString("</ul>")

	// create the posts file
	htmlFile, _ := os.Create(filepath.Join(outFolder, "index.html"))

	// don't forget to close it
	defer htmlFile.Close()

	// read the header/footer templates
	header, _ := os.ReadFile(filepath.Join(templateFolder, "header.html"))
	footer, _ := os.ReadFile(filepath.Join(templateFolder, "footer.html"))

	// combine them
	var (
		output    strings.Builder
		forOutput = [...]string{
			string(header),
			postList.String(),
			string(footer),
		}
	)
	for _, v := range forOutput {
		output.WriteString(v)
	}
	fmt.Fprintln(htmlFile, output.String())
}

func createAboutPage(outFolder, templateFolder string) error {

	log.Println("Your output folder: \t", outFolder)
	log.Println("Your templates folder: \t", templateFolder)

	// create the about file
	aboutFile, pluginErr := os.Create(filepath.Join(outFolder, "about.html"))
	if pluginErr != nil {
		log.Println("aboutFile: ", pluginErr)
		return pluginErr
	}

	// read the header/footer templates
	header, pluginErr := os.ReadFile(filepath.Join(templateFolder, "header.html"))
	if pluginErr != nil {
		return pluginErr
	}
	footer, pluginErr := os.ReadFile(filepath.Join(templateFolder, "footer.html"))
	if pluginErr != nil {
		return pluginErr
	}

	// explainer text for the about.html page
	// the way this entire function is structured could be a lot better
	// it's not that it's wrong, it's just messy and ugly
	siteExplainer := "<b class=\"info\">about this site</b><br>"
	// log.Println("siteExplainer", siteExplainer)

	// content vars
	siteName := "name:&ensp;" + site.Name + "<br>"
	siteDesc := "bio:&ensp;" + site.Description + "<br>"
	siteLink := "url:&ensp;<a href='" + site.Link + "'>" + site.Link + "</a><br>"
	siteLicense := "license:&ensp;" + site.License + "<br><br><br>"

	// log.Println("site info:", siteName, siteDesc, siteLink, siteLicense)

	// author vars
	authorExplainer := "<b class=\"info\">author information</b><br>"
	authorName := "name:&ensp;" + author.Name + "<br>"
	authorBio := "bio:&ensp;" + author.Bio + "<br>"
	authorLinks := "author links:"
	for _, link := range author.Links {
		authorLinks += "<br>👉&emsp;<a href='" + link + "'>" + link + "</a>"
	}
	authorLinks += "<br><br>"

	// log.Println("authorInfo: ", authorExplainer, authorName, authorBio, authorLinks)
	// plugin vars

	var pluginsSection strings.Builder
	pluginsSection.WriteRune(' ')

	plugins, pluginErr := os.ReadDir(pluginsFolder)
	if pluginErr != nil {
		log.Println("plugin err: ", pluginErr)
		return pluginErr
	}

	if len(plugins) > 0 {

		if len(plugins) == 1 {
			log.Printf("Extensions:\t %d plugin loaded", len(plugins))
		} else {
			log.Printf("Extensions:\t %d plugins loaded", len(plugins))
		}
		pluginsSection.Reset()
		pluginsSection.WriteString("<b>plugin credits</b>")
		for _, plugin := range plugins {
			file, err := os.Open(filepath.Join(pluginsFolder, plugin.Name(), "plugin.json"))
			if err != nil {
				log.Println("plugin error: ", err)
				return err
			}

			var pluginData map[string]string
			err = json.NewDecoder(file).Decode(&pluginData)
			if err != nil {
				log.Println("pluginData map: ", err)
				return err
			}

			toWrite := [...]string{
				"<li>",
				pluginData["plugin_name"],
				" v", pluginData["plugin_version"],
				" by ", pluginData["plugin_author"],
				" - ", pluginData["plugin_description"],
				"<br>", pluginData["plugin_license"], "<br>",
				pluginData["plugin_link"],
				"</li",
			}
			for _, v := range toWrite {
				pluginsSection.WriteString(v)
			}
		}
		pluginsSection.WriteString("</ul>")
	}

	// log.Println("pluginSection: ", pluginsSection)

	// log.Println("writeline: ", aboutFile, string(header)+siteExplainer+siteName+siteDesc+siteLink+siteLicense+authorExplainer+authorName+authorBio+authorLinks+pluginsSection+string(footer))
	// combine the content and write to the about file
	var (
		output    strings.Builder
		forOutput = [...]string{
			string(header),
			siteExplainer,
			siteName,
			siteDesc,
			siteLink,
			siteLicense,
			authorExplainer,
			authorName,
			authorBio,
			authorLinks,
			pluginsSection.String(),
			string(footer),
		}
	)
	for _, v := range forOutput {
		output.WriteString(v)
	}
	fmt.Fprintln(aboutFile, output.String())

	aboutFile.Close()

	return nil
}

// checkFlags looks at the run flags like --output when we start up
func checkFlags() {

	flag.StringVar(&inFolder, "input", inFolder, "the input folder for markdown files")
	flag.StringVar(&outFolder, "output", outFolder, "the output folder for html files")
	flag.StringVar(&templateFolder, "templates", templateFolder, "the templates folder for header and footer html files")
	flag.StringVar(&pluginsFolder, "plugins", pluginsFolder, "the plugins folder for plugins")

	watchFlag := flag.Bool("watch", false, "watch the content directories for changes")
	flag.Parse()

	isWatching = *watchFlag

	// there is some concern whether there is potential for infinite write loop
	// when using --watch and setting your folders to the same directory
	// the assumption is that if you're outputting your built html to a
	// 'watched' directory, fsnotify will trigger a rebuild each time a build
	// deposits files into the watched directory.
	// this should prevent that.
	if inFolder == outFolder || inFolder == templateFolder || inFolder == pluginsFolder || outFolder == templateFolder || outFolder == pluginsFolder || templateFolder == pluginsFolder {
		message := "Error: The input, output, templates, and plugins folders must be different directories"
		log.Fatalf(message)
	}

	if isWatching {
		// make a list of folders to keep an eye on
		foldersToWatch := []string{templateFolder, inFolder}

		// the process to watch it goes in a goroutine
		watchFoldersForChanges(foldersToWatch)

		// give the user some type of confirmation
		fmt.Println("Waiting for changes: ", foldersToWatch)

		// select {} is a blocking operation that keeps
		// the program from closing
		select {}
	}

	// now, if nothing has gone wrong, we process the html
	createAboutPage(outFolder, templateFolder)

	markdownToHTML(inFolder, outFolder, templateFolder)
	createPostList(inFolder, outFolder, templateFolder)
}
