package grpc

import (
	"context"

	"github.com/sheelendar196/go-projects/grpc_project/internal/core/domain"
)

type EmployeeService interface {
	GetEmployeeEntryByID(ctx context.Context, empID string) (*domain.Employee, error)
	InsertEmplyee(ctx context.Context, emp *domain.Employee) error
}
