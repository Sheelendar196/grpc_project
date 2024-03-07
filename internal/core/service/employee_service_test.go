package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRoleInteractor(t *testing.T) {
	tests := []struct {
		name string
		repo EmployeeRepo
	}{
		{
			name: "repo nil test",
			repo: nil,
		},
		{
			name: "repo not nil test",
			repo: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			es := NewEmployeeInteractor(tt.repo)
			if tt.repo == nil {
				assert.Nil(t, es)
			} else {
				assert.NotNil(t, es)
			}

		})
	}
}
