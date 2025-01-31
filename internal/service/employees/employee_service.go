package employees

import (
	"github.com/almarino_meli/grupo-5-wave-15/pkg/models/employees"
)

type EmployeeService interface {
	Create(e *employees.Employee) (createdEmployee employees.Employee, err error)
	GetAll() (e map[int]employees.Employee, err error)
	GetByID(id int) (e employees.Employee, err error)
	Update(e employees.EmployeePatch) (updatedEmployee employees.Employee, err error)
	Delete(id int) (err error)
}
