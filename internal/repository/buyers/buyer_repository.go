package buyers

import "github.com/almarino_meli/grupo-5-wave-15/pkg/models/buyer"

type BuyerRepository interface {
	FindAll() (v map[int]buyer.Buyer, err error)
	FindById(vId int) (v buyer.Buyer, err error)
	FindByCardId(vCardId string) (v buyer.Buyer, err error)
	SetNewBuyer(vNew buyer.Buyer) (v buyer.Buyer, err error)
	UpdateBuyer(vNew buyer.Buyer) (v buyer.Buyer, err error)
	DeleteBuyer(vId int) (vv bool, err error)
}
