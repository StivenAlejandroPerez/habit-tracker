package habit_tracker

import (
	"context"
	"time"
)

type Goal struct {
	ID          uint64
	Description string
}

type Goals []Goal

type GoalRepository interface {
	InsertGoals(ctx context.Context, habits Goals, now time.Time) error
}
