package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Danil-Zaripov/golang-course/project/gateway/internal/usecases"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RepoHandler struct {
	usecase *usecases.RepoUsecase
}

func NewRepoHandler(u *usecases.RepoUsecase) *RepoHandler {
	return &RepoHandler{usecase: u}
}

// @Summary      Fetch repo information
// @Produce      json
// @Param        owner path      string  true  "Owner"
// @Param        name  path      string  true  "Name"
// @Router       /api/repo/{owner}/{name} [get]
func (rh *RepoHandler) HandleRepo(w http.ResponseWriter, r *http.Request) {
	s, err := rh.usecase.ProcessAndForward(r.URL)
	if err != nil {
		code := status.Code(err)
		if code != codes.InvalidArgument {
			switch code {
			case codes.NotFound:
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("Repository not found"))
			case codes.PermissionDenied:
				w.WriteHeader(http.StatusForbidden)
				w.Write([]byte("Permission denied"))
			default:
				w.WriteHeader(http.StatusBadGateway)
				w.Write([]byte(fmt.Sprintf("%d grpc error", code)))
			}
			return
		}
		log.Printf("failed to process %v", err)
		return
	}
	j, err := json.Marshal(&s)
	if err != nil {
		log.Print(err)
		return
	}
	w.Write(j)
}
