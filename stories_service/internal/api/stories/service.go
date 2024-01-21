package stories

import (
	"github.com/YANcomp/yanbackend/stories_service/internal/service"
	desc "github.com/YANcomp/yanbackend/stories_service/pkg/stories_v1"
)

type Implementation struct {
	desc.UnimplementedStoriesV1Server
	storiesService service.StoriesService
}

func NewImplementation(storiesService service.StoriesService) *Implementation {
	return &Implementation{
		storiesService: storiesService,
	}
}
