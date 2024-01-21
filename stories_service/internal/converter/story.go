package converter

import (
	"github.com/YANcomp/yanbackend/stories_service/internal/model"
	desc "github.com/YANcomp/yanbackend/stories_service/pkg/stories_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToStoryFromService(story *model.Story) *desc.Story {
	var updatedAt *timestamppb.Timestamp
	if story.UpdatedAt.Valid {
		updatedAt = timestamppb.New(story.UpdatedAt.Time)
	}

	return &desc.Story{
		ID:             story.ID,
		IsActive:       story.IsActive,
		IsActiveMobile: story.IsActiveMobile,
		Preview:        story.Preview,
		Slides:         ToSlidesFromService(story.Slides),
		Title:          story.Title,
		CreatedAt:      timestamppb.New(story.CreatedAt),
		UpdatedAt:      updatedAt,
	}
}

func ToSlidesFromService(slides []*model.Slide) []*desc.Slide {
	slideList := make([]*desc.Slide, 0, len(slides))
	for _, slide := range slides {
		slideList = append(slideList, &desc.Slide{
			ID:                 slide.ID,
			BackgroundImage:    slide.BackgroundImage,
			Caption:            slide.Caption,
			Content:            slide.Content,
			Delay:              slide.Delay,
			IsHideShadowBottom: slide.IsHideShadowBottom,
			TextPosition:       slide.TextPosition,
		})
	}

	return slideList
}
