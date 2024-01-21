package repository

import (
	"context"
	"github.com/YANcomp/yanbackend/stories_service/internal/model"
)

type StoriesRepository interface {
	Get(ctx context.Context, id int64) (*model.Story, error)
}
