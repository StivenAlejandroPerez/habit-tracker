package habit_tracker

import (
	"context"
	"time"
)

type Event struct {
	ID      uint64
	HabitID uint64
	Subject string
	StartAt time.Time
	EndAt   time.Time
}

type Events []Event

//go:generate mockery --name EventRepository --filename event_repository.go --outpkg mocks --structname EventRepository --disable-version-string
type EventRepository interface {
	InsertEvents(ctx context.Context, events Events, now time.Time) error
	UpdateEvents(ctx context.Context, events Events, now time.Time) error
}
