package employees

import (
	employeeRepo "github.com/almarino_meli/grupo-5-wave-15/internal/repository/employees"
	warehouseRepo "github.com/almarino_meli/grupo-5-wave-15/internal/repository/warehouses"
	"github.com/almarino_meli/grupo-5-wave-15/pkg/errors"
	"github.com/almarino_meli/grupo-5-wave-15/pkg/models/employees"
)

type EmployeeDefault struct {
	repository          employeeRepo.EmployeeRepository
	warehouseRepository warehouseRepo.WarehouseRepository
}

func NewEmployeeService(repository employeeRepo.EmployeeRepository, warehouseRepository warehouseRepo.WarehouseRepository) EmployeeDefault {
	return EmployeeDefault{repository: repository, warehouseRepository: warehouseRepository}
}

func (es EmployeeDefault) Create(e *employees.Employee) (createdEmployee employees.Employee, err error) {
	if e.WarehouseID != 0 {
		_, err = es.warehouseRepository.GetByID(e.WarehouseID)
		if err != nil {
			return employees.Employee{}, &errors.NotFoundError{Message: "Warehouse not found"}
		}
	}

	createdEmployee, err = es.repository.Create(e)
	return
}

func (es EmployeeDefault) GetAll() (e map[int]employees.Employee, err error) {
	e, err = es.repository.GetAll()

	if len(e) == 0 {
		return e, &errors.NotFoundError{Message: "No employees found"}
	}

	return
}

func (es EmployeeDefault) GetByID(id int) (e employees.Employee, err error) {
	e, err = es.repository.GetByID(id)
	return
}

func (es EmployeeDefault) Update(e employees.EmployeePatch) (updatedEmployee employees.Employee, err error) {
	existingEmployee, err := es.repository.GetByID(*e.Id)
	if err != nil {
		return employees.Employee{}, err
	}

	es.ApplyPatch(&existingEmployee, e)

	if *e.WarehouseID != 0 {
		_, err = es.warehouseRepository.GetByID(*e.WarehouseID)
		if err != nil {
			return employees.Employee{}, &errors.NotFoundError{Message: "Warehouse not found"}
		}
	}

	updatedEmployee, err = es.repository.Update(&existingEmployee)
	if err != nil {
		return employees.Employee{}, err
	}
	return
}

func (es EmployeeDefault) Delete(id int) (err error) {
	err = es.repository.Delete(id)
	return
}
