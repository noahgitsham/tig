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

func findBaseDir(path string) (string, err) {
	for (dirNotExists(filepath.Join(path, ".git")) && path != "/") {
		path = filepath.Dir(path)
	} 
	if (path == "/") {
		return nil, os.ErrNotExist
	}
	return path, nil
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

func getPublicLinkFiles(path string) ([]string, err) {
	file, err := findTigDir(path)
	if (err != nil) {
		return err
	}
	files, err := os.ReadDir(filepath.Join(file, "links", "public"))
	if (err != nil) {
		return err
	}
	for _, file := range files {
        fmt.Println(file.Name(), file.IsDir())
    }
	return 
}
