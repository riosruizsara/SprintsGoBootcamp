package products

import (
	"github.com/almarino_meli/grupo-5-wave-15/pkg/models"
	"github.com/go-playground/validator/v10"
)

type ProductDTO struct {
	ID                             int     `json:"id,omitempty"`
	ProductCode                    string  `json:"product_code"`
	Description                    string  `json:"description"`
	Width                          float64 `json:"width"`
	Height                         float64 `json:"height"`
	Length                         float64 `json:"length"`
	NetWeight                      float64 `json:"net_weight"`
	ExpirationRate                 int     `json:"expiration_rate"`
	RecommendedFreezingTemperature float64 `json:"recommended_freezing_temperature"`
	FreezingRate                   int     `json:"freezing_rate"`
	ProductTypeID                  int     `json:"product_type_id"`
	SellerID                       *int    `json:"seller_id,omitempty"` // may or may not come
}

func (p *ProductDTO) ToModel(validate *validator.Validate) (Product, error) {
	modelID := models.ID{ID: p.ID}
	modelProductTypeID := models.ID{ID: p.ProductTypeID}
	modelSellerID := models.ID{}
	if p.SellerID != nil {
		modelSellerID = models.ID{ID: *p.SellerID}
	}

	product := Product{
		ID:                             modelID,
		ProductCode:                    p.ProductCode,
		Description:                    p.Description,
		Dimentions:                     Dimentions{Width: p.Width, Height: p.Height, Length: p.Length},
		NetWeight:                      p.NetWeight,
		ExpirationRate:                 p.ExpirationRate,
		RecommendedFreezingTemperature: p.RecommendedFreezingTemperature,
		FreezingRate:                   p.FreezingRate,
		ProductTypeID:                  modelProductTypeID,
		SellerID:                       modelSellerID,
	}

	// Validate the product
	if validate != nil {
		if err := validate.Struct(product); err != nil {
			return Product{}, err
		}
	}

	return product, nil
}
