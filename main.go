package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/russross/blackfriday"
)

var (
	inFolder       = "./markdown"
	outFolder      = "./output"
	templateFolder = "./templates"
)

func init() {
	folders := []string{inFolder, outFolder, templateFolder}
	createFoldersErr := createFolders(folders)
	if createFoldersErr != nil {
		log.Println(createFoldersErr)
	}
}

func main() {
	ConvertMarkdownToHTML(inFolder, outFolder, templateFolder)
	CreatePostListHTML(inFolder, outFolder, templateFolder)
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

// ConvertMarkdownToHTML converts markdown files in a folder to HTML files
func ConvertMarkdownToHTML(inFolder, outFolder, templateFolder string) {
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

// CreatePostListHTML creates an HTML file with a list of posts by title
func CreatePostListHTML(inFolder, outFolder, templateFolder string) {
	files, _ := os.ReadDir(inFolder)

	// Sort the files by modification time in reverse order
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
