package postgres

import (
	"context"
	"fmt"
	"strings"
	"time"

	"habit-tracker"
)

type EventRepository struct {
	db Drivers
}

func NewEventRepository(db Drivers) habit_tracker.EventRepository {
	return &EventRepository{
		db: db,
	}
}

func (er *EventRepository) InsertEvents(ctx context.Context, events habit_tracker.Events, now time.Time) error {
	_, err := er.db.ExecContext(ctx, buildInsertEventsQuery(events, now))
	if err != nil {
		return fmt.Errorf("[err:%w]", err)
	}

	return nil
}

func buildInsertEventsQuery(events habit_tracker.Events, now time.Time) string {
	str := strings.Builder{}
	nowStr := now.Format(time.RFC3339)
	for _, event := range events {
		str.WriteString(
			fmt.Sprintf(
				`(%d, '%s', '%s', '%s', '%s', '%s'), `,
				event.HabitID, event.Subject, event.StartAt.Format(time.RFC3339),
				event.EndAt.Format(time.RFC3339), nowStr, nowStr,
			),
		)
	}

	return fmt.Sprintf(insertEventsQuery, strings.TrimSuffix(str.String(), ", "))
}
