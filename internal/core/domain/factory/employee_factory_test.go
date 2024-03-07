package factory

import (
	"context"
	"errors"
	"testing"

	"github.com/sheelendar196/go-projects/graphql-project/internal/core/domain"
	"github.com/sheelendar196/go-projects/graphql-project/internal/core/domain/factory/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	empServiceTest := mocks.NewEmployeeService(t)
	testEF := NewEmployeeFactory(empServiceTest)
	tests := []struct {
		name        string
		ctx         context.Context
		emp         *domain.Employee
		expectedErr error
	}{
		{
			name:        "create_error_case",
			ctx:         context.Background(),
			emp:         &domain.Employee{},
			expectedErr: errors.New("not insert into db"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			empServiceTest.On("SaveEmployee", mock.Anything, tt.emp).Return(errors.New("not insert into db")).Once()
			err := testEF.Create(tt.ctx, tt.emp)
			assert.Equal(t, err, tt.expectedErr)

		})
	}
}
