package tables

import (
	"time"

	"github.com/sheelendar196/go-projects/grpc_project/internal/core/domain"
)

func (Employee) TableName() string {
	return "employee"
}

type Employee struct {
	ID         string  `gorm:"column:id;primaryKey"`
	Name       string  `gorm:"column:name"`
	EmpID      string  `gorm:"column:emp_id"`
	Mobile     *string `gorm:"column:mobile"`
	Email      *string `gorm:"column:email"`
	Address    *string `gorm:"column:address"`
	ManagerID  *string `gorm:"column:manager_id"`
	IsActive   *bool   `gorm:"column:is_active"`
	Department *string `gorm:"column:department"`

	CreatedAt     time.Time  `gorm:"column:created_at"`
	CreatedBy     string     `gorm:"column:created_by"`
	LastUpdatedAt *time.Time `gorm:"column:lastupdated_at"`
	LastUpdatedBy *string    `gorm:"column:lastupdated_by"`
}

func (e Employee) ToDomain() *domain.Employee {
	return &domain.Employee{
		Name:       e.Name,
		EmpID:      e.EmpID,
		Mobile:     e.Mobile,
		Email:      e.Email,
		Address:    e.Address,
		ManagerID:  e.ManagerID,
		Department: e.Department,
		IsActive:   e.IsActive,
	}
}

func ToTable(e *domain.Employee) *Employee {
	return &Employee{
		Name:       e.Name,
		EmpID:      e.EmpID,
		Mobile:     e.Mobile,
		Email:      e.Email,
		Address:    e.Address,
		ManagerID:  e.ManagerID,
		Department: e.Department,
		IsActive:   e.IsActive,
	}
}

// UpdateEmployee update employee into table once actual table we can remove this func.
func UpdateEmployee(emp *Employee, input domain.Employee) {
	if input.Name != "" {
		emp.Name = input.Name
	}
	if input.Address != nil {
		emp.Address = input.Address
	}
	if input.Department != nil {
		emp.Department = input.Department
	}
	if input.Email != nil {
		*emp.Email = input.EmpID
	}
	if input.Mobile != nil {
		emp.Mobile = input.Mobile
	}
	if input.IsActive != nil {
		emp.IsActive = input.IsActive
	}
	if input.ManagerID != nil {
		emp.ManagerID = input.ManagerID
	}
}
