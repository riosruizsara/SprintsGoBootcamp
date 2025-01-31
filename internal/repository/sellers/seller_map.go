package sellers

import (
	customErrors "github.com/almarino_meli/grupo-5-wave-15/pkg/errors"
	"github.com/almarino_meli/grupo-5-wave-15/pkg/models/sellers"
)

func NewSellerMap(db map[int]sellers.Seller) *SellerMap {
	defaultDb := make(map[int]sellers.Seller)
	if db != nil {
		defaultDb = db
	}
	return &SellerMap{db: defaultDb}
}

type SellerMap struct {
	db map[int]sellers.Seller
}

func (r *SellerMap) Create(s sellers.Seller) (res sellers.Seller, err error) {
	if r.companyExists(s.CompanyId) == true {
		err = &customErrors.ConflictError{Message: "repository: Company already exists"}
		return
	}
	var id int = r.getNextId()
	s.Id = id
	r.db[id] = s
	return s, nil
}

func (r *SellerMap) companyExists(cid int) bool {
	for _, seller := range r.db {
		if seller.CompanyId == cid {
			return true
		}
	}
	return false
}

func (r *SellerMap) getNextId() (maxId int) {
	maxId = 1
	for _, v := range r.db {
		if v.Id > maxId {
			maxId = v.Id
		}
	}
	maxId = maxId + 1
	return
}


func (r *SellerMap) GetAll()(res map[int]sellers.Seller, err error){
	return r.db, nil
}

func (r *SellerMap) GetById(id int) (res sellers.Seller, err error){
	if _, exists := r.db[id]; exists{
		return r.db[id], nil
	} 
	return sellers.Seller{}, &customErrors.NotFoundError{Message: "repository: Seller does not exist"}
}

func (r *SellerMap) Delete(id int) (err error){
	if _, exists := r.db[id]; exists{
		delete(r.db, id)
		return nil
	} 
	return &customErrors.NotFoundError{Message: "repository: Seller does not exist"}
}

func (r *SellerMap) Update(s sellers.Seller) (res sellers.Seller, err error) {
	for _, seller := range r.db {
		if seller.CompanyId == s.CompanyId && seller.Id != s.Id {
			err = &customErrors.ConflictError{Message: "repository: Company already exists"}
			return sellers.Seller{}, err
		}
	}
	r.db[s.Id] = s
	res = s
	return
}
