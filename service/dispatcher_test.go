package service_test

import (
	"context"
	"testing"

	"data-enricher-dispatcher/config"
	"data-enricher-dispatcher/model"
	"data-enricher-dispatcher/service"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAPIClient struct {
	mock.Mock
}

// func (m *MockAPIClient) GetUsers() ([]model.User, error) {
// 	args := m.Called()
// 	return args.Get(0).([]model.User), args.Error(1)
// }

func (m *MockAPIClient) GetUsers(ctx context.Context) ([]model.User, error) {
	args := m.Called(ctx)
	if users, ok := args.Get(0).([]model.User); ok {
		return users, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockAPIClient) PostUser(ctx context.Context, user model.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

type MockLogger struct {
	mock.Mock
}

func (m *MockLogger) Println(v ...interface{}) {
	m.Called(v...)
}

func (m *MockLogger) Error(v ...interface{}) {
	m.Called(v...)
}

func (m *MockLogger) Debug(v ...interface{}) {
	m.Called(v...)
}

func (m *MockLogger) Fatal(v ...interface{}) {
	m.Called(v...)
}

func (m *MockLogger) Info(v ...interface{}) {
	m.Called(v...)
}

func (m *MockLogger) Warn(v ...interface{}) {
	m.Called(v...)
}

func TestDispatcher_Start(t *testing.T) {
	mockClient := new(MockAPIClient)
	mockLogger := new(MockLogger)
	cfg := &config.Config{ExcludePostfixes: []string{"@test.com"}}

	users := []model.User{
		{Name: "John Doe", Email: "john@test.com"},
		{Name: "", Email: "jane@test.com"},
		{Name: "Other User", Email: "other@other.com"},
	}

	mockClient.On("GetUsers", mock.MatchedBy(func(ctx context.Context) bool {
		_, ok := ctx.Deadline()
		return ok
	})).Return(users, nil)
	mockClient.On("PostUser", mock.MatchedBy(func(ctx context.Context) bool {
		return ctx != nil
	}), users[0]).Return(nil)
	mockLogger.On("Println", mock.Anything).Once()
	mockLogger.On("Info", mock.Anything).Once()

	ctx := context.Background()
	d := service.NewDispatcher(mockClient, mockLogger, cfg)
	err := d.Start(ctx)
	assert.NoError(t, err)

	mockClient.AssertExpectations(t)
	mockLogger.AssertExpectations(t)
}
