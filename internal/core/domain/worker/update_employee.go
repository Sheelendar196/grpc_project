package worker

import (
	"context"

	"github.com/sheelendar196/go-projects/graphql-project/internal/core/domain"
)

type UpdateEmployee struct {
	employeeInteractor EmployeeService
}

func NewUpdateEmployee(employeeService EmployeeService) *UpdateEmployee {
	return &UpdateEmployee{employeeInteractor: employeeService}
}

func (ef *UpdateEmployee) Update(ctx context.Context, emp domain.Employee) (*domain.Employee, error) {
	// check employee details and verify all fields fo employee before create
	return ef.employeeInteractor.UpdateEmployeeDetails(ctx, emp)
}

func (ef *UpdateEmployee) Detete(ctx context.Context, empID string) (*domain.Employee, error) {
	// check employee details and verify all fields fo employee before create
	return ef.employeeInteractor.DeleteEmployee(ctx, empID)
}
