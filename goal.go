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

//go:generate mockery --name GoalRepository --filename goal_repository.go --outpkg mocks --structname GoalRepository --disable-version-string
type GoalRepository interface {
	InsertGoals(ctx context.Context, habits Goals, now time.Time) error
	UpdateGoals(ctx context.Context, habits Goals, now time.Time) error
}
