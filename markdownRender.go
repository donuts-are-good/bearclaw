package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/russross/blackfriday"
)

// markdownToHTML converts markdown documents to HTML
func markdownToHTML(inFolder, outFolder, templateFolder string) {
	files, _ := os.ReadDir(inFolder)

	for _, file := range files {
		// only markdown files
		if filepath.Ext(file.Name()) == ".md" {

			// open the selected markdown file
			markdownFile, _ := os.Open(inFolder + "/" + file.Name())

			// create the html file
			htmlFile, _ := os.Create(outFolder + "/" + file.Name() + ".html")

			// read the md
			reader := bufio.NewReader(markdownFile)
			markdown, _ := io.ReadAll(reader)

			// send the md to blackfriday
			html := blackfriday.MarkdownCommon(markdown)

			// read in our templates
			header, _ := os.ReadFile(templateFolder + "/header.html")
			footer, _ := os.ReadFile(templateFolder + "/footer.html")

			// assemble in order
			completeHTML := string(header) + strings.TrimSpace(string(html)) + string(footer)

			// pass the assembled html into ScanForPluginCalls
			htmlAfterPlugins, htmlAfterPluginsErr := ScanForPluginCalls(completeHTML)
			if htmlAfterPluginsErr != nil {
				log.Println("error inserting plugin content: ", htmlAfterPluginsErr)
				log.Println("returning content without plugin...")

				// if there's an error, let's just take the html from before the error and use that.
				htmlAfterPlugins = completeHTML
			}

			// don't forget to close it when done
			markdownFile.Close()
			htmlFile.Close()

			fmt.Fprintln(htmlFile, htmlAfterPlugins)
		}
	}
}
