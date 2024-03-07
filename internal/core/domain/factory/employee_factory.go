package factory

import (
	"context"

	"github.com/sheelendar196/go-projects/graphql-project/internal/core/domain"
)

type EmployeeFactory struct {
	employeeInteractor EmployeeService
}

func NewEmployeeFactory(employeeService EmployeeService) *EmployeeFactory {
	return &EmployeeFactory{employeeInteractor: employeeService}
}

func (ef *EmployeeFactory) Create(ctx context.Context, emp *domain.Employee) error {
	// check employee details and verify all fields fo employee before create
	return ef.employeeInteractor.SaveEmployee(ctx, emp)
}
