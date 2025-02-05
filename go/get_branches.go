package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func get_branches(client *http.Client, Bearer string, repository Repository) []Branch {

	fetch_branches_url := fmt.Sprintf("%s/repos/%s/%s/branches", GITHUB_BASE_URL, USER_NAME, repository.Name)
	req, err := http.NewRequest("GET", fetch_branches_url, nil)
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

	var branches []Branch

	err = json.Unmarshal(body, &branches)
	if err != nil {
		log.Println("Error while marshalling the response bytes:", err)
	}

	return branches
}
