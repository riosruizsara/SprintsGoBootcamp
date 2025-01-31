package sellers

import (
	"encoding/json"
	"os"

	"github.com/almarino_meli/grupo-5-wave-15/pkg/models/sellers"
)

func NewSellerJSONFile(path string) *SellerJSONFile {
	return &SellerJSONFile{
		path: path,
	}
}

type SellerJSONFile struct {
	path string
}

// Load is a method that loads the Sellers
func (l *SellerJSONFile) Load() (s map[int]sellers.Seller, err error) {
	file, err := os.Open(l.path)
	if err != nil {
		return
	}
	defer file.Close()

	var SellersJSON []sellers.SellerDTO
	err = json.NewDecoder(file).Decode(&SellersJSON)
	if err != nil {
		return
	}

	// serialize Sellers
	s = make(map[int]sellers.Seller)
	for _, seller := range SellersJSON {
		s[seller.ID] = sellers.Seller{
			Id:          seller.ID,
			CompanyId:   seller.CompanyId,
			CompanyName: seller.CompanyName,
			Address:     seller.Address,
			Telephone:   seller.Telephone,
		}
	}
	return
}
