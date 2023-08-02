package postgres

import (
	"context"
	"fmt"
	"strings"
	"time"

	"habit-tracker"
)

type HabitRepository struct {
	db Drivers
}

func NewHabitRepository(db Drivers) habit_tracker.HabitRepository {
	return &HabitRepository{
		db: db,
	}
}

func (hr *HabitRepository) InsertHabits(ctx context.Context,
	habits habit_tracker.Habits, now time.Time) error {
	_, err := hr.db.ExecContext(ctx, buildInsertHabitsQuery(habits, now))
	if err != nil {
		return fmt.Errorf("[err:%w]", err)
	}

	return nil
}

func (hr *HabitRepository) InsertHabitCategories(ctx context.Context,
	habitCategories habit_tracker.HabitCategories, now time.Time) error {
	_, err := hr.db.ExecContext(ctx, buildInsertHabitCategoriesQuery(habitCategories, now))
	if err != nil {
		return fmt.Errorf("[err:%w]", err)
	}

	return nil
}

func (hr *HabitRepository) InsertHabitRecords(ctx context.Context,
	habitRecords habit_tracker.HabitRecords, now time.Time) error {
	_, err := hr.db.ExecContext(ctx, buildInsertHabitRecordsQuery(habitRecords, now))
	if err != nil {
		return fmt.Errorf("[err:%w]", err)
	}

	return nil
}

func (hr *HabitRepository) UpdateHabits(ctx context.Context,
	habits habit_tracker.Habits, now time.Time) error {
	return nil
}

func (hr *HabitRepository) UpdateHabitCategories(ctx context.Context,
	habitCategories habit_tracker.HabitCategories, now time.Time) error {
	return nil
}

func (hr *HabitRepository) UpdateHabitRecords(ctx context.Context,
	habitRecords habit_tracker.HabitRecords, now time.Time) error {
	return nil
}

func buildInsertHabitsQuery(habits habit_tracker.Habits, now time.Time) string {
	str := strings.Builder{}
	nowStr := now.Format(time.RFC3339)
	for _, habit := range habits {
		str.WriteString(
			fmt.Sprintf(
				"(%d, '%s', '%s', '%s', '%s')", habit.CategoryID, habit.Name, habit.Description, nowStr, nowStr,
			),
		)
	}

	return fmt.Sprintf(insertHabitsQuery, strings.TrimSuffix(str.String(), ", "))
}

func buildInsertHabitCategoriesQuery(categories habit_tracker.HabitCategories, now time.Time) string {
	str := strings.Builder{}
	nowStr := now.Format(time.RFC3339)
	for _, category := range categories {
		str.WriteString(
			fmt.Sprintf("('%s', '%s', '%s')", category.CategoryName, nowStr, nowStr),
		)
	}

	return fmt.Sprintf(insertHabitCategoriesQuery, strings.TrimSuffix(str.String(), ", "))
}

func buildInsertHabitRecordsQuery(records habit_tracker.HabitRecords, now time.Time) string {
	str := strings.Builder{}
	nowStr := now.Format(time.RFC3339)
	for _, record := range records {
		str.WriteString(
			fmt.Sprintf(
				"(%d, '%s', '%s', '%s', '%s', '%s')",
				record.HabitID,
				record.RecordDate.Format(time.RFC3339),
				record.Result,
				record.Description,
				nowStr,
				nowStr,
			),
		)
	}

	return fmt.Sprintf(insertHabitRecordsQuery, strings.TrimSuffix(str.String(), ", "))
}
