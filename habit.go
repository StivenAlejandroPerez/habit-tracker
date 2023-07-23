package habit_tracker

import (
	"context"
	"time"
)

type Habit struct {
	ID          uint64 `sql:"id"`
	CategoryID  uint64 `sql:"category_id"`
	Name        string `sql:"name"`
	Description string `sql:"description"`
}

type HabitCategory struct {
	ID           uint64 `sql:"id"`
	CategoryName string `sql:"category_name"`
}

type HabitRecord struct {
	ID          uint64    `sql:"id"`
	HabitID     uint64    `sql:"habit_id"`
	RecordDate  time.Time `sql:"record_date"`
	Result      string    `sql:"result"`
	Description string    `sql:"description"`
}

type Habits []Habit
type HabitCategories []HabitCategory
type HabitRecords []HabitRecord

type HabitRepository interface {
	InsertHabits(ctx context.Context, habits Habits, now time.Time) error
	InsertHabitCategories(ctx context.Context, habitCategories HabitCategories, now time.Time) error
	InsertHabitRecords(ctx context.Context, habitRecords HabitRecords, now time.Time) error
}
