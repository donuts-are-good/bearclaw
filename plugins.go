package main

import (
	"archive/zip"
	"bytes"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func FindZips(folderPath string) error {
	files, err := os.ReadDir(folderPath)
	if err != nil {
		return err
	}

	// loop through the files in the plugins dir
	for _, file := range files {
		// just look at the zips
		if filepath.Ext(file.Name()) == ".zip" {

			// open each zip file
			zipFilePath := filepath.Join(folderPath, file.Name())
			r, err := zip.OpenReader(zipFilePath)
			if err != nil {
				return err
			}
			defer r.Close()

			// extract properly
			for _, zipFile := range r.File {

				// get just the name, no extension
				fpath := filepath.Join(folderPath, strings.TrimSuffix(file.Name(), ".zip"))

				// handle directories
				if zipFile.FileInfo().IsDir() {
					if strings.HasPrefix(zipFile.Name, fpath) {
						continue
					}
					fpath = filepath.Join(fpath, filepath.Base(zipFile.Name))
					os.MkdirAll(fpath, os.ModePerm)
					continue
				}
				if err = os.MkdirAll(fpath, os.ModePerm); err != nil {
					return err
				}

				// skip nonsense files
				if strings.HasPrefix(zipFile.Name, "._") {
					continue
				}

				readcloser, err := zipFile.Open()
				if err != nil {
					return err
				}
				defer readcloser.Close()

				destFile, err := os.Create(filepath.Join(fpath, filepath.Base(zipFile.Name)))
				if err != nil {
					return err
				}
				defer destFile.Close()

				if _, err = io.Copy(destFile, readcloser); err != nil {
					return err
				}
			}

			// delete the zip after we're done.
			if err = os.Remove(zipFilePath); err != nil {
				return err
			}
		}
	}
	return nil
}

func ScanForPluginCalls(html []byte) ([]byte, error) {
	log.Println("Processing:\t plugin calls")
	// regexp to find the plugin call
	re := regexp.MustCompile(`<!-- plugin "(.+)" -->`)
	matches := re.FindAllSubmatch(html, -1)
	if len(matches) == 0 {
		// we probably messed up the plugins call, or it didn't match anything.
		return html, nil
	}

	// swap the plugin calls for the content of the plugin file
	for _, match := range matches {
		filePath := string(match[1])
		fileContent, err := os.ReadFile(filePath)
		if err != nil {
			return nil, err
		}
		html = bytes.Replace(html, match[0], fileContent, -1)
	}

	return html, nil
}
