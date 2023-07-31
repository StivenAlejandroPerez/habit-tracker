package postgres

import (
	"context"
	"fmt"
	"strings"
	"time"

	"habit-tracker"
)

type GoalRepository struct {
	db Drivers
}

func NewGoalRepository(db Drivers) habit_tracker.GoalRepository {
	return &GoalRepository{
		db: db,
	}
}

func (gr *GoalRepository) InsertGoals(ctx context.Context, goals habit_tracker.Goals, now time.Time) error {
	_, err := gr.db.ExecContext(ctx, buildInsertGoalsQuery(goals, now))
	if err != nil {
		return fmt.Errorf("[err:%w]", err)
	}

	return nil
}

func buildInsertGoalsQuery(goals habit_tracker.Goals, now time.Time) string {
	str := strings.Builder{}
	nowStr := now.Format(time.RFC3339)
	for _, goal := range goals {
		str.WriteString(fmt.Sprintf("('%s', '%s', '%s'), ", goal.Description, nowStr, nowStr))
	}

	return fmt.Sprintf(insertGoalsQuery, strings.TrimSuffix(str.String(), ", "))
}
