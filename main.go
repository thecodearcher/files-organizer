package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"mime"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	sortedFiles := make(map[string][]string)
	sortDetails := make(map[string]int)
	var folderPath string
	fmt.Println("Enter file folder path: ")
	fmt.Scanln(&folderPath)

	fmt.Println("Organizing: ", folderPath)
	fmt.Println("Please wait...")
	files, err := ioutil.ReadDir(folderPath)

	if err != nil {
		log.Panicln(err)
	}

	for _, file := range files {
		if !file.IsDir() {
			var fileType string
			mimeType := mime.TypeByExtension(filepath.Ext(file.Name()))
			if strings.Contains(mimeType, "word") {
				fileType = "document"
			} else if strings.Contains(mimeType, "pdf") {
				fileType = "pdf"
			} else if strings.Contains(mimeType, "spreadsheet") || strings.Contains(mimeType, "excel") {
				fileType = "spreadsheets"
			} else if mimeType == "" {
				fileType = "others"
			} else {
				fileType = strings.Split(mimeType, "/")[0]
			}

			sortedFiles[fileType] = append(sortedFiles[fileType], file.Name())
			sortDetails["filesCount"] = sortDetails["filesCount"] + 1
		}
	}

	fmt.Println(".......")
	fmt.Println("Number of files scanned: ", sortDetails["filesCount"])
	fmt.Println(".......")

	for key, files := range sortedFiles {
		subfolder := filepath.Join(folderPath, key)
		os.Mkdir(subfolder, os.ModePerm)
		sortDetails["foldersCount"] = sortDetails["foldersCount"] + 1
		for _, fileName := range files {
			oldPath := filepath.Join(folderPath, fileName)
			newPath := filepath.Join(subfolder, fileName)
			os.Rename(oldPath, newPath)
			sortDetails["moveCount"] = sortDetails["moveCount"] + 1
		}
	}

	fmt.Println("Created/Modified: ", sortDetails["foldersCount"], "folders")
	fmt.Println("......")
	fmt.Println("Moved: ", sortDetails["moveCount"], "files")
	fmt.Println("......")
	fmt.Println("Done.")
}
