package main

import (
	"errors"
	"os"
	flag "github.com/spf13/pflag"
)

// find the argument at a given certain index
// for commands add, visual, group, switch, init
func getArgument(num int) (string, error) {
	args := os.Args

	// maybe implement recursive check
	// in a backwards grammar
	if len(args) <= num {
		return "", errors.New("Missing argument")
	} else {
		argument := args[num]
		if (num == 1 &&
		(argument == "add" || 
		argument == "visual" ||
		argument == "group" ||
		argument == "switch" ||
		argument == "init")) {
			return argument, nil
		} else if (num == 2 || num == 3) {
			return argument, nil
		} else {
			return "", errors.New("Invalid argument")
		}
	}
}

// find flags
// for commands group add, group delete, switch
func getFlags() (string) {
	val := flag.String("type", "public", "Type: public or private")
	flag.Parse()
	return *val
}

// find the filepaths from arguments
// for command add
func getFilepaths() ([]string, error) {
	args := os.Args

	if len(args) <= 2 {
		return nil, errors.New("No filepaths")
	} else {
		return args[2:], nil
	}
}

