package sellers

import "github.com/almarino_meli/grupo-5-wave-15/pkg/models/sellers"

type SellerLoader interface {
	// Load is a method that loads the sellers
	Load() (v map[int]sellers.Seller, err error)
}
