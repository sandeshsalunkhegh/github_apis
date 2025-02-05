package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const GITHUB_BASE_URL = "https://api.github.com"
const USER_NAME = "sandeshsalunkhegh"

var GITHUB_ACCESS_TOKEN string

func main() {

	err := godotenv.Load("../secrets.env")
	if err != nil {
		log.Fatalln(err)
	}

	GITHUB_ACCESS_TOKEN = os.Getenv("GITHUB_ACCESS_TOKEN")
	Bearer := fmt.Sprintf("Bearer %s", GITHUB_ACCESS_TOKEN)

	repositories := get_repositories(&http.Client{}, Bearer)

	for i := range repositories {

		fmt.Println(repositories[i])

		pull_requests := get_pull_requests(&http.Client{}, Bearer, repositories[i])
		if len(pull_requests) > 0 {
			fmt.Println("Pull Requests for Repository:\t", repositories[i].Name)
			for _, pull_request := range pull_requests {
				fmt.Println(pull_request)
			}
		} else {
			fmt.Println("No Pull Requests for Repository:\t", repositories[i].Name)
		}

		commits := get_commits(&http.Client{}, Bearer, repositories[i])
		if len(commits) > 0 {
			fmt.Println("Commits for Repository:\t", repositories[i].Name)
			for _, commit := range commits {
				fmt.Println(commit)
			}
		} else {
			fmt.Println("No Commits for Repository:\t", repositories[i].Name)
		}

		branches := get_branches(&http.Client{}, Bearer, repositories[i])
		if len(branches) > 0 {
			fmt.Println("Branches for Repository:\t", repositories[i].Name)
			for _, branch := range branches {
				fmt.Println(branch)
			}
		} else {
			fmt.Println("No Branches for Repository:\t", repositories[i].Name)
		}
	}
}
