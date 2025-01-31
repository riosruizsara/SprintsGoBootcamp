package employees

import (
	"encoding/json"
	"os"

	"github.com/almarino_meli/grupo-5-wave-15/pkg/models/employees"
)

type EmployeeJSONFile struct {
	path string
}

func NewEmployeeJSON(path string) *EmployeeJSONFile {
	return &EmployeeJSONFile{
		path: path,
	}
}

func (l *EmployeeJSONFile) Load() (e map[int]employees.Employee, err error) {
	file, err := os.Open(l.path)
	if err != nil {
		return
	}
	defer file.Close()

	var employeesJSON []employees.EmployeeDTO
	err = json.NewDecoder(file).Decode(&employeesJSON)
	if err != nil {
		return
	}

	e = make(map[int]employees.Employee)
	for _, employeeDTO := range employeesJSON {
		employee := employees.ToEmployeeModel(employeeDTO)
		e[*employee.Id] = employee
	}

	return

}
