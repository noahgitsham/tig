package main

import (
	// "fmt"
)

func main() {
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
}
