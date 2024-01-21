package stories

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/YANcomp/platform_common/pkg/db"
	"github.com/YANcomp/yanbackend/stories_service/internal/model"
	"github.com/YANcomp/yanbackend/stories_service/internal/repository"
	"github.com/YANcomp/yanbackend/stories_service/internal/repository/stories/converter"
	modelRepo "github.com/YANcomp/yanbackend/stories_service/internal/repository/stories/model"
)

const (
	tableName = "stories"

	idColumn             = "id"
	previewColumn        = "preview"
	titleColumn          = "title"
	isActiveColumn       = "isActive"
	isActiveMobileColumn = "isActiveMobile"
	createdAtColumn      = "created_at"
	updatedAtColumn      = "updated_at"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.StoriesRepository {
	return &repo{db: db}
}

func (r *repo) Get(ctx context.Context, id int64) (*model.Story, error) {
	builder := sq.Select(idColumn, previewColumn, titleColumn, isActiveColumn, isActiveMobileColumn, createdAtColumn, updatedAtColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{idColumn: id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "stories_repository.Get",
		QueryRaw: query,
	}

	var story modelRepo.Story
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&story.ID, &story.Preview, &story.Title, &story.IsActive,
		&story.IsActiveMobile, &story.CreatedAt, &story.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return converter.ToStoryFromRepo(&story), nil
}
