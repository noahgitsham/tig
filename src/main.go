package main

import (
	// "fmt"
)

func main() {
	// tree, err := parseGitLog()
	// check(err)
	// fmt.Println(tree)
	// find argument at position 1
	argument, err := getArgument(1)
	check(err)
	if argument == "add" {
		// find filepaths argument
		filepaths, err := getFilepaths()
		check(err)
		// add filepaths to links
		addLinks(filepaths)
	} else if argument == "visual" {
		// launch visualisation
		openInBrowser("http://localhost:8080")
		visualise()
	} else if argument == "init" {
		tigInit()
	} else if argument == "group" {
		// find argument at position 2
		argument, err := getArgument(2)
		check(err)
		if argument == "add" {
			name, err := getArgument(3)
			check(err)
			visibility := getFlags()
			addGroup(name, visibility)
			check(err)
		} else if argument == "delete" {
			name, err := getArgument(3)
			check(err)
			visibility := getFlags()
			deleteGroup(name, visibility)
			check(err)
		}
	} else if argument == "switch" {
		// find argument at position 2
		name, err := getArgument(2)
		check(err)
		visibility := getFlags()
		switchGroup(name, visibility)
	}

	//files, err := hashToFilepaths("31bb6d3e5495c5ba1638b32a788b012920272355")
	//check(err)
	//fmt.Println(files)
}
