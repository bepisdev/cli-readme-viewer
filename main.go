package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Repository struct {
	Content string `json:"content"`
}

func main() {
	// FIXME: Replace this with CLI flags later
	repoOwner := "joshburnsxyz"
	repoName := "protorec"

	// Perform GET request to github API for repos readme file.
	resp, err := http.Get(fmt.Sprintf("https://api.github.com/repos/%s/%s/readme",repoOwner,repoName))

	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// Extract JSON response
	body,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Unmarshal response
	var repo Repository
	err = json.Unmarshal(body, &repo)
	if err != nil {
		fmt.Println(err)
		return
	}

	data, err := base64.StdEncoding.DecodeString(repo.Content)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print result
	fmt.Println(string(data))

	return
}
