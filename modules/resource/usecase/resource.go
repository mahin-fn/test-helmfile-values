package usecase

import (
	"context"

	"github.com/mokhlesurr031/sre-project/backend/domain"
)

// New return new usecase for user
func New(repo domain.ResourceRepository) domain.ResourceUseCase {
	return &ResourceUseCase{
		repo: repo,
	}
}

type ResourceUseCase struct {
	repo domain.ResourceRepository
}

func (resource *ResourceUseCase) Post(ctx context.Context, ctr *domain.Resource) (*domain.ResourceCriteria, error) {
	return resource.repo.Post(ctx, ctr)
}

func (resource *ResourceUseCase) Get(ctx context.Context) ([]*domain.Resource, error) {
	return resource.repo.Get(ctx)
}
