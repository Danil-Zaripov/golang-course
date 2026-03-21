package adapters

import (
	"context"
	"time"

	pb "github.com/Danil-Zaripov/golang-course/project/api/v1"
	"github.com/Danil-Zaripov/golang-course/project/gateway/internal/domain"
)

type CollectorAdapter struct {
	client pb.CollectorClient
}

func (c *CollectorAdapter) SendAndGet(p domain.RepoPath) (domain.RepoInfo, error) {
	ret := domain.RepoInfo{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*6)
	defer cancel()
	r, err := c.client.GetRepoData(ctx, &pb.RepoRequest{Name: p.Name, Owner: p.Owner})
	if err != nil {
		return ret, err
	}
	ret.FullName = r.Fullname

	ret.CreatedAt, err = time.Parse("2006-01-02", r.CreatedAt)
	if err != nil {
		return ret, err
	}
	ret.Description = r.Desc
	ret.ForksCount = int(r.ForksCount)
	ret.StarsCount = int(r.StarsCount)
	return ret, nil
}

func NewCollectorAdapter(c pb.CollectorClient) *CollectorAdapter {
	return &CollectorAdapter{client: c}
}
