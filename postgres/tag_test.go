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

func TestTagRepository_InsertTags(t *testing.T) {
	type fields struct {
		db Drivers
	}
	type args struct {
		ctx  context.Context
		tags habit_tracker.Tags
		now  time.Time
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
	tags
(name, description, created_at, updated_at)
VALUES ('Health', 'New description', '2023-07-30T12:00:00Z', '2023-07-30T12:00:00Z');`

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
					tags: habit_tracker.Tags{
						{
							Name:        "Health",
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
					tags: habit_tracker.Tags{
						{
							Name:        "Health",
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
				tr := NewTagRepository(tt.fields.db)

				tt.wantErr(t, tr.InsertTags(tt.args.ctx, tt.args.tags, tt.args.now))
			},
		)
	}
}

func TestTagRepository_UpdateTags(t *testing.T) {
	type fields struct {
		db Drivers
	}
	type args struct {
		ctx  context.Context
		tags habit_tracker.Tags
		now  time.Time
	}
	type test struct {
		name    string
		fields  fields
		args    args
		wantErr assert.ErrorAssertionFunc
	}

	query := `
UPDATE
	tags
SET
	"name" = 'Health',
	description = 'New description',
	updated_at = '2023-07-30T12:00:00Z'
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
					tags: habit_tracker.Tags{
						{
							ID:          1,
							Name:        "Health",
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
				name: "Success",
				fields: fields{
					db: &Postgres{
						db: db,
					},
				},
				args: args{
					ctx: context.Background(),
					tags: habit_tracker.Tags{
						{
							ID:          1,
							Name:        "Health",
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
				tr := &TagRepository{
					db: tt.fields.db,
				}
				tt.wantErr(t, tr.UpdateTags(tt.args.ctx, tt.args.tags, tt.args.now))
			},
		)
	}
}

func Test_buildInsertTagsQuery(t *testing.T) {
	type args struct {
		tags habit_tracker.Tags
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
				tags: habit_tracker.Tags{
					{
						ID:          1,
						Name:        "Health",
						Description: "New description",
					},
				},
				now: time.Date(2023, 7, 30, 12, 0, 0, 0, time.UTC),
			},
			want: `
INSERT
	INTO
	tags
(name, description, created_at, updated_at)
VALUES ('Health', 'New description', '2023-07-30T12:00:00Z', '2023-07-30T12:00:00Z');`,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				assert.Equal(t, tt.want, buildInsertTagsQuery(tt.args.tags, tt.args.now))
			},
		)
	}
}

func Test_buildUpdateTagsQuery(t *testing.T) {
	type args struct {
		tags habit_tracker.Tags
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
				tags: habit_tracker.Tags{
					{
						ID:          1,
						Name:        "Health",
						Description: "New description",
					},
				},
				now: time.Date(2023, 7, 30, 12, 0, 0, 0, time.UTC),
			},
			want: `
UPDATE
	tags
SET
	"name" = 'Health',
	description = 'New description',
	updated_at = '2023-07-30T12:00:00Z'
WHERE
	id = 1;`,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				assert.Equal(t, tt.want, buildUpdateTagsQuery(tt.args.tags, tt.args.now))
			},
		)
	}
}
