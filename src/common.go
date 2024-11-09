package main

import (
    "os"
	"errors"
    "path/filepath"
)

func getCurrentDir() string {
	currentDirectory, err := os.Getwd()
	if (err != nil) {
		panic(err)
	}
	return currentDirectory
}

func dirNotExists(path string) bool {
	_, err := os.Stat(path)
	return errors.Is(err, os.ErrNotExist)
}

// Finds base directory ()
func findGitDir(path string) string {
	for (dirNotExists(filepath.Join(path, ".git")) && path != "/") {
		path = filepath.Dir(path)
	} 
	if (path == "/") {
		return ""
	}
	return path
}

func getHeadCommit() string {
	return "0"
}
