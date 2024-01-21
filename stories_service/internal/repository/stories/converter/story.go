package converter

import (
	"github.com/YANcomp/yanbackend/stories_service/internal/model"
	modelRepo "github.com/YANcomp/yanbackend/stories_service/internal/repository/stories/model"
)

func ToStoryFromRepo(note *modelRepo.Story) *model.Story {
	return &model.Story{
		ID:             note.ID,
		IsActive:       note.IsActive,
		IsActiveMobile: note.IsActiveMobile,
		Preview:        note.Preview,
		Slides:         ToSlidesFromRepo(note.Slides),
		Title:          note.Title,
		CreatedAt:      note.CreatedAt,
		UpdatedAt:      note.UpdatedAt,
	}
}

func ToSlidesFromRepo(slides []*modelRepo.Slide) []*model.Slide {
	slideList := make([]*model.Slide, 0, len(slides))
	for _, slide := range slides {
		slideList = append(slideList, &model.Slide{
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
