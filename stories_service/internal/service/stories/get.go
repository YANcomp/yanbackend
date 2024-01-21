package stories

import (
	"context"
	"github.com/YANcomp/yanbackend/stories_service/internal/model"
)

func (s *serv) Get(ctx context.Context, id int64) (*model.Story, error) {
	story, err := s.storiesRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return story, nil
}
