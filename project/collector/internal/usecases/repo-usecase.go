package usecases

import (
	pb "github.com/Danil-Zaripov/golang-course/project/api/v1"
	"github.com/Danil-Zaripov/golang-course/project/collector/internal/domain"
)

type RepoUsecase struct {
	fetcher domain.InfoFetcher
}

func NewRepoUsecase(fetcher domain.InfoFetcher) *RepoUsecase {
	return &RepoUsecase{fetcher: fetcher}
}

func (ru *RepoUsecase) ProcessAndFetch(request *pb.RepoRequest) (domain.RepoInfo, error) {
	drequest := domain.RepoRequest{
		Owner: request.GetOwner(),
		Name:  request.GetName(),
	}
	return ru.fetcher.FetchInfo(drequest)
}
