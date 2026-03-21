package domain

import (
	"errors"
	"time"
)

type RepoInfo struct {
	FullName    string    `json:"full_name"`
	Description string    `json:"description"`
	ForksCount  int       `json:"forks_count"`
	StarsCount  int       `json:"stargazers_count"`
	CreatedAt   time.Time `json:"created_at"`
}

type RepoRequest struct {
	Owner string
	Name  string
}

type InfoFetcher interface {
	FetchInfo(RepoRequest) (RepoInfo, error)
}

var (
	NotFoundError  = errors.New("repo not found")
	ForbiddenError = errors.New("forbidden")
)
