package incrementor

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/bryopsida/go-background-svc-template/interfaces"
	"github.com/stretchr/testify/mock"
)

// MockNumberRepository is a mock implementation of the INumberRepository interface.
type printMockNumberRepository struct {
	mock.Mock
}

func (m *printMockNumberRepository) Save(number interfaces.Number) error {
	args := m.Called(number)
	return args.Error(0)
}

func (m *printMockNumberRepository) FindByID(findID string) (*interfaces.Number, error) {
	args := m.Called(findID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*interfaces.Number), args.Error(1)
}

func (m *printMockNumberRepository) DeleteByID(findID string) error {
	args := m.Called(findID)
	return args.Error(0)
}

func TestPrint(t *testing.T) {
	mockRepo := new(printMockNumberRepository)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Test case: Record not found
	mockRepo.On("FindByID", getID()).Return(nil, errors.New("key not found"))

	go Print(ctx, mockRepo)
	time.Sleep(2 * time.Second)
	cancel()

	mockRepo.AssertCalled(t, "FindByID", getID())

	// Test case: Record found
	number := &interfaces.Number{ID: getID(), Number: 42}
	mockRepo.On("FindByID", getID()).Return(number, nil)

	ctx, cancel = context.WithCancel(context.Background())
	defer cancel()

	go Print(ctx, mockRepo)
	time.Sleep(2 * time.Second)
	cancel()

	mockRepo.AssertCalled(t, "FindByID", getID())
}
