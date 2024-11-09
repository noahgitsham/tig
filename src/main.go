package main

import(
	"fmt"
)

func main() {
	// find command argument
	command, err := getCommand()
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
	}
}
