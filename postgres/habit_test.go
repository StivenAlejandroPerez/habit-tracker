package postgres

import (
	"context"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"habit-tracker"
)

func TestHabitRepository_InsertHabits(t *testing.T) {
	type fields struct {
		db Drivers
	}
	type args struct {
		ctx    context.Context
		habits habit_tracker.Habits
		now    time.Time
	}
	type test struct {
		name    string
		fields  fields
		args    args
		wantErr assert.ErrorAssertionFunc
	}

	query := `
INSERT
	INTO
	habits
(category_id, "name", description, created_at, updated_at)
VALUES (1, 'Exercise', 'New description', '2023-07-30T12:00:00Z', '2023-07-30T12:00:00Z');`

	tests := []test{
		func() test {
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)

			mock.ExpectExec(regexp.QuoteMeta(query)).WillReturnResult(sqlmock.NewResult(1, 1))

			return test{
				name: "Success",
				fields: fields{
					db: &Postgres{
						db: db,
					},
				},
				args: args{
					ctx: context.Background(),
					habits: habit_tracker.Habits{
						{
							CategoryID:  1,
							Name:        "Exercise",
							Description: "New description",
						},
					},
					now: time.Date(2023, 7, 30, 12, 0, 0, 0, time.UTC),
				},
				wantErr: assert.NoError,
			}
		}(),
		func() test {
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)

			mock.ExpectExec(regexp.QuoteMeta(query)).WillReturnError(assert.AnError)

			return test{
				name: "ErrorExecContext",
				fields: fields{
					db: &Postgres{
						db: db,
					},
				},
				args: args{
					ctx: context.Background(),
					habits: habit_tracker.Habits{
						{
							CategoryID:  1,
							Name:        "Exercise",
							Description: "New description",
						},
					},
					now: time.Date(2023, 7, 30, 12, 0, 0, 0, time.UTC),
				},
				wantErr: assert.Error,
			}
		}(),
	}
	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				hr := NewHabitRepository(tt.fields.db)

				tt.wantErr(t, hr.InsertHabits(tt.args.ctx, tt.args.habits, tt.args.now))
			},
		)
	}
}

func TestHabitRepository_InsertHabitCategories(t *testing.T) {
	type fields struct {
		db Drivers
	}
	type args struct {
		ctx             context.Context
		habitCategories habit_tracker.HabitCategories
		now             time.Time
	}
	type test struct {
		name    string
		fields  fields
		args    args
		wantErr assert.ErrorAssertionFunc
	}

	query := `
INSERT
	INTO
	habit_categories
(category_name, created_at, updated_at)
VALUES ('Health', '2023-07-30T12:00:00Z', '2023-07-30T12:00:00Z');`

	tests := []test{
		func() test {
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)

			mock.ExpectExec(regexp.QuoteMeta(query)).WillReturnResult(sqlmock.NewResult(1, 1))

			return test{
				name: "Success",
				fields: fields{
					db: &Postgres{
						db: db,
					},
				},
				args: args{
					ctx: context.Background(),
					habitCategories: habit_tracker.HabitCategories{
						{
							CategoryName: "Health",
						},
					},
					now: time.Date(2023, 7, 30, 12, 0, 0, 0, time.UTC),
				},
				wantErr: assert.NoError,
			}
		}(),
		func() test {
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)

			mock.ExpectExec(regexp.QuoteMeta(query)).WillReturnError(assert.AnError)

			return test{
				name: "ErrorExecContext",
				fields: fields{
					db: &Postgres{
						db: db,
					},
				},
				args: args{
					ctx: context.Background(),
					habitCategories: habit_tracker.HabitCategories{
						{
							CategoryName: "Health",
						},
					},
					now: time.Date(2023, 7, 30, 12, 0, 0, 0, time.UTC),
				},
				wantErr: assert.Error,
			}
		}(),
	}
	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				hr := &HabitRepository{
					db: tt.fields.db,
				}
				tt.wantErr(t, hr.InsertHabitCategories(tt.args.ctx, tt.args.habitCategories, tt.args.now))
			},
		)
	}
}

