package habit_tracker

import (
	"context"
	"time"
)

type Tag struct {
	ID          uint64
	Name        string
	Description string
}

type Tags []Tag

type TagRepository interface {
	InsertTags(ctx context.Context, tags Tags, now time.Time) error
}
