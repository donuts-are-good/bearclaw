package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/russross/blackfriday"
)

var (
	inFolder       = "./markdown"
	outFolder      = "./output"
	templateFolder = "./templates"
)

func init() {
	foldersToCreate := []string{inFolder, outFolder, templateFolder}
	createFoldersErr := createFolders(foldersToCreate)
	if createFoldersErr != nil {
		log.Println(createFoldersErr)
	}
}

func main() {
	watchFlag := flag.Bool("watch", false, "watch the current directory for changes")
	flag.Parse()

	if *watchFlag {
		foldersToWatch := []string{templateFolder, inFolder}
		go watchFoldersForChanges(foldersToWatch)
		fmt.Println("Waiting for changes: ", foldersToWatch)
		select {}
	}

	markdownToHTML(inFolder, outFolder, templateFolder)
	createPostList(inFolder, outFolder, templateFolder)
}

func watchFoldersForChanges(folders []string) {
	for _, folder := range folders {
		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			log.Fatal(err)
		}
		defer watcher.Close()

		done := make(chan bool)
		go func() {
			for {
				select {
				case event := <-watcher.Events:
					if event.Op&fsnotify.Write == fsnotify.Write {
						log.Println("modified:", event.Name, " - rebuilding files..")
						markdownToHTML(inFolder, outFolder, templateFolder)
						createPostList(inFolder, outFolder, templateFolder)
					}
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

func createFolders(folders []string) error {
	for _, folder := range folders {
		if _, err := os.Stat(folder); os.IsNotExist(err) {
			err = os.MkdirAll(folder, os.ModePerm)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func markdownToHTML(inFolder, outFolder, templateFolder string) {
	files, _ := os.ReadDir(inFolder)

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".md" {
			markdownFile, _ := os.Open(inFolder + "/" + file.Name())
			defer markdownFile.Close()
			htmlFile, _ := os.Create(outFolder + "/" + file.Name() + ".html")
			defer htmlFile.Close()
			reader := bufio.NewReader(markdownFile)
			markdown, _ := io.ReadAll(reader)
			html := blackfriday.MarkdownCommon(markdown)
			header, _ := os.ReadFile(templateFolder + "/header.html")
			footer, _ := os.ReadFile(templateFolder + "/footer.html")
			fmt.Fprintln(htmlFile, string(header)+strings.TrimSpace(string(html))+string(footer))
		}
	}
}

func createPostList(inFolder, outFolder, templateFolder string) {
	files, _ := os.ReadDir(inFolder)
	sort.Slice(files, func(i, j int) bool {
		fi, _ := os.Stat(inFolder + "/" + files[i].Name())
		fj, _ := os.Stat(inFolder + "/" + files[j].Name())
		return fi.ModTime().After(fj.ModTime())
	})
	postList := "<ul>"
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".md" {
			title := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
			postList += "<li><a href='" + file.Name() + ".html'>" + title + "</a></li>"
		}
	}
	postList += "</ul>"
	htmlFile, _ := os.Create(outFolder + "/posts.html")
	defer htmlFile.Close()
	header, _ := os.ReadFile(templateFolder + "/header.html")
	footer, _ := os.ReadFile(templateFolder + "/footer.html")
	fmt.Fprintln(htmlFile, string(header)+postList+string(footer))
}
