package buyer

type Buyer struct {
	Id           int
	CardNumberId string
	FirstName    string
	LastName     string
}

func (b Buyer) MapBuyerToDto() BuyerDoc {
	return BuyerDoc{
		Id:           b.Id,
		CardNumberId: b.CardNumberId,
		FirstName:    b.FirstName,
		LastName:     b.LastName,
	}
}
