package main

import (
	"errors"
	"os"
	"path/filepath"
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
	_, err = file.WriteString("./public/default")
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
	_, err = file.WriteString("./"+visibility+"/"+name)
	check(err)
	file.Close()
}
