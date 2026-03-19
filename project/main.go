package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const UsageInformation string = `Usage: go run . OWNER/REPO
Example: go run . Danil-Zaripov/golang-course`

type ParsedJsonData struct {
	Full_Name, Description        string
	Stargazers_count, Forks_count int
	Created_At                    time.Time
}

func main() {
	argsWithoutProc := os.Args[1:]
	argc := len(argsWithoutProc)
	if argc != 1 {
		fmt.Println(UsageInformation)
		return
	}
	repo_name := argsWithoutProc[0]
	words := strings.Split(repo_name, "/")
	if len(words) != 2 {
		fmt.Printf("Incorrect repository name %s\n\n%s", repo_name, UsageInformation)
		return
	}
	client := &http.Client{}
	url := fmt.Sprintf("https://api.github.com/repos/%s", repo_name)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.Status)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	var data ParsedJsonData
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf(
		`Repository: %s
Description: %s
%d stars
%d forks
Created at: %s
`,
		data.Full_Name,
		data.Description,
		data.Stargazers_count,
		data.Forks_count,
		data.Created_At.Format("2006-01-02"),
	)
}
