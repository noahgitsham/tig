package main

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func check(err error){
	if err != nil{
		panic(err)
	}
}

func getCurrentDir() string {
	currentDirectory, err := os.Getwd()
	check(err)
	return currentDirectory
}

func dirNotExists(path string) bool {
	_, err := os.Stat(path)
	return errors.Is(err, os.ErrNotExist)
}

// Finds base directory ()
func findGitDir(path string) string {
	for dirNotExists(filepath.Join(path, ".git")) && path != "/" {
		path = filepath.Dir(path)
	}
	if path == "/" {
		return ""
	}
	return path
}

func getHeadCommit() string {
	return "0"
}

func tigInit() {
	err := os.Mkdir(".tig", 0755)
	check(err)
	err = os.Mkdir(".tig/public", 0755)
	check(err)	
	err = os.Mkdir(".tig/private", 0755)
	check(err)	
	file, err := os.Create(".tig/public/default")
	check(err)		
	file.Close()
	file, err = os.Create(".tig/HEAD")
	check(err)
	_, err = file.WriteString(".tig/public/default")
	check(err)
	file.Close()
}

func addGroup(name string, visibility string) {
	file, err := os.Create(".tig/"+visibility+"/"+name)
	check(err)
	file.Close()
}

func deleteGroup(name string, visibility string) {
	err := os.Remove((".tig/"+visibility+"/" + name))
	check(err)
}

func switchGroup(name string, visibility string) {
	file, err := os.Create(".tig/HEAD")
	check(err)
	_, err = file.WriteString(".tig/"+visibility+"/"+name)
	check(err)
	file.Close()
}

func addLinks(filepaths []string) {
	// find the filepath stored in git HEAD
	filepathBytes, err := os.ReadFile(".git/HEAD")
	check(err)
	// convert from []byte to string
	filepath := string(filepathBytes)

	// find the hash at found filepath
	extendedFilepath := ".git/"+strings.Split(filepath, " ")[1]
	extendedFilepath = strings.TrimSpace(extendedFilepath)
	hashBytes, err := os.ReadFile(extendedFilepath)
    check(err)
    // convert from []byte to string
    hash := string(hashBytes)
    hash = strings.TrimSpace(hash)

	// find the filepath stored in tig HEAD
	filepathBytes, err = os.ReadFile(".tig/HEAD")
	check(err)
	// convert from []byte to string
	filepath = string(filepathBytes)
	
	// find the file contents at found filepath
	fileContentBytes, err := os.ReadFile(filepath)
	check(err)
	// convert from []byte to string
	fileContent := string(fileContentBytes)

	// add code to check if file already there

	// iterate through links to update array
	var links []string
	filepathsString := strings.Join(filepaths, " ")
	if fileContent == "" {
		links = []string{(hash + " " + filepathsString)}
	} else {
		added := false
		links = strings.Split(fileContent, "\n")
		for i, link := range links {
			linkHash := strings.Split(link, " ")[0]
			if hash == linkHash{
				links[i] = links[i] + " " + filepathsString
				added = true
			}
		}
		if (added == false) {
			links = append(links, (hash + " " + filepathsString))
		}
	}

	// write commit array back to file
	fileContent = strings.Join(links, "\n")
	file, err := os.Create(filepath)
	check(err)
	_, err = file.WriteString(fileContent)
	check(err)
	file.Close()
}
