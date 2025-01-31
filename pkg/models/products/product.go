package products

import (
	"github.com/almarino_meli/grupo-5-wave-15/pkg/models"
	"github.com/go-playground/validator/v10"
)

type Product struct {
	ID                             models.ID  `validate:"required"`         // unique identifier
	ProductCode                    string     `validate:"required"`         // unique identifier
	Description                    string     `validate:"required,max=512"` // description of the product's characteristics
	Dimentions                     Dimentions `validate:"required"`
	NetWeight                      float64    `validate:"required,gt=0,lt=1000"` // expressed in kg
	ExpirationRate                 int        `validate:"required,gt=0"`
	RecommendedFreezingTemperature float64    `validate:"required,gte=-512,lte=512"`
	FreezingRate                   int        `validate:"required,gt=0"`
	ProductTypeID                  models.ID  `validate:"required"` // related to a particular "ProductType"
	SellerID                       models.ID  `validate:"required"` // related to a particular "Seller" (foreign key), if set to 0 it means that the product was loaded without a seller
}

type ProductPatch struct {
	ProductCode                    *string  `json:"product_code" validate:"omitempty,required"`
	Description                    *string  `json:"description" validate:"omitempty,required,max=512"`
	Width                          *float64 `json:"width" validate:"omitempty,required,gt=0"`
	Height                         *float64 `json:"height" validate:"omitempty,required,gt=0"`
	Length                         *float64 `json:"length" validate:"omitempty,required,gt=0"`
	NetWeight                      *float64 `json:"net_weight" validate:"omitempty,required,gt=0,lt=1000"`
	ExpirationRate                 *int     `json:"expiration_rate" validate:"omitempty,required,gt=0"`
	RecommendedFreezingTemperature *float64 `json:"recommended_freezing_temperature" validate:"omitempty,required,gte=-512,lte=512"`
	FreezingRate                   *int     `json:"freezing_rate" validate:"omitempty,required,gt=0"`
	ProductTypeID                  *int     `json:"product_type_id" validate:"omitempty,required,gte=0"`
	SellerID                       *int     `json:"seller_id" validate:"omitempty,required,gte=0"`
}

type ProductType struct {
	ID   int
	Name string
}

type ProductCode struct {
	Code string
}

type Description struct {
	Description string
}

type Dimentions struct {
	Width  float64 `validate:"required,gt=0"` // expressed in cm
	Height float64 `validate:"required,gt=0"` // expressed in cm
	Length float64 `validate:"required,gt=0"` // expressed in cm
}

type Weight struct {
	NetWeight float64 `validate:"required,gt=0,lt=1000"` // expressed in kg
}

type ExpirationRate struct {
	Rate int `validate:"required,gt=0"`
}

type Temperature struct {
	Temperature float64 `validate:"required,gte=-512,lte=512"` // expressed in degrees celsius
}

type FreezingRate struct {
	Rate int `validate:"required,gt=0"`
}

func (p Product) GetID() int {
	return p.ID.ID
}

func (p Product) ToDTO() ProductDTO {
	return ProductDTO{
		ID:                             p.GetID(),
		ProductCode:                    p.ProductCode,
		Description:                    p.Description,
		Width:                          p.Dimentions.Width,
		Height:                         p.Dimentions.Height,
		Length:                         p.Dimentions.Length,
		NetWeight:                      p.NetWeight,
		ExpirationRate:                 p.ExpirationRate,
		RecommendedFreezingTemperature: p.RecommendedFreezingTemperature,
		FreezingRate:                   p.FreezingRate,
		ProductTypeID:                  p.ProductTypeID.ID,
		SellerID:                       &p.SellerID.ID,
	}
}

func (product *Product) Validate(validate *validator.Validate) error {
	return validate.Struct(product)
}

func (patch *ProductPatch) Validate(validate *validator.Validate) error {
	return validate.Struct(patch)
}
