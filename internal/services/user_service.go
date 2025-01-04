package services

import (
	"context"

	"github.com/dfanso/go-echo-boilerplate/internal/models"
	"github.com/dfanso/go-echo-boilerplate/internal/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) GetAll(ctx context.Context) ([]models.User, error) {
	return s.repo.FindAll(ctx)
}

func (s *UserService) GetByID(ctx context.Context, id primitive.ObjectID) (*models.User, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *UserService) Create(ctx context.Context, user *models.User) error {
	return s.repo.Create(ctx, user)
}

func (s *UserService) Update(ctx context.Context, user *models.User) error {
	return s.repo.Update(ctx, user)
}

func (s *UserService) Delete(ctx context.Context, id primitive.ObjectID) error {
	return s.repo.Delete(ctx, id)
}
