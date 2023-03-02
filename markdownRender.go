package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/russross/blackfriday"
)

// markdownToHTML converts markdown documents to HTML
func markdownToHTML(inFolder, outFolder, templateFolder string) {
	files, _ := os.ReadDir(inFolder)

	var wg sync.WaitGroup
	input := make(chan string, 128)

	wg.Add(runtime.NumCPU())
	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			fileProcessor(input, inFolder, outFolder, templateFolder)
			wg.Done()
		}()
	}

	for _, file := range files {
		// only markdown files
		if filepath.Ext(file.Name()) == ".md" {
			input <- file.Name()
		}
	}
	close(input)
	wg.Wait()
}

func fileProcessor(input <-chan string, inFolder, outFolder, templateFolder string) {
	header, _ := os.ReadFile(path.Join(templateFolder, "header.html"))
	footer, _ := os.ReadFile(path.Join(templateFolder, "footer.html"))
	// FIXME: do literally any kind of error handling

	result := bytes.NewBuffer(make([]byte, 4096))

	for infile := range input {
		result.Reset()

		// open the selected markdown file
		markdownFile, _ := os.Open(path.Join(inFolder, infile))

		// create the html file
		htmlFile, _ := os.Create(path.Join(outFolder, infile+".html"))

		// read the md
		reader := bufio.NewReader(markdownFile)
		markdown, _ := io.ReadAll(reader)

		// send the md to blackfriday
		html := blackfriday.MarkdownCommon(markdown)

		// assemble in order
		result.Write(header)
		result.Write(bytes.TrimSpace(html))
		result.Write(footer)

		// pass the assembled html into ScanForPluginCalls
		var resultCopy []byte
		copy(resultCopy, result.Bytes())
		htmlAfterPlugins, htmlAfterPluginsErr := ScanForPluginCalls(result.Bytes())
		if htmlAfterPluginsErr != nil {
			log.Println("error inserting plugin content: ", htmlAfterPluginsErr)
			log.Println("returning content without plugin...")

			// if there's an error, let's just take the html from before the error and use that.
			htmlAfterPlugins = resultCopy
		}

		io.Copy(htmlFile, bytes.NewReader(htmlAfterPlugins))

		// don't forget to close it when done
		markdownFile.Close()
		htmlFile.Close()
	}
}
