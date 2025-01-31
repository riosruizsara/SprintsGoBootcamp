package buyers

import (
	"github.com/almarino_meli/grupo-5-wave-15/pkg/errors"
	"github.com/almarino_meli/grupo-5-wave-15/pkg/models/buyer"

	rpbuy "github.com/almarino_meli/grupo-5-wave-15/internal/repository/buyers"
)

// Estructura Servicio
type BuyerDefault struct {
	rp rpbuy.BuyerRepository
}

// Instancia Servicio
func NewBuyerDefault(rp rpbuy.BuyerRepository) *BuyerDefault {
	return &BuyerDefault{rp: rp}
}

// Implementaciones

// Encontrar todos los valores
func (r *BuyerDefault) FindAll() (v map[int]buyer.Buyer, err error) {
	return r.rp.FindAll()
}

// Encontrar un solo registro por id
func (r *BuyerDefault) FindById(vId int) (v buyer.Buyer, err error) {
	d, err := r.rp.FindById(vId)
	if d.Id == 0 {
		return d, &errors.NotFoundError{Message: "Buyer not found"}
	} else {
		v = d
	}
	return
}

// Crear un nuevo registro con la informacion proporcionada
func (r *BuyerDefault) SetNewBuyer(vNew buyer.Buyer) (v buyer.Buyer, err error) {
	isValid, _ := r.VerifyUniqueCardIdFromBuyer(vNew.CardNumberId)
	if isValid {
		return r.rp.SetNewBuyer(vNew)
	} else {
		return v, &errors.DuplicateError{Message: "Invalid Uniqueness: CardNumberID found in database"}
	}
}

// Actualizar el registro almacenado
func (r *BuyerDefault) UpdateBuyer(vId int, vNew buyer.BuyerDocPatched) (updatedBuyer buyer.Buyer, err error) {
	updatedBuyer, err = r.FindById(vId)
	if err != nil {
		return updatedBuyer, &errors.NotFoundError{Message: "Buyer not found"}
	}
	if vNew.CardNumberId != nil {
		isValid, _ := r.VerifyUniqueCardIdFromBuyer(*vNew.CardNumberId)
		if isValid {
			updatedBuyer.CardNumberId = *vNew.CardNumberId
		} else {
			return updatedBuyer, &errors.DuplicateError{Message: "Invalid Uniqueness: CardNumberID found in database"}
		}
	}
	if vNew.LastName != nil {
		updatedBuyer.LastName = *vNew.LastName
	}
	if vNew.FirstName != nil {
		updatedBuyer.FirstName = *vNew.FirstName
	}
	r.rp.UpdateBuyer(updatedBuyer)
	return
}

// Eliminar el registro
func (r *BuyerDefault) DeleteBuyer(vId int) (err error) {
	d, err := r.rp.DeleteBuyer(vId)
	if !d || err != nil {
		return &errors.NotFoundError{Message: "Buyer not found"}
	}
	return nil
}

// Validar que atributo Card ID sea unico
func (r *BuyerDefault) VerifyUniqueCardIdFromBuyer(vCardId string) (isValid bool, err error) {
	v, err := r.rp.FindByCardId(vCardId)
	if v.CardNumberId == "" {
		return true, nil
	}
	return
}
