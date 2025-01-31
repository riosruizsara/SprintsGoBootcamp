package buyers

import "github.com/almarino_meli/grupo-5-wave-15/pkg/models/buyer"

type BuyerLoader interface {
	Load() (v map[int]buyer.Buyer, err error)
}
