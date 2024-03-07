package repository

import (
	"context"
	"errors"

	"github.com/sheelendar196/go-projects/grpc_project/internal/core/domain"
	"github.com/sheelendar196/go-projects/grpc_project/internal/infrastructure/models/tables"
	"gorm.io/gorm"
)

type EmployeeRepository struct {
	gormDB *gorm.DB
	list   []*tables.Employee
}

func NewEmployeeRepository(g *gorm.DB) *EmployeeRepository {
	return &EmployeeRepository{
		gormDB: g,
	}
}

func (er *EmployeeRepository) GetEmployeeEntryByID(ctx context.Context, empID string) (*domain.Employee, error) {
	for _, emp := range er.list {
		if emp.EmpID == empID {
			return emp.ToDomain(), nil
		}
	}
	return nil, errors.New("employee not found in db")
}
func (er *EmployeeRepository) InsertEmplyee(ctx context.Context, emp *domain.Employee) error {
	er.list = append(er.list, tables.ToTable(emp))
	return nil
}
