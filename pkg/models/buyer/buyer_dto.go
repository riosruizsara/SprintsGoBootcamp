package buyer

import "github.com/go-playground/validator/v10"

type BuyerDoc struct {
	Id           int    `json:"id,omitempty"`
	CardNumberId string `json:"card_number_id" validate:"required,numeric"`
	FirstName    string `json:"first_name" validate:"required,min=2,max=50"`
	LastName     string `json:"last_name" validate:"required,min=2,max=50"`
}

type BuyerDocPatched struct {
	Id           int     `json:"id,omitempty"`
	CardNumberId *string `json:"card_number_id" validate:"omitempty,numeric"`
	FirstName    *string `json:"first_name" validate:"omitempty,min=2,max=50"`
	LastName     *string `json:"last_name" validate:"omitempty,min=2,max=50"`
}

func (b BuyerDoc) MapBuyerToModel() Buyer {
	return Buyer{
		Id:           b.Id,
		CardNumberId: b.CardNumberId,
		FirstName:    b.FirstName,
		LastName:     b.LastName,
	}
}

func (b BuyerDocPatched) MapBuyerToModel() Buyer {
	return Buyer{
		Id:           b.Id,
		CardNumberId: *b.CardNumberId,
		FirstName:    *b.FirstName,
		LastName:     *b.LastName,
	}
}

func (e *BuyerDoc) Validate(validate *validator.Validate) error {
	return validate.Struct(e)
}

func (e *BuyerDocPatched) Validate(validate *validator.Validate) error {
	return validate.Struct(e)
}
