package buyers

import "github.com/almarino_meli/grupo-5-wave-15/pkg/models/buyer"

type BuyerService interface {
	FindAll() (v map[int]buyer.Buyer, err error)
	FindById(vId int) (v buyer.Buyer, err error)
	SetNewBuyer(vNew buyer.Buyer) (v buyer.Buyer, err error)
	UpdateBuyer(vId int, vNew buyer.BuyerDocPatched) (v buyer.Buyer, err error)
	DeleteBuyer(vId int) (err error)
	VerifyUniqueCardIdFromBuyer(vCardId string) (isValid bool, err error)
}
