package stories

import (
	"github.com/YANcomp/yanbackend/stories_service/internal/repository"
	"github.com/YANcomp/yanbackend/stories_service/internal/service"
)

type serv struct {
	storiesRepository repository.StoriesRepository
}

func NewService(
	storiesRepository repository.StoriesRepository,
) service.StoriesService {
	return &serv{
		storiesRepository: storiesRepository,
	}
}
