package employees

import (
	"github.com/almarino_meli/grupo-5-wave-15/pkg/models/employees"
)

func (es EmployeeDefault) ApplyPatch(e *employees.Employee, patch employees.EmployeePatch) {
	if patch.CardNumberID != nil && *patch.CardNumberID != "" {
		e.CardNumberID = *patch.CardNumberID
	}
	if patch.FirstName != nil && *patch.FirstName != "" {
		e.FirstName = *patch.FirstName
	}
	if patch.LastName != nil && *patch.LastName != "" {
		e.LastName = *patch.LastName
	}
	if patch.WarehouseID != nil && *patch.WarehouseID != 0 {
		e.WarehouseID = *patch.WarehouseID
	}

	return
}
