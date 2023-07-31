// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"
	habit_tracker "habit-tracker"

	mock "github.com/stretchr/testify/mock"
)

// EventRepository is an autogenerated mock type for the EventRepository type
type EventRepository struct {
	mock.Mock
}

// InsertEvents provides a mock function with given fields: ctx, events
func (_m *EventRepository) InsertEvents(ctx context.Context, events habit_tracker.Event) error {
	ret := _m.Called(ctx, events)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, habit_tracker.Event) error); ok {
		r0 = rf(ctx, events)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewEventRepository creates a new instance of EventRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEventRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *EventRepository {
	mock := &EventRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
