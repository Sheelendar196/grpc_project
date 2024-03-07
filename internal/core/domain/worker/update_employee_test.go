package worker

import (
	"context"
	"errors"
	"testing"

	"github.com/sheelendar196/go-projects/graphql-project/internal/core/domain"
	"github.com/sheelendar196/go-projects/graphql-project/internal/core/domain/worker/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	empServiceTest := mocks.NewEmployeeService(t)
	testUE := NewUpdateEmployee(empServiceTest)
	tests := []struct {
		name        string
		ctx         context.Context
		emp         domain.Employee
		expectedRes *domain.Employee
		expectedErr error
		mockF       func(empServiceTest *mocks.EmployeeService, emp domain.Employee)
	}{
		{
			name:        "update_error",
			ctx:         context.Background(),
			emp:         domain.Employee{},
			expectedErr: errors.New("not able to update into db"),
			mockF: func(empServiceTest *mocks.EmployeeService, emp domain.Employee) {
				empServiceTest.On("UpdateEmployeeDetails", mock.Anything, emp).Return(nil, errors.New("not able to update into db")).Once()
			},
		},
		{
			name:        "update_success",
			ctx:         context.Background(),
			emp:         domain.Employee{},
			expectedRes: &domain.Employee{Name: "sheelendar", EmpID: "322422"},
			expectedErr: nil,
			mockF: func(empServiceTest *mocks.EmployeeService, emp domain.Employee) {
				empServiceTest.On("UpdateEmployeeDetails", mock.Anything, emp).Return(&domain.Employee{Name: "sheelendar", EmpID: "322422"}, nil).Once()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if tt.mockF != nil {
				tt.mockF(empServiceTest, tt.emp)
			}
			res, err := testUE.Update(tt.ctx, tt.emp)
			assert.Equal(t, res, tt.expectedRes)
			assert.Equal(t, err, tt.expectedErr)

		})
	}
}
