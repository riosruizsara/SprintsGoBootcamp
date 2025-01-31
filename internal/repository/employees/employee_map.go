package employees

import (
	"github.com/almarino_meli/grupo-5-wave-15/pkg/errors"
	"github.com/almarino_meli/grupo-5-wave-15/pkg/models/employees"
)

type EmployeeMap struct {
	db map[int]employees.Employee
}

func NewEmployeeRepository(db map[int]employees.Employee) *EmployeeMap {
	defaultDb := make(map[int]employees.Employee)
	if db != nil {
		defaultDb = db
	}

	return &EmployeeMap{db: defaultDb}
}

func (er *EmployeeMap) Create(employee *employees.Employee) (createdEmployee employees.Employee, err error) {
	if er.employeeExistsByCardID(employee.CardNumberID) {
		return createdEmployee, &errors.DuplicateError{Message: "Error: Employee already exists with provided card ID"}
	}

	id := er.getNextID()
	employee.Id = &id
	er.db[*employee.Id] = *employee
	createdEmployee = *employee

	return
}

func (er *EmployeeMap) GetAll() (e map[int]employees.Employee, err error) {
	e = make(map[int]employees.Employee)
	for id, employee := range er.db {
		e[id] = employee
	}
	return
}

func (er *EmployeeMap) GetByID(id int) (e employees.Employee, err error) {
	e, ok := er.db[id]
	if !ok {
		return employees.Employee{}, &errors.NotFoundError{Message: "Employee not found"}
	}
	return
}

func (er *EmployeeMap) Update(e *employees.Employee) (updatedEmployee employees.Employee, err error) {
	_, ok := er.db[*e.Id]
	if !ok {
		return updatedEmployee, &errors.NotFoundError{Message: "Employee not found"}
	}

	for _, emp := range er.db {
		if emp.CardNumberID == e.CardNumberID && emp.Id != e.Id {
			return updatedEmployee, &errors.DuplicateError{Message: "Error: Employee already exists with provided card ID"}
		}
	}

	er.db[*e.Id] = *e
	updatedEmployee = *e

	return
}

func (er *EmployeeMap) Delete(id int) (err error) {
	_, ok := er.db[id]
	if !ok {
		return &errors.NotFoundError{Message: "Employee not found"}
	}

	delete(er.db, id)
	return
}

func (er *EmployeeMap) employeeExistsByCardID(cardID string) bool {
	for _, emp := range er.db {
		if emp.CardNumberID == cardID {
			return true
		}
	}
	return false
}

func (er *EmployeeMap) getNextID() int {
	newId := 0

	for _, employee := range er.db {
		if *employee.Id > newId {
			newId = *employee.Id
		}
	}

	return newId + 1
}
