package main

import (
	"fmt"
)

func main() {
	// find command argument
	command, err := getCommand(1)
	if err != nil {
		panic(err)
	} else if command == "add" {
		// find filepaths argument
		filepaths, err := getFilepaths()
		if err != nil {
			panic(err)
		} else {
			// add filepaths to links
			fmt.Println(filepaths)
		}
	} else if command == "visual" {
		// launch visualisation
		fmt.Println("VISUALISE")
	} else if command == "init" {
		err = tigInit()
		if err != nil {
			panic(err)
		}
	} else if command == "group" {
		command, err := getCommand(2)
		if err != nil {
			panic(err)
		} else if command == "add" {
			name, err := getCommand(3)
			fmt.Printf(name)
			err = addGroup(name)
			if err != nil {
				panic(err)
			}
		} else if command == "delete" {
			name, err := getCommand(3)
			err = deleteGroup(name)
			if err != nil {
				panic(err)
			}
		}
	} else if command == "switch" {

	}
}
