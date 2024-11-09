package main

import (
	"errors"
	"os"
	"path/filepath"
)

func getCurrentDir() string {
	currentDirectory, err := os.Getwd()
	if err != nil {
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

func tigInit() error {
	err := os.Mkdir(".tig", 0755)
	if err != nil {
		return err
	} else {
		err = os.Mkdir(".tig/public", 0755)
		if err != nil {
			return err
		}
		err = os.Mkdir(".tig/private", 0755)
		if err != nil {
			return err
		}
		file, err := os.Create(".tig/public/default")
		if err != nil {
			return err
		} else {
			file.Close()
		}
		file, err = os.Create(".tig/HEAD")
		if err != nil {
			return err
		} else {
			_, err := file.WriteString("./public/default")
			if err != nil {
				return err
			}
			file.Close()
		}
	}
	return nil
}

func addGroup(name string) error {
	file, err := os.Create(".tig/public/"+name)
	if err != nil {
		return err
	} else {
		file.Close()
	}
	return nil
}
func deleteGroup(name string) error {
	err := os.Remove((".tig/public/" + name))
	if err != nil {
		return err
	}
	return nil
}
