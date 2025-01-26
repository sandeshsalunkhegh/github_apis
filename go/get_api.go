package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
)

func main() {
	client := github.NewClient(nil)
	context := context.Background()
	repositories, response, err := client.Repositories.Get(context, "sandeshsalunkhegh", "go_apis")
	if err != nil {
		fmt.Println("Error:\t", err)
	}
	fmt.Println(repositories.GetPullsURL())
	fmt.Println(response.Request.GetBody())
}
