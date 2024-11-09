package main

import (
	"github.com/gin-gonic/gin"
	"strings"
	"os/exec"
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

func getGitLog() (string, error) {
	out, err := exec.Command("git", "log", "--all", "--author-date-order", "--pretty=format:\"%h,%p,%s,%D\"").Output()
	if (err != nil) {
		return "", err
	}
	return string(out), nil
}

// Returns hash, parentHash, branchCommit?, mergeCommit?
func parseLogLine(line string) (hash string, parentHash string, branchCommit string, mergeParentHashes []string) {
	split := strings.Split(line, ",")
	hash = split[0]

	parents := strings.Split(split[1], " ")
	if len(parents) != 1 {
		parentHash = parents[0]
	}
	if len(parents) > 1 {
		mergeParentHashes = parents[1:]
	}
	return
}

func parseGitLog() string {
	// Hashmap for commit hash -> branch index
	// branchOf := make(map[string]int)

	// var parsedOutput string
	return "hello"
}
