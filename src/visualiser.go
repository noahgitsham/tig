package main

import (
	"github.com/gin-gonic/gin"
	"strings"
	"os/exec"
	"bufio"
	"slices"
	"strconv"
	"net/http"
	"fmt"
)

func visualise() {
	r := gin.Default()
	r.Static("/static", "./static")
	r.Static("/images", "./media")
	r.LoadHTMLGlob("static/*")
	r.GET("/", func(c *gin.Context) {
		// c.JSON(200, gin.H{
		// 	"message": "pong",
		// })
		tree, err := parseGitLog()
		check(err)
		fmt.Println(tree)
		
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":   "Tig",
			"content": "visualise your repository",
			"commitInfo":tree,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}



///////////////////////////////
// GIT GRAPH MERMAID CONVERT //
///////////////////////////////

func getGitLog() ([]string, error) {
	out, err := exec.Command("git", "log", "--all", "--author-date-order", "--pretty=format:%h,%p,%s,%D").Output()
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
		mergeHashes = []string {}
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

func descendentOf(child commit, parent commit) bool {

	commitMap, err := parseLogs()
	if (err != nil) {
		panic(err)
	}

	if parent.children == nil {
		return false
	}
	if parent.hash == child.hash {
		return true
	}
	for _, childToCheck := range parent.children {
		found := descendentOf(child, commitMap[childToCheck])
		if found {
			return true
		}
	}
	return false
}

func peek(stack []string) string {
	return stack[len(stack) - 1]
}

func parseGitLog() (string, error) {
	commits, err := getCommitsReverseChronological()
	if (err != nil) {
		return "", err
	}
	commitMap, err := parseLogs()
	if (err != nil) {
		return "", err
	}
	// Iterate over commits, and add them as child to their parent nodes
	for _, commit := range commits {
		child := commitMap[commit]
		temp := commitMap[child.parentHash]
		temp.children = append(temp.children, commit)
		commitMap[child.parentHash] = temp
		for _, hash := range commitMap[commit].mergeHashes {
			temp := commitMap[child.parentHash]
			temp.children = append(temp.children, hash)
			commitMap[child.parentHash] = temp
		}
	}

	// BIG MONEY
	stackList := [][]string{}
	for _, commit := range commits {
		commitStruct := commitMap[commit]

		inserted := false
		var index int
		for i := 0; i < len(stackList); i++ {
			if slices.Contains(commitStruct.children, peek(stackList[i])) {
				newStack := append(stackList[i], commit)
				stackList[i] = newStack
				inserted = true
				index = i
				break
			}
		}
		if !inserted {
			stackList = append(stackList, []string{commit})
			index = len(stackList) - 1
		}


		if len(commitStruct.mergeHashes) != 0 {
			for _, parent := range commitStruct.mergeHashes {
				// Right to left
				for i, stack := range stackList[:index] {
					if descendentOf(commitMap[parent], commitMap[peek(stack)]) {
						// Swap
						temp := stackList[index]
						stackList[index] = stackList[i]
						stackList[i] = temp
						break
					}
				}
			}
		}
	}


	// BIGGEST MONEY EVER
	commits, err = getCommitsChronological()
	if (err != nil) {
		return "", err
	}
	outputString := "gitGraph\n"
	// outputString += "  commit\n"
	for _, commit := range commits {
		for index, stack := range stackList {
			if slices.Contains(stack, commit) {
				if stack[len(stack) - 1] == commit {
<<<<<<< HEAD
					outputString += "  branch b" + strconv.Itoa(index) + "\n"
				}
				outputString += "  checkout b" + strconv.Itoa(index) + "\n"
				if len(commitMap[commit].mergeHashes) > 0 {
					for _, mergeCommit := range commitMap[commit].mergeHashes {
						branch := "b"
						for index, stack := range stackList {
							if slices.Contains(stack, mergeCommit) {
								branch += strconv.Itoa(index)
								break
							}
						}
						outputString += "  merge " + branch  + " id: \"" + commit + "\"\n"
					}
				} else {
					outputString += "  commit id: \"" + commit + "\"\n"
				}
			}
		}
	}

	// var parsedOutput string
	return outputString, nil
}
