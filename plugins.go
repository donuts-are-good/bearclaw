package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func ScanForPluginCalls(html string) (string, error) {

	// regexp to find the plugin call
	re := regexp.MustCompile(`<!-- plugin "(.+)" -->`)
	matches := re.FindAllStringSubmatch(html, -1)
	if len(matches) == 0 {

		// we probably messed up the plugins call, or it didn't match anything.
		return html, nil
	}

	// swap the plugin calls for the content of the plugin file
	for _, match := range matches {
		filePath := match[1]
		fileContent, err := os.ReadFile(filePath)
		if err != nil {
			return "", err
		}
		html = strings.Replace(html, match[0], string(fileContent), -1)
	}

	return html, nil
}

func ExtractPlugins() error {

	// get the list of files in the plugins folder
	files, err := os.ReadDir(pluginsFolder)
	if err != nil {
		return err
	}

	// loop through each file
	for _, file := range files {

		// check if the file is a zip file
		if filepath.Ext(file.Name()) == ".zip" {

			// open the zip file
			zipReader, err := zip.OpenReader(pluginsFolder + "/" + file.Name())
			if err != nil {
				return err
			}

			defer zipReader.Close()

			// loop through each file in the zip
			for _, zipFile := range zipReader.File {
				destination := pluginsFolder + "/" + zipFile.Name
				if _, err := os.Stat(destination); err == nil {
					// ask the user if they want to overwrite the file
					fmt.Printf("%s already exists. Would you like to overwrite it? (y/n)\n", destination)
					var response string
					fmt.Scanln(&response)
					if response != "y" {
						continue
					}
				}

				// extract the file to the plugins folder
				err = extractFile(zipFile, pluginsFolder)
				if err != nil {
					return err
				}

			}

			// delete the zip file
			err = os.Remove(pluginsFolder + "/" + file.Name())
			if err != nil {
				return err
			}

		}
	}
	return nil
}

func extractFile(file *zip.File, pluginsFolder string) error {

	// get the file info
	fileInfo := file.FileInfo()

	// create the destination file path
	destinationFilePath := filepath.Join(pluginsFolder, file.Name)

	// create the destination directory
	err := os.MkdirAll(filepath.Dir(destinationFilePath), fileInfo.Mode())
	if err != nil {
		return err
	}

	// open the source file
	sourceFile, err := file.Open()
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	// create the destination file
	destinationFile, err := os.OpenFile(destinationFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, fileInfo.Mode())
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	// copy the source file to the destination file
	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}

	return nil
}
