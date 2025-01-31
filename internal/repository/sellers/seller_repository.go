package sellers

import "github.com/almarino_meli/grupo-5-wave-15/pkg/models/sellers"

type SellerRepository interface {
	Create(s sellers.Seller)(res sellers.Seller, err error)
	GetAll() (res map[int]sellers.Seller, err error)
	GetById(id int) (res sellers.Seller, err error)
	Update(s sellers.Seller)(res sellers.Seller, err error)
	Delete(id int) (err error)
}

