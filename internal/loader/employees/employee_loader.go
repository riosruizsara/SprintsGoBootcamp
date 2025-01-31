package employees

import "github.com/almarino_meli/grupo-5-wave-15/pkg/models/employees"

type EmployeeLoader interface {
	Load() (v map[int]employees.Employee, err error)
}
