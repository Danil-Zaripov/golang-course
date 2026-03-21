package usecases

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/Danil-Zaripov/golang-course/project/gateway/internal/domain"
)

type RepoUsecase struct {
	forwarder domain.DataForwarder
}

func processUrl(u *url.URL) (domain.RepoPath, error) {
	ret := domain.RepoPath{}
	path := u.Path
	words := strings.Split(path, "/")
	if len(words) != 5 {
		path = strings.Join(words[3:], "/")
		return ret, fmt.Errorf("incorrect repo path %s", path)
	}
	ret.Owner = words[3]
	ret.Name = words[4]
	return ret, nil
}

func (r *RepoUsecase) ProcessAndForward(u *url.URL) (domain.RepoInfo, error) {
	ret := domain.RepoInfo{}
	path, err := processUrl(u)
	if err != nil {
		return ret, err
	}
	return r.forwarder.SendAndGet(path)
}

func NewRepoUsecase(f domain.DataForwarder) *RepoUsecase {
	return &RepoUsecase{forwarder: f}
}
