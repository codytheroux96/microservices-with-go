package metadata

import (
	"context"
	"errors"

	"movieexample.com/metadata/internal/repository"
	"movieexample.com/metadata/pkg/model"
)

var ErrNotFound = errors.New("not found")

type metadataRepository interface {
	Get(ctx context.Context, id string) (*model.Metadata, error)
}

// Defines a metadata service controller
type Controller struct {
	repo metadataRepository
}

// Creates new metadata service controller
func New(repo metadataRepository) *Controller {
	return &Controller{repo}
}

func (c *Controller) Get(ctx context.Context, id string) (*model.Metadata, error) {
	res, err := c.repo.Get(ctx, id)
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		return nil, ErrNotFound
	}

	return res, err
}
