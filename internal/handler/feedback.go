package handler

import (
	"context"

	fs "github.com/wintermonth2298/feedback-service/api/feedback-service"
	"github.com/wintermonth2298/feedback-service/internal/model"
)

type FeedbackService interface {
	Create(ctx context.Context, feedback model.Feedback) error
	Get(ctx context.Context, feedbackID int64) (model.Feedback, error)
}

type FeedbackHandler struct {
	feedbackService FeedbackService
	fs.UnimplementedFeedbackServiceServer
}

func NewFeedbackHandler(feedbackService FeedbackService) *FeedbackHandler {
	return &FeedbackHandler{feedbackService: feedbackService}
}

func (h *FeedbackHandler) CreateFeedback(ctx context.Context, req *fs.CreateFeedbackRequest) (*fs.CreateFeedbackResponse, error) {
	feedback := model.Feedback{
		ID:   req.GetId(),
		Rate: req.GetRate(),
		Text: req.GetText(),
	}

	err := h.feedbackService.Create(ctx, feedback)
	return &fs.CreateFeedbackResponse{}, err
}

func (h *FeedbackHandler) GetFeedback(ctx context.Context, req *fs.GetFeedbackRequest) (*fs.GetFeedbackResponse, error) {
	feedback, err := h.feedbackService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	resp := &fs.GetFeedbackResponse{
		Id:   feedback.ID,
		Rate: feedback.Rate,
		Text: feedback.Text,
	}

	return resp, nil
}
