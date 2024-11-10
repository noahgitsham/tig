package main

import (
	"github.com/gin-gonic/gin"
	"strings"
	"os/exec"
	"bufio"
	"slices"
)

func startServer() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}



///////////////////////////////
// GIT GRAPH MERMAID CONVERT //
///////////////////////////////

func getGitLog() ([]string, error) {
	out, err := exec.Command("git", "log", "--all", "--author-date-order", "--pretty=format:\"%h,%p,%s,%D\"").Output()
	if (err != nil) {
		return nil, err
	}
	var outputList []string
	scanner := bufio.NewScanner(strings.NewReader(string(out)))
	for scanner.Scan() {
		outputList = append(outputList, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return outputList, nil
}

const (  // iota is reset to 0
        HASH = iota  // c0 == 0
        PARENT_HASH = iota  // c1 == 1
        MERGE_HASHES = iota  // c2 == 2
)

func parseLogLine(line string) (hash string, parentHash string, mergeHashes []string) {
	split := strings.Split(line, ",")
	hash = split[0]

	parents := strings.Split(split[1], " ")
	if len(parents) >= 1 {
		parentHash = parents[0]
	}
	if len(parents) > 1 {
		mergeHashes = parents[1:]
	}
	return
}

type commit struct {
	hash string
	parentHash string
	mergeHashes []string
	children []string
}

func parseLogs() (map[string]commit, error){
	logs, err := getGitLog()
	if (err != nil) {
		return nil, err
	}
	commitsMap := make(map[string]commit)
	for _, log := range logs {
		hash, parentHash, mergeHashes := parseLogLine(log)
		commitsMap[hash] = commit{hash: hash, parentHash: parentHash, mergeHashes: mergeHashes}
	}
	return commitsMap, nil
}

func getCommitsReverseChronological() ([]string, error) {
	logs, err := getGitLog()
	if (err != nil) {
		return nil, err
	}
	var outputList []string
	for _, log := range logs {
		hash, _, _ := parseLogLine(log)
		outputList = append(outputList, hash)
	}
	return outputList, nil
}

func getCommitsChronological() ([]string, error) {
	commits, err := getCommitsReverseChronological()
	if (err != nil) {
		return nil, err
	}
	slices.Reverse(commits)
	return commits, nil
}

func parseGitLog() (string, error) {
	// Hashmap for commit hash -> list of children
	children := make(map[string][]string)
	commits, err := getCommitsReverseChronological()
	if (err != nil) {
		return "", err
	}
	commitMap, err := parseLogs()
	if (err != nil) {
		return "", err
	}
	for _, commit := range commits {
		parent := commitMap[commit]
		commitMap[parent.parentHash].children = append(commitMap[parent.parentHash].children, commit)
		for _, hash := range commitMap[commit].mergeHashes {

		}
	}

	// for line in 
	

	// var parsedOutput string
	return "hello", nil
}
