package service

import (
	"context"

	"github.com/wintermonth2298/feedback-service/internal/model"
)

type FeedbackRepository interface {
	Create(ctx context.Context, feedback model.Feedback) error
	Get(ctx context.Context, feedbackID int64) (model.Feedback, error)
}

type FeedbackService struct {
	feedbackRepo FeedbackRepository
}

func NewFeedbackService(feedbackRepo FeedbackRepository) *FeedbackService {
	return &FeedbackService{feedbackRepo: feedbackRepo}
}

func (f *FeedbackService) Create(ctx context.Context, feedback model.Feedback) error {
	return f.feedbackRepo.Create(ctx, feedback)
}

func (f *FeedbackService) Get(ctx context.Context, feedbackID int64) (model.Feedback, error) {
	return f.feedbackRepo.Get(ctx, feedbackID)
}
