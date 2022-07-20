package repository

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/wintermonth2298/feedback-service/internal/model"
)

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

type FeedbackRepository struct {
	db *sqlx.DB
}

func NewFeedbackRepository(db *sqlx.DB) *FeedbackRepository {
	return &FeedbackRepository{db: db}
}

func (r *FeedbackRepository) Create(ctx context.Context, feedback model.Feedback) error {
	query, args, err := psql.Insert("feedback").Columns("rate, text").
		Values(feedback.Rate, feedback.Text).ToSql()
	if err != nil {
		return fmt.Errorf("feedbackRepository - Create() - sq: %w", err)
	}

	_, err = r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("feedbackRepository - Create() - ExecContext(): %w", err)
	}

	return nil
}

func (r *FeedbackRepository) Get(ctx context.Context, feedbackID int64) (model.Feedback, error) {
	var feedback model.Feedback

	query, args, err := psql.Select("*").From("feedback").Where(sq.Eq{"id": feedbackID}).ToSql()
	if err != nil {
		return model.Feedback{}, fmt.Errorf("feedbackRepository - Get() - sq: %w", err)
	}

	err = r.db.QueryRowxContext(ctx, query, args...).StructScan(&feedback)
	if err != nil {
		return model.Feedback{}, fmt.Errorf("feedbackRepository - getAll() - QueryRowxContext(): %w", err)
	}

	return feedback, nil
}
