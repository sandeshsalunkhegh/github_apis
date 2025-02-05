package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func get_repositories(client *http.Client, Bearer string) []Repository {

	fetch_repositories_url := fmt.Sprintf("%s/users/%s/repos", GITHUB_BASE_URL, USER_NAME)

	req, err := http.NewRequest("GET", fetch_repositories_url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("Authorization", Bearer)
	req.Header.Set("Accept", "application/vnd.github+json")

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}

	var repositories []Repository

	err = json.Unmarshal(body, &repositories)
	if err != nil {
		log.Println("Error while marshalling the response bytes:", err)
	}

	return repositories
}