func TestHabitRepository_InsertHabitRecords(t *testing.T) {
	type fields struct {
		db Drivers
	}
	type args struct {
		ctx          context.Context
		habitRecords habit_tracker.HabitRecords
		now          time.Time
	}
	type test struct {
		name    string
		fields  fields
		args    args
		wantErr assert.ErrorAssertionFunc
	}

	query := `
INSERT
	INTO
	habit_records
(habit_id, record_date, "result", description, created_at, updated_at)
VALUES (1, '2023-07-30T12:00:00Z', 'Success', 'New description', '2023-07-30T12:00:00Z', '2023-07-30T12:00:00Z');`

	tests := []test{
		func() test {
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)

			mock.ExpectExec(regexp.QuoteMeta(query)).WillReturnResult(sqlmock.NewResult(1, 1))

			return test{
				name: "Success",
				fields: fields{
					db: &Postgres{
						db: db,
					},
				},
				args: args{
					ctx: context.Background(),
					habitRecords: habit_tracker.HabitRecords{
						{
							HabitID:     1,
							RecordDate:  time.Date(2023, 7, 30, 12, 0, 0, 0, time.UTC),
							Result:      "Success",
							Description: "New description",
						},
					},
					now: time.Date(2023, 7, 30, 12, 0, 0, 0, time.UTC),
				},
				wantErr: assert.NoError,
			}
		}(),
		func() test {
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)

			mock.ExpectExec(regexp.QuoteMeta(query)).WillReturnError(assert.AnError)

			return test{
				name: "ErrorExecContext",
				fields: fields{
					db: &Postgres{
						db: db,
					},
				},
				args: args{
					ctx: context.Background(),
					habitRecords: habit_tracker.HabitRecords{
						{
							HabitID:     1,
							RecordDate:  time.Date(2023, 7, 30, 12, 0, 0, 0, time.UTC),
							Result:      "Success",
							Description: "New description",
						},
					},
					now: time.Date(2023, 7, 30, 12, 0, 0, 0, time.UTC),
				},
				wantErr: assert.Error,
			}
		}(),
	}
	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				hr := &HabitRepository{
					db: tt.fields.db,
				}
				tt.wantErr(t, hr.InsertHabitRecords(tt.args.ctx, tt.args.habitRecords, tt.args.now))
			},
		)
	}
}

func Test_buildInsertHabitsQuery(t *testing.T) {
	type args struct {
		tags habit_tracker.Habits
		now  time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Success",
			args: args{
				tags: habit_tracker.Habits{
					{
						CategoryID:  1,
						Name:        "Exercise",
						Description: "New description",
					},
				},
				now: time.Date(2023, 7, 30, 12, 0, 0, 0, time.UTC),
			},
			want: `
INSERT
	INTO
	habits
(category_id, "name", description, created_at, updated_at)
VALUES (1, 'Exercise', 'New description', '2023-07-30T12:00:00Z', '2023-07-30T12:00:00Z');`,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				assert.Equal(t, tt.want, buildInsertHabitsQuery(tt.args.tags, tt.args.now))
			},
		)
	}
}

func Test_buildInsertHabitCategoriesQuery(t *testing.T) {
	type args struct {
		categories habit_tracker.HabitCategories
		now        time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Success",
			args: args{
				categories: habit_tracker.HabitCategories{
					{
						CategoryName: "Health",
					},
				},
				now: time.Date(2023, 7, 30, 12, 0, 0, 0, time.UTC),
			},
			want: `
INSERT
	INTO
	habit_categories
(category_name, created_at, updated_at)
VALUES ('Health', '2023-07-30T12:00:00Z', '2023-07-30T12:00:00Z');`,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				assert.Equal(t, tt.want, buildInsertHabitCategoriesQuery(tt.args.categories, tt.args.now))
			},
		)
	}
}

func Test_buildInsertHabitRecordsQuery(t *testing.T) {
	type args struct {
		records habit_tracker.HabitRecords
		now     time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Success",
			args: args{
				records: habit_tracker.HabitRecords{
					{
						HabitID:     1,
						RecordDate:  time.Date(2023, 7, 30, 12, 0, 0, 0, time.UTC),
						Result:      "Success",
						Description: "New description",
					},
				},
				now: time.Date(2023, 7, 30, 12, 0, 0, 0, time.UTC),
			},
			want: `
INSERT
	INTO
	habit_records
(habit_id, record_date, "result", description, created_at, updated_at)
VALUES (1, '2023-07-30T12:00:00Z', 'Success', 'New description', '2023-07-30T12:00:00Z', '2023-07-30T12:00:00Z');`,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				assert.Equal(t, tt.want, buildInsertHabitRecordsQuery(tt.args.records, tt.args.now))
			},
		)
	}
}
