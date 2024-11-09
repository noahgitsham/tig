package main

import (
    "os"
	"errors"
    "path/filepath"
)

func getHeadCommit() string {
	return find
}

func findGitDir(path string) (string, err) {
	file, err := findBaseDir(path)
	if (err != nil) {
		return err
	}
	return filepath.Join(file, ".git")
}

func findTigDir(path string) (string, err) {
	file, err := findBaseDir(path)
	if (err != nil) {
		return err
	}
	return filepath.Join(file, ".tig")
}

func getPublicLinks(path string) (string, err) {
	file, err := findTigDir(path)
	if (err != nil) {
		return err
	}
	for _, file := range filepath.Join(file, "links", "public") {
        fmt.Println(file.Name(), file.IsDir())
    }
	return 
}
