package sellers

import (
	"github.com/go-playground/validator/v10"
)

type Seller struct {
	Id          int    `validate:"numeric"`
	CompanyId   int    `validate:"numeric"`
	CompanyName string `validate:"required,min=3"`
	Address     string `validate:"required,min=3"`
	Telephone   string `validate:"numeric"`
}

type SellerPatch struct {
	CompanyId   *int    `json:"cid" validate:"omitempty,numeric,min=1"`
	CompanyName *string `json:"company_name" validate:"omitempty,min=3"`
	Address     *string `json:"address" validate:"omitempty,min=3"`
	Telephone   *string `json:"telephone" validate:"omitempty,numeric,min=6"`
}

func (s *Seller) Validate(validate *validator.Validate) error {
    return validate.Struct(s)
}

func (s *SellerPatch) Validate(validate *validator.Validate) error {
    return validate.Struct(s)
}