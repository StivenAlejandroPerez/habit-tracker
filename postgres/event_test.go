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

func TestEventRepository_InsertEvents(t *testing.T) {
	type fields struct {
		db Drivers
	}
	type args struct {
		ctx    context.Context
		events habit_tracker.Events
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
	events
(habit_id, subject, start_at, end_at, created_at, updated_at)
VALUES (2, 'Go to gym', '2023-07-27T12:00:00Z', '2023-07-27T14:00:00Z', '2023-07-20T15:32:00Z', '2023-07-20T15:32:00Z'), (3, 'Painting class', '2023-07-27T14:00:00Z', '2023-07-27T16:00:00Z', '2023-07-20T15:32:00Z', '2023-07-20T15:32:00Z');`
	tests := []test{
		func() test {
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)

			mock.ExpectExec(regexp.QuoteMeta(query)).WillReturnResult(sqlmock.NewResult(3, 2))

			return test{
				name: "Success",
				fields: fields{
					db: &Postgres{
						db: db,
					},
				},
				args: args{
					ctx: context.Background(),
					events: habit_tracker.Events{
						{
							HabitID: 2,
							Subject: "Go to gym",
							StartAt: time.Date(2023, 7, 27, 12, 0, 0, 0, time.UTC),
							EndAt:   time.Date(2023, 7, 27, 14, 0, 0, 0, time.UTC),
						},
						{
							HabitID: 3,
							Subject: "Painting class",
							StartAt: time.Date(2023, 7, 27, 14, 0, 0, 0, time.UTC),
							EndAt:   time.Date(2023, 7, 27, 16, 0, 0, 0, time.UTC),
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
					events: habit_tracker.Events{
						{
							HabitID: 2,
							Subject: "Go to gym",
							StartAt: time.Date(2023, 7, 27, 12, 0, 0, 0, time.UTC),
							EndAt:   time.Date(2023, 7, 27, 14, 0, 0, 0, time.UTC),
						},
						{
							HabitID: 3,
							Subject: "Painting class",
							StartAt: time.Date(2023, 7, 27, 14, 0, 0, 0, time.UTC),
							EndAt:   time.Date(2023, 7, 27, 16, 0, 0, 0, time.UTC),
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
				er := NewEventRepository(tt.fields.db)

				tt.wantErr(t, er.InsertEvents(tt.args.ctx, tt.args.events, tt.args.now))
			},
		)
	}
}

func Test_buildInsertEventsQuery(t *testing.T) {
	type args struct {
		events habit_tracker.Events
		now    time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Success",
			args: args{
				events: habit_tracker.Events{
					{
						HabitID: 2,
						Subject: "Go to gym",
						StartAt: time.Date(2023, 7, 27, 12, 0, 0, 0, time.UTC),
						EndAt:   time.Date(2023, 7, 27, 14, 0, 0, 0, time.UTC),
					},
					{
						HabitID: 3,
						Subject: "Painting class",
						StartAt: time.Date(2023, 7, 27, 14, 0, 0, 0, time.UTC),
						EndAt:   time.Date(2023, 7, 27, 16, 0, 0, 0, time.UTC),
					},
				},
				now: time.Date(2023, 7, 20, 15, 32, 0, 0, time.UTC),
			},
			want: `
INSERT
	INTO
	events
(habit_id, subject, start_at, end_at, created_at, updated_at)
VALUES (2, 'Go to gym', '2023-07-27T12:00:00Z', '2023-07-27T14:00:00Z', '2023-07-20T15:32:00Z', '2023-07-20T15:32:00Z'), (3, 'Painting class', '2023-07-27T14:00:00Z', '2023-07-27T16:00:00Z', '2023-07-20T15:32:00Z', '2023-07-20T15:32:00Z');`,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				assert.Equal(t, tt.want, buildInsertEventsQuery(tt.args.events, tt.args.now))
			},
		)
	}
}
