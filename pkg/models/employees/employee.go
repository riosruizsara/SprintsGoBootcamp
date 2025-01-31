package employees

import "github.com/go-playground/validator/v10"

type Employee struct {
	Id           *int   `json:"id"`
	CardNumberID string `json:"card_number_id" validate:"required,numeric"`
	FirstName    string `json:"first_name" validate:"required,min=2,max=50"`
	LastName     string `json:"last_name" validate:"required,min=2,max=50"`
	WarehouseID  int    `json:"warehouse_id" validate:"numeric"`
}

type EmployeePatch struct {
	Id           *int    `json:"id"`
	CardNumberID *string `json:"card_number_id" validate:"omitempty,numeric"`
	FirstName    *string `json:"first_name" validate:"omitempty,min=2,max=50"`
	LastName     *string `json:"last_name" validate:"omitempty,min=2,max=50"`
	WarehouseID  *int    `json:"warehouse_id" validate:"omitempty,numeric"`
}

func ToEmployeeDTO(e Employee) EmployeeDTO {
	return EmployeeDTO{
		Id:           e.Id,
		CardNumberID: e.CardNumberID,
		FirstName:    e.FirstName,
		LastName:     e.LastName,
		WarehouseID:  e.WarehouseID,
	}
}

func (e *Employee) Validate(validate *validator.Validate) error {
	return validate.Struct(e)
}

func (e *EmployeePatch) Validate(validate *validator.Validate) error {
	return validate.Struct(e)
}
