package service

import (
    "context"
    "mytemplate/internal/model"
    "mytemplate/internal/repository"
)

type UserService struct {
    Repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
    return &UserService{Repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, user *model.User) error {
    return s.Repo.Create(ctx, user)
}
