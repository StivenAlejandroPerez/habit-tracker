// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"
	habit_tracker "habit-tracker"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// TagRepository is an autogenerated mock type for the TagRepository type
type TagRepository struct {
	mock.Mock
}

// InsertTags provides a mock function with given fields: ctx, tags, now
func (_m *TagRepository) InsertTags(ctx context.Context, tags habit_tracker.Tags, now time.Time) error {
	ret := _m.Called(ctx, tags, now)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, habit_tracker.Tags, time.Time) error); ok {
		r0 = rf(ctx, tags, now)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateTags provides a mock function with given fields: ctx, tags, now
func (_m *TagRepository) UpdateTags(ctx context.Context, tags habit_tracker.Tags, now time.Time) error {
	ret := _m.Called(ctx, tags, now)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, habit_tracker.Tags, time.Time) error); ok {
		r0 = rf(ctx, tags, now)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewTagRepository creates a new instance of TagRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTagRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *TagRepository {
	mock := &TagRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}