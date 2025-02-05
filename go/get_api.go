package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/sandeshsalunkhegh/github_apis/go/models"

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

	organization_repositories_url := fmt.Sprintf("%s/users/%s/repos", GITHUB_BASE_URL, USER_NAME)
	req, err := http.NewRequest("GET", organization_repositories_url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("Authorization", Bearer)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}
	// repositories := string([]byte(body))
	log.Println(json.Valid(body))
	var repositories []models.Repository
	err = json.Unmarshal(body, &repositories)
	if err != nil {
		log.Println("Error while marshalling the response bytes:", err)
	}
	for i := range repositories {
		fmt.Println(repositories[i])
	}
}
