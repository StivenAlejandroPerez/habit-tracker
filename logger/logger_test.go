package logger

import (
	"context"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestGetLogger(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want func(t *testing.T, actual interface{}) bool
	}{
		{
			name: "WithPrefix",
			args: args{
				ctx: context.WithValue(context.Background(), "Logger", New("test")),
			},
			want: func(t *testing.T, actual interface{}) bool {
				return assert.Equal(t, "test", GetPrefix(actual.(*logrus.Entry)))
			},
		},
		{
			name: "WithoutPrefix",
			args: args{
				ctx: context.WithValue(context.Background(), "Logger", New("-")),
			},
			want: func(t *testing.T, actual interface{}) bool {
				assert.NotEmpty(t, GetPrefix(actual.(*logrus.Entry)))
				return assert.NotEqual(t, "-", GetPrefix(actual.(*logrus.Entry)))
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				tt.want(t, GetLogger(tt.args.ctx))
			},
		)
	}
}

func TestGetPrefix(t *testing.T) {
	type args struct {
		logger *logrus.Entry
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Success",
			args: args{
				logger: New("test"),
			},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				assert.Equal(t, tt.want, GetPrefix(tt.args.logger))
			},
		)
	}
}

func TestNew(t *testing.T) {
	type args struct {
		prefix string
	}

	tests := []struct {
		name string
		args args
		want func(t *testing.T, actual interface{}) bool
	}{
		{
			name: "Success",
			args: args{
				prefix: "test",
			},
			want: func(t *testing.T, actual interface{}) bool {
				return assert.Equal(t, "test", actual)
			},
		},
		{
			name: "Default",
			args: args{
				prefix: "-",
			},
			want: func(t *testing.T, actual interface{}) bool {
				assert.NotEmpty(t, actual)
				return assert.NotEqual(t, "-", actual)
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				log := New(tt.args.prefix)
				tt.want(t, GetPrefix(log))
			},
		)
	}
}

func TestNewContext(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{
			name: "Success",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				_, ok := NewContext().Value(loggerKey).(*logrus.Entry)
				assert.Equal(t, tt.want, ok)
			},
		)
	}
}

func TestSetLogger(t *testing.T) {
	type args struct {
		ctx    context.Context
		logger *logrus.Entry
	}

	log := New("test")

	tests := []struct {
		name string
		args args
		want context.Context
	}{
		{
			name: "Success",
			args: args{
				ctx:    context.Background(),
				logger: log,
			},
			want: context.WithValue(context.Background(), "Logger", log),
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				assert.Equal(t, tt.want, SetLogger(tt.args.ctx, tt.args.logger))
			},
		)
	}
}
