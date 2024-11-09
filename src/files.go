package main

func getCurrentDir() (string, err) {
	currentDirectory, err := os.Getwd()
	if (err != nil) {
		return err
	}
	return currentDirectory
}

func dirNotExists(path string) bool {
	_, err := os.Stat(path)
	return errors.Is(err, os.ErrNotExist)
}

// Finds base directory ()
func findBaseDir(path string) (string, err) {
	for (dirNotExists(filepath.Join(path, ".git")) && path != "/") {
		path = filepath.Dir(path)
	} 
	if (path == "/") {
		return nil, os.ErrNotExist
	}
	return path, nil
}
