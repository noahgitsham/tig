package main

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"runtime"
	"os/exec"
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

// function to return all associated
// filespaths from a given hash
func hashToFilepaths(hash string) ([]string, error) {
	// find the filepath stored in tig HEAD
	filepathBytes, err := os.ReadFile(".tig/HEAD")
    check(err)
    // convert from []byte to string
	filepath := string(filepathBytes)

    // find the file contents at found filepath
    fileContentBytes, err := os.ReadFile(filepath)
    check(err)
    // convert from []byte to string
    fileContent := string(fileContentBytes)

	// find the line with the correct hash
	lines := strings.Split(fileContent, "\n")
	for _, line := range lines {
		entries := strings.Split(line, " ")
		if entries[0] == hash {
			return entries[1:], nil
		}
	}

	return nil, errors.New("Hash not found")
}

// returns file contents
func fileContents() []string {
	// find the filepath stored in tig HEAD
	filepathBytes, err := os.ReadFile(".tig/HEAD")
    check(err)
    // convert from []byte to string
	filepath := string(filepathBytes)
    // find the file contents at found filepath
    fileContentBytes, err := os.ReadFile(filepath)
    check(err)
    // convert from []byte to string
    fileContent := string(fileContentBytes)
	// find the line with the correct hash
	lines := strings.Split(fileContent, "\n")
	return lines
	}

// from https://gist.github.com/sevkin/9798d67b2cb9d07cb05f89f14ba682f8
func openInBrowser(link string) error {
    var cmd string
    var args []string

    switch runtime.GOOS {
    case "windows":
        cmd = "cmd"
        args = []string{"/c", "start"}
    case "darwin":
        cmd = "open"
    default: // "linux", "freebsd", "openbsd", "netbsd"
        cmd = "xdg-open"
    }
    args = append(args, link)
    return exec.Command(cmd, args...).Start()}
