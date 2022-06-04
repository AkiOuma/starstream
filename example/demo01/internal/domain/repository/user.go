package repository

import (
	"context"
	"github.com/AkiOuma/demo01/internal/domain/entity"
	"github.com/AkiOuma/demo01/internal/domain/valueobject"
)

type User interface {
	FindUsers(ctx context.Context, querier *valueobject.UserQuerier) (*FindUsersResult, error)
	SaveUsers(ctx context.Context, data []*entity.User) ([]*entity.User, error)
	RemoveUsers(ctx context.Context, root []int) error
}

// Finding Results
type FindUsersResult struct {
	Count int
	Data  []*entity.User
}
