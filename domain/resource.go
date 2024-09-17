package domain

import (
	"context"
	"time"
)

type Resource struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}

type ResourceCriteria struct {
	Name *string `json:"name"`
	URL  *string `json:"url"`
}

type ResourceUseCase interface {
	Post(ctx context.Context, ctr *Resource) (*ResourceCriteria, error)
	Get(ctx context.Context) ([]*Resource, error)
}

type ResourceRepository interface {
	Post(ctx context.Context, ctr *Resource) (*ResourceCriteria, error)
	Get(ctx context.Context) ([]*Resource, error)
}
