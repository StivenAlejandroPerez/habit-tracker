package postgres

import (
	"context"
	"fmt"
	"strings"
	"time"

	"habit-tracker"
)

type TagRepository struct {
	db Drivers
}

func NewTagRepository(db Drivers) habit_tracker.TagRepository {
	return &TagRepository{
		db: db,
	}
}

func (tr *TagRepository) InsertTags(ctx context.Context, tags habit_tracker.Tags, now time.Time) error {
	_, err := tr.db.ExecContext(ctx, buildInsertTagsQuery(tags, now))
	if err != nil {
		return fmt.Errorf("[err:%w]", err)
	}

	return nil
}

func buildInsertTagsQuery(tags habit_tracker.Tags, now time.Time) string {
	str := strings.Builder{}
	nowStr := now.Format(time.RFC3339)
	for _, tag := range tags {
		str.WriteString(
			fmt.Sprintf("('%s', '%s', '%s', '%s')", tag.Name, tag.Description, nowStr, nowStr),
		)
	}

	return fmt.Sprintf(insertTagsQuery, strings.TrimSuffix(str.String(), ", "))
}
