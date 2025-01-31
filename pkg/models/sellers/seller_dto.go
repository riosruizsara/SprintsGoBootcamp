package sellers

type SellerDTO struct {
	ID          int    `json:"id"`
	CompanyId   int    `json:"cid"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	Telephone   string `json:"telephone"`
}

func MapDTOToValueObject(dto SellerDTO) (s *Seller, e error) {
	s = &Seller{
		Id:          dto.ID,
		CompanyId:   dto.CompanyId,
		CompanyName: dto.CompanyName,
		Address:     dto.Address,
		Telephone:   dto.Telephone,
	}
	return
}

func MapToSellerDTO(seller Seller) SellerDTO {
	return SellerDTO{
		ID: seller.Id,
        CompanyId: seller.CompanyId,
        CompanyName: seller.CompanyName,
        Address: seller.Address,
        Telephone: seller.Telephone,
	}

}
