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

func TestGoalRepository_InsertGoals(t *testing.T) {
	type fields struct {
		db Drivers
	}
	type args struct {
		ctx   context.Context
		goals habit_tracker.Goals
		now   time.Time
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
	goals
(description, created_at, updated_at)
VALUES ('New goal', '2023-07-20T15:32:00Z', '2023-07-20T15:32:00Z');`
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
					goals: habit_tracker.Goals{
						{
							Description: "New goal",
						},
					},
					now: time.Date(2023, 7, 20, 15, 32, 0, 0, time.UTC),
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
					goals: habit_tracker.Goals{
						{
							Description: "New goal",
						},
					},
					now: time.Date(2023, 7, 20, 15, 32, 0, 0, time.UTC),
				},
				wantErr: assert.Error,
			}
		}(),
	}
	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				gr := NewGoalRepository(tt.fields.db)
				tt.wantErr(t, gr.InsertGoals(tt.args.ctx, tt.args.goals, tt.args.now))
			},
		)
	}
}

func TestGoalRepository_UpdateGoals(t *testing.T) {
	type fields struct {
		db Drivers
	}
	type args struct {
		ctx   context.Context
		goals habit_tracker.Goals
		now   time.Time
	}
	type test struct {
		name    string
		fields  fields
		args    args
		wantErr assert.ErrorAssertionFunc
	}

	query := `
UPDATE
	goals
SET
	description = 'New goal',
	updated_at = '2023-07-20T15:32:00Z'
WHERE
	id = 1;`
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
					goals: habit_tracker.Goals{
						{
							ID:          1,
							Description: "New goal",
						},
					},
					now: time.Date(2023, 7, 20, 15, 32, 0, 0, time.UTC),
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
					goals: habit_tracker.Goals{
						{
							ID:          1,
							Description: "New goal",
						},
					},
					now: time.Date(2023, 7, 20, 15, 32, 0, 0, time.UTC),
				},
				wantErr: assert.Error,
			}
		}(),
	}
	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				gr := &GoalRepository{
					db: tt.fields.db,
				}
				tt.wantErr(t, gr.UpdateGoals(tt.args.ctx, tt.args.goals, tt.args.now))
			},
		)
	}
}

func Test_buildInsertGoalsQuery(t *testing.T) {
	type args struct {
		goals habit_tracker.Goals
		now   time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Success",
			args: args{
				goals: habit_tracker.Goals{
					{
						Description: "New goal",
					},
				},
				now: time.Date(2023, 7, 20, 15, 32, 0, 0, time.UTC),
			},
			want: `
INSERT
	INTO
	goals
(description, created_at, updated_at)
VALUES ('New goal', '2023-07-20T15:32:00Z', '2023-07-20T15:32:00Z');`,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				assert.Equal(t, tt.want, buildInsertGoalsQuery(tt.args.goals, tt.args.now))
			},
		)
	}
}

func Test_buildUpdateGoalsQuery(t *testing.T) {
	type args struct {
		goals habit_tracker.Goals
		now   time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Success",
			args: args{
				goals: habit_tracker.Goals{
					{
						ID:          1,
						Description: "New goal",
					},
				},
				now: time.Date(2023, 7, 20, 15, 32, 0, 0, time.UTC),
			},
			want: `
UPDATE
	goals
SET
	description = 'New goal',
	updated_at = '2023-07-20T15:32:00Z'
WHERE
	id = 1;`,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				assert.Equal(t, tt.want, buildUpdateGoalsQuery(tt.args.goals, tt.args.now))
			},
		)
	}
}
