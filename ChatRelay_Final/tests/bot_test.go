// Unit test for bot logic
package bot_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"chatrelay/bot"
)

type MockBackendClient struct {
	mock.Mock
}

func (m *MockBackendClient) StreamChat(ctx context.Context, userID, query string) (<-chan string, error) {
	args := m.Called(ctx, userID, query)
	return args.Get(0).(<-chan string), args.Error(1)
}

func TestHandleMessage_Success(t *testing.T) {
	ctx := context.Background()
	mockClient := new(MockBackendClient)
	textChunks := make(chan string, 2)
	textChunks <- "Hello, "
	textChunks <- "World!"
	close(textChunks)

	mockClient.On("StreamChat", ctx, "U123", "Hi").Return((<-chan string)(textChunks), nil)

	handler := bot.NewHandler(mockClient)
	err := handler.HandleMessage(ctx, "U123", "Hi")

	assert.NoError(t, err)
	mockClient.AssertExpectations(t)
}

func TestHandleMessage_Error(t *testing.T) {
	ctx := context.Background()
	mockClient := new(MockBackendClient)
	expectedErr := errors.New("backend error")
	mockClient.On("StreamChat", ctx, "U123", "Hi").Return(nil, expectedErr)

	handler := bot.NewHandler(mockClient)
	err := handler.HandleMessage(ctx, "U123", "Hi")

	assert.EqualError(t, err, "backend error")
	mockClient.AssertExpectations(t)
}
