package usecase

import (
	"context"
	"github.com/AkiOuma/demo01/internal/domain/entity"
	"github.com/AkiOuma/demo01/internal/domain/repository"
	"github.com/AkiOuma/demo01/internal/domain/service"
	"github.com/AkiOuma/demo01/internal/domain/valueobject"
)

// Interface
type User interface {
	FindUsers(ctx context.Context, querier *valueobject.UserQuerier) (*repository.FindUsersResult, error)
	SaveUsers(ctx context.Context, data []*entity.User) ([]*entity.User, error)
	RemoveUsers(ctx context.Context, querier *valueobject.UserQuerier) error
}

// Implement
type user struct {
	repo repository.User
	svc  service.User
}

var _ User = (*user)(nil)

// Constructor
func NewUser(
	repo repository.User,
	svc service.User,
) User {
	return &user{svc: svc, repo: repo}
}

// Implement methods
func (u *user) FindUsers(ctx context.Context, querier *valueobject.UserQuerier) (*repository.FindUsersResult, error) {
	return u.repo.FindUsers(ctx, querier)
}
func (u *user) SaveUsers(ctx context.Context, data []*entity.User) ([]*entity.User, error) {
	return u.repo.SaveUsers(ctx, data)
}
func (u *user) RemoveUsers(ctx context.Context, querier *valueobject.UserQuerier) error {
	result, err := u.FindUsers(ctx, querier)
	if err != nil {
		return err
	}
	root := make([]int, 0, result.Count)
	for _, v := range result.Data {
		root = append(root, v.GetId())
	}
	return u.repo.RemoveUsers(ctx, root)

}
