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

//go:generate mockery --name TagRepository --filename tag_repository.go --outpkg mocks --structname TagRepository --disable-version-string
type TagRepository interface {
	InsertTags(ctx context.Context, tags Tags, now time.Time) error
	UpdateTags(ctx context.Context, tags Tags, now time.Time) error
}
