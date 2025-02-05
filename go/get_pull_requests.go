package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func get_pull_requests(client *http.Client, Bearer string, repository Repository) []PullRequest {

	fetch_pull_request_url := fmt.Sprintf("%s/repos/%s/%s/pulls", GITHUB_BASE_URL, USER_NAME, repository.Name)
	req, err := http.NewRequest("GET", fetch_pull_request_url, nil)
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

	var pull_requests []PullRequest

	err = json.Unmarshal(body, &pull_requests)
	if err != nil {
		log.Println("Error while marshalling the response bytes:", err)
	}

	return pull_requests
}
