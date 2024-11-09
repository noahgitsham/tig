package main

import (
	"errors"
	"os"
)

// find the command the user wants to run
func getCommand(num int) (string, error) {
	args := os.Args

	if len(args) <= num {
		return "", errors.New("No command")
	} else {
		command := args[num]
		if (num == 1 &&
		(command == "add" || 
		command == "visual" ||
		command == "group" ||
		command == "switch" ||
		command == "init")) {
			return command, nil
		} else if(num == 2 &&
		(command == "add" || 
		command == "delete")){
			return command, nil
		}else if(num == 3){
			return command,nil
		}else{
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

