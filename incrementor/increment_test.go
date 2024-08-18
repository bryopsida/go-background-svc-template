package incrementor

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/bryopsida/go-background-svc-template/interfaces"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockNumberRepository is a mock implementation of the INumberRepository interface.
type incMockNumberRepository struct {
	mock.Mock
}

func (m *incMockNumberRepository) Save(number interfaces.Number) error {
	args := m.Called(number)
	return args.Error(0)
}

func (m *incMockNumberRepository) FindByID(calledID string) (*interfaces.Number, error) {
	args := m.Called(calledID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*interfaces.Number), args.Error(1)
}

func (m *incMockNumberRepository) DeleteByID(calledID string) error {
	args := m.Called(calledID)
	return args.Error(0)
}

func TestIncrement(t *testing.T) {
	mockRepo := new(incMockNumberRepository)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Test case: Record not found, initialize record
	mockRepo.On("FindByID", getID()).Return(nil, errors.New("key not found"))
	mockRepo.On("Save", mock.Anything).Return(nil)

	go Increment(ctx, mockRepo)
	time.Sleep(2 * time.Second)
	cancel()

	mockRepo.AssertCalled(t, "FindByID", getID())
	mockRepo.AssertCalled(t, "Save", mock.Anything)

	// Test case: Record found, increment number
	number := &interfaces.Number{ID: getID(), Number: 0}
	mockRepo.On("FindByID", getID()).Return(number, nil)
	mockRepo.On("Save", *number).Return(nil)

	ctx, cancel = context.WithCancel(context.Background())
	defer cancel()

	go Increment(ctx, mockRepo)
	time.Sleep(2 * time.Second)
	cancel()

	mockRepo.AssertCalled(t, "FindByID", getID())
	mockRepo.AssertCalled(t, "Save", *number)
	assert.Equal(t, uint64(0), number.Number)
}
