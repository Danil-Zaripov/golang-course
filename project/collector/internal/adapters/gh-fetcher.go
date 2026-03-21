package adapters

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/Danil-Zaripov/golang-course/project/collector/internal/domain"
)

type GHFetcher struct {
	client *http.Client
}

func (gf *GHFetcher) FetchInfo(request domain.RepoRequest) (domain.RepoInfo, error) {
	info := domain.RepoInfo{}
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s", request.Owner, request.Name)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return info, err
	}
	resp, err := gf.client.Do(req)
	if err != nil {
		return info, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		switch resp.StatusCode {
		case http.StatusNotFound:
			return info, domain.NotFoundError
		case http.StatusForbidden:
			return info, domain.ForbiddenError
		default:
			return info, errors.New(resp.Status)
		}
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return info, err
	}
	err = json.Unmarshal(body, &info)
	if err != nil {
		return info, err
	}
	return info, nil
}

func NewGHFetcher(client *http.Client) *GHFetcher {
	return &GHFetcher{client: client}
}
