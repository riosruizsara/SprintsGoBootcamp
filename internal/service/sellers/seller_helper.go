package sellers

import (
	"github.com/almarino_meli/grupo-5-wave-15/pkg/models/sellers"
)


func applyPatch(seller *sellers.Seller, patch sellers.SellerPatch) error {
	if patch.CompanyId != nil {
		seller.CompanyId = *patch.CompanyId
	}
	if patch.CompanyName != nil {
		seller.CompanyName = *patch.CompanyName
	}
	if patch.Address != nil {
		seller.Address = *patch.Address
	}
	if patch.Telephone != nil {
		seller.Telephone = *patch.Telephone
	}
	return nil
}