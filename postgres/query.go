package postgres

// Inserts
const (
	insertEventsQuery = `
INSERT
	INTO
	events
(habit_id, subject, start_at, end_at, created_at, updated_at)
VALUES %s;`
	insertGoalsQuery = `
INSERT
	INTO
	goals
(description, created_at, updated_at)
VALUES %s;`
	insertTagsQuery = `
INSERT
	INTO
	tags
(name, description, created_at, updated_at)
VALUES %s;`
	insertHabitsQuery = `
INSERT
	INTO
	habits
(category_id, "name", description, created_at, updated_at)
VALUES %s;`
	insertHabitCategoriesQuery = `
INSERT
	INTO
	habit_categories
(category_name, created_at, updated_at)
VALUES %s;`
	insertHabitRecordsQuery = `
INSERT
	INTO
	habit_records
(habit_id, record_date, "result", description, created_at, updated_at)
VALUES %s;`
)

// Updates
const (
	UpdateEventsQuery = `
UPDATE
	events
SET
	habit_id = %d,
	subject = '%s',
	start_at = '%s',
	end_at = '%s',
	updated_at = '%s'
WHERE
	id = %d;`
	UpdateGoalsQuery = `
UPDATE
	goals
SET
	description = '%s',
	updated_at = '%s'
WHERE
	id = %d;`
	UpdateTagsQuery = `
UPDATE
	tags
SET
	"name" = '%s',
	description = '%s',
	updated_at = '%s'
WHERE
	id = %d;`
)
