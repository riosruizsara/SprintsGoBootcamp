package sellers

import (
	sellersRp "github.com/almarino_meli/grupo-5-wave-15/internal/repository/sellers"
	"github.com/almarino_meli/grupo-5-wave-15/pkg/errors"
	"github.com/almarino_meli/grupo-5-wave-15/pkg/models/sellers"
)

func NewSellerDefault(rp sellersRp.SellerRepository) *SellerDefault {
	return &SellerDefault{rp: rp}
}

type SellerDefault struct {
	rp sellersRp.SellerRepository
}

func (sv *SellerDefault) Create(s sellers.Seller) (res sellers.Seller, err error) {
	res, err = sv.rp.Create(s)
	return
}

func (sv *SellerDefault) GetAll() (res map[int]sellers.Seller, err error) {
	res, err = sv.rp.GetAll()
	return
}

func (sv *SellerDefault) GetById(id int) (res sellers.Seller, err error) {
	res, err = sv.rp.GetById(id)
	return
}

func (sv *SellerDefault) Delete(id int) (err error) {
	err = sv.rp.Delete(id)
	return
}

func (sv *SellerDefault) Update(s sellers.SellerPatch, id int) (res sellers.Seller, err error) {
	existingSeller, err := sv.rp.GetById(id)
	if err != nil {
		return sellers.Seller{}, err
	}
	if err := applyPatch(&existingSeller, s); err != nil {
		return sellers.Seller{}, &errors.ValidationError{Message: err.Error()}
	}
	res, err = sv.rp.Update(existingSeller)
	if err != nil {
		return sellers.Seller{}, err
	}

	return res, nil
}
