package buyers

import (
	"github.com/almarino_meli/grupo-5-wave-15/pkg/models/buyer"
)

// Mapa de valores Buyers - Data Temporal
type BuyerMap struct {
	db map[int]buyer.Buyer
}

// Instancia de mapa de valores
func NewBuyerMap(db map[int]buyer.Buyer) *BuyerMap {
	defaultDb := make(map[int]buyer.Buyer)
	if db != nil {
		defaultDb = db
	}
	return &BuyerMap{db: defaultDb}
}

// Implementaciones

// Encontrar todos los valores
func (r *BuyerMap) FindAll() (v map[int]buyer.Buyer, err error) {
	v = make(map[int]buyer.Buyer)
	for key, value := range r.db {
		v[key] = value
	}
	return
}

// Encontrar un solo registro por id
func (r *BuyerMap) FindById(vId int) (v buyer.Buyer, err error) {
	for key, value := range r.db {
		if vId == key {
			v = value
			break
		}
	}
	return
}

// Crear un nuevo registro con la informacion proporcionada
func (r *BuyerMap) SetNewBuyer(vNew buyer.Buyer) (v buyer.Buyer, err error) {
	idNew := len(r.db) + 1 // Obtener Nuevo ID (Es un incremental si ya existe un valor asignado
	vNew.Id = idNew        //Set ID
	r.db[idNew] = vNew     // Asignacion
	v = r.db[idNew]        // Callback
	return
}

// Actualizar el registro almacenado
func (r *BuyerMap) UpdateBuyer(vNew buyer.Buyer) (v buyer.Buyer, err error) {
	r.db[vNew.Id] = vNew
	return r.db[vNew.Id], nil
}

// Eliminar el registro
func (r *BuyerMap) DeleteBuyer(vId int) (ok bool, err error) {
	if r.db[vId].Id != 0 { // Existe en base de datos
		delete(r.db, vId)
		return true, nil
	} else {
		return false, nil
	}

}

// Encontrar registro por CardId
func (r *BuyerMap) FindByCardId(vCardId string) (v buyer.Buyer, err error) {
	for _, value := range r.db {
		if value.CardNumberId == vCardId {
			v = value
			break
		}
	}
	return
}
