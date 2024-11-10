package main

import(
	"os"
	"errors"
)

// find the command the user wants to run
func getCommand() (string, error) {
	args := os.Args

	if len(args) <= 1 {
		return "", errors.New("No command")
	} else {
		command := args[1]
		if command == "add" || command == "visual"{
			return command, nil
		} else{
			return "", errors.New("Invalid command")
		}
	}
}

// find the filepaths from arguments
func getFilepaths() ([]string, error) {
	args := os.Args

	if len(args) <= 2 {
		return nil, errors.New("No filepaths")
	} else {
		return args[2:], nil
	}
}
