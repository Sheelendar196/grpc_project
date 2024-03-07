package service

import (
	"context"

	"github.com/sheelendar196/go-projects/graphql-project/internal/core/domain"
)

type EmployeeInteractor struct {
	Repo EmployeeRepo
}

// NewEmployeeInteractor return valid employee interactor.
func NewEmployeeInteractor(repo EmployeeRepo) *EmployeeInteractor {
	if repo == nil {
		return nil
	}
	return &EmployeeInteractor{Repo: repo}
}

func (es EmployeeInteractor) SaveEmployee(ctx context.Context, employee *domain.Employee) error {
	return es.Repo.SaveEmployee(ctx, employee)
}
func (es EmployeeInteractor) GetEmployeeByID(ctx context.Context, empID string) (*domain.Employee, error) {
	return es.Repo.GetEmployeeByID(ctx, empID)
}
func (es EmployeeInteractor) GetEmployeeList(ctx context.Context) ([]*domain.Employee, error) {
	return es.Repo.GetEmployeeList(ctx)
}

func (es EmployeeInteractor) UpdateEmployeeDetails(ctx context.Context, input domain.Employee) (*domain.Employee, error) {
	return es.Repo.UpdateEmployeeDetails(ctx, input)
}

func (es EmployeeInteractor) DeleteEmployee(ctx context.Context, empID string) (*domain.Employee, error) {
	return es.Repo.DeleteEmployee(ctx, empID)
}
