package domain

import "time"

type EmployeeRole string

const (
	EMPLOEE EmployeeRole = "EMPLOYEE"
	MANAGER EmployeeRole = "MANAGER"
)

type Employee struct {
	ID         string  `validate:""`
	Name       string  `validate:"required"`
	EmpID      string  `validate:"required"`
	Mobile     *string `validate:""`
	Email      *string `validate:"required"`
	Address    *string `validate:""`
	ManagerID  *string `validate:""`
	IsActive   *bool   `validate:"required"`
	Department *string `validate:"required"`

	CreatedAt     time.Time  `validate:"required"`
	CreatedBy     string     `validate:"required,ascii,max=100"`
	LastUpdatedAt *time.Time `validate:""`
	LastUpdatedBy *string    `validate:"omitempty,ascii,max=100"`
}
