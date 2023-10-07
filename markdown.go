package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

// markdownToHTML converts markdown documents to HTML
func markdownToHTML(inFolder, outFolder, templateFolder string) {
	files, err := os.ReadDir(inFolder)
	if err != nil {
		fmt.Printf("%s: couldn't read '%s': %v\n", Bold(Red("ERROR")), inFolder, err)
		return
	}

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
	header, err := os.ReadFile(filepath.Join(templateFolder, "header.html"))
	if err != nil {
		fmt.Printf("%s: couldn't read 'header.html': %v\n", Bold(Red("ERROR")), err)
		return
	}
	footer, err := os.ReadFile(filepath.Join(templateFolder, "footer.html"))
	if err != nil {
		fmt.Printf("%s: couldn't read 'footer.html': %v\n", Bold(Red("ERROR")), err)
		return
	}

	result := bytes.NewBuffer(make([]byte, 4096))

	for infile := range input {

		log.Println("Processing:\t", infile)
		result.Reset()

		// open the selected markdown file
		markdownFile, err := os.Open(filepath.Join(inFolder, infile))
		if err != nil {
			fmt.Printf("%s: couldn't read '%s': %v\n", Bold(Red("ERROR")), infile, err)
			continue
		}

		// create the html file
		htmlFile, err := os.Create(filepath.Join(outFolder, infile+".html"))
		if err != nil {
			fmt.Printf("%s: couldn't create '%s': %v\n", Bold(Red("ERROR")), infile+".html", err)
			markdownFile.Close()
			continue
		}

		// read the md
		reader := bufio.NewReader(markdownFile)
		markdown, err := io.ReadAll(reader)
		if err != nil {
			fmt.Printf("%s: couldn't read from '%s': %v\n", Bold(Red("ERROR")), infile+".html", err)
			markdownFile.Close()
			htmlFile.Close()
			continue
		}

		// send the md to blackfriday
		html := MarkdownCommon(markdown)

		// assemble in order
		result.Write(header)
		result.Write(bytes.TrimSpace(html))
		result.Write(footer)

		// pass the assembled html into ScanForPluginCalls
		// var resultCopy []byte

		resultCopy := make([]byte, result.Len())
		copy(resultCopy, result.Bytes())
		htmlAfterPlugins, err := ScanForPluginCalls(result.Bytes())
		if err != nil {
			fmt.Printf("%s: error parsing plugin in '%s': %v\n", Bold(Red("ERROR")), infile+".html", err)
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
