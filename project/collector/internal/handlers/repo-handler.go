package handlers

import (
	"context"
	"errors"

	pb "github.com/Danil-Zaripov/golang-course/project/api/v1"
	"github.com/Danil-Zaripov/golang-course/project/collector/internal/domain"
	"github.com/Danil-Zaripov/golang-course/project/collector/internal/usecases"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RepoHandler struct {
	usecase *usecases.RepoUsecase
	pb.UnimplementedCollectorServer
}

func (rh *RepoHandler) GetRepoData(_ context.Context, in *pb.RepoRequest) (*pb.RepoReply, error) {
	reply := pb.RepoReply{}
	info, err := rh.usecase.ProcessAndFetch(in)
	if err != nil {
		switch {
		case errors.Is(err, domain.NotFoundError):
			return &reply, status.Error(codes.NotFound, err.Error())
		case errors.Is(err, domain.ForbiddenError):
			return &reply, status.Error(codes.PermissionDenied, err.Error())
		default:
			return &reply, err
		}
	}
	reply.Fullname = info.FullName
	reply.Desc = info.Description
	reply.ForksCount = int32(info.ForksCount)
	reply.StarsCount = int32(info.StarsCount)
	reply.CreatedAt = info.CreatedAt.Format("2006-01-02")

	return &reply, nil
}

func NewRepoHandler(usecase *usecases.RepoUsecase) *RepoHandler {
	return &RepoHandler{usecase: usecase}
}
