package service_test

import (
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

func (m *MockAPIClient) GetUsers() ([]model.User, error) {
	args := m.Called()
	return args.Get(0).([]model.User), args.Error(1)
}

func (m *MockAPIClient) PostUser(user model.User) error {
	args := m.Called(user)
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

	mockClient.On("GetUsers").Return(users, nil)
	mockClient.On("PostUser", users[0]).Return(nil)
	mockLogger.On("Println", mock.Anything).Twice()
	mockLogger.On("Info", mock.Anything).Once()

	d := service.NewDispatcher(mockClient, mockLogger, cfg)
	err := d.Start()
	assert.NoError(t, err)

	mockClient.AssertExpectations(t)
	mockLogger.AssertExpectations(t)
}
