package domain

import "time"

type RepoInfo struct {
	FullName    string    `json:"full_name"`
	Description string    `json:"description"`
	ForksCount  int       `json:"forks_count"`
	StarsCount  int       `json:"stargazers_count"`
	CreatedAt   time.Time `json:"created_at"`
}

type RepoPath struct {
	Owner string
	Name  string
}

type DataForwarder interface {
	SendAndGet(RepoPath) (RepoInfo, error)
}
