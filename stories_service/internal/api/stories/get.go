package stories

import (
	"context"
	"github.com/YANcomp/yanbackend/stories_service/internal/converter"
	desc "github.com/YANcomp/yanbackend/stories_service/pkg/stories_v1"
	"log"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	storyObj, err := i.storiesService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	log.Printf("id: %d", storyObj.ID)

	return &desc.GetResponse{
		Story: converter.ToStoryFromService(storyObj),
	}, nil
}
