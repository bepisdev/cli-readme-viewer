package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"flag"
	markdown "github.com/MichaelMure/go-term-markdown"
)

type Repository struct {
	Content string `json:"content"`
}

var (
	repoOwner string
	repoName string
)

// Initialize flags
func init()  {
	flag.StringVar(&repoOwner, "owner", "", "The name of the account that owns the repo")
	flag.StringVar(&repoName, "name", "", "The name of the repo itself")
}

func main() {

	flag.Parse()
	
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
	dataStr := string(data)

	fmt.Println(string(markdown.Render(dataStr, 80, 6)))

	return
}
