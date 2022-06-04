package service

import (
	"github.com/AkiOuma/demo01/internal/domain/repository"
)

type User interface{}

var _ User = (*user)(nil)

type user struct {
	repo repository.User
}
