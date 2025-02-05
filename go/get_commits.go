package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func get_commits(client *http.Client, Bearer string, repository Repository) []Commit {

	fetch_commits_url := fmt.Sprintf("%s/repos/%s/%s/commits", GITHUB_BASE_URL, USER_NAME, repository.Name)
	req, err := http.NewRequest("GET", fetch_commits_url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("Authorization", Bearer)

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}

	var commits []Commit

	err = json.Unmarshal(body, &commits)
	if err != nil {
		log.Println("Error while marshalling the response bytes:", err)
	}

	return commits
}
