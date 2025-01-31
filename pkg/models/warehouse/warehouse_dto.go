package warehouse

import "github.com/almarino_meli/grupo-5-wave-15/pkg/models"

type WarehouseDTO struct {
    ID                *int    `json:"id,omitempty"`
    WarehouseCode     string  `json:"warehouse_code"`
    Address           string  `json:"address"`
    Telephone         string  `json:"telephone"`
    MinimumCapacity   int     `json:"minimum_capacity"`
    MinimumTemperature float64 `json:"minimum_temperature"`
}

func (w *WarehouseDTO) ToModel() (Warehouse, error) {
    modelID, err := models.NewID(*w.ID)
    if err != nil {
        return Warehouse{}, err
    }
    modelWarehouseCode, err := NewWarehouseCode(w.WarehouseCode)
    if err != nil {
        return Warehouse{}, err
    }
    modelAddress, err := NewAddress(w.Address)
    if err != nil {
        return Warehouse{}, err
    }
    modelTelephone, err := NewTelephone(w.Telephone)
    if err != nil {
        return Warehouse{}, err
    }
    modelMinimumCapacity, err := NewCapacity(w.MinimumCapacity)
    if err != nil {
        return Warehouse{}, err
    }
    modelMinimumTemperature, err := NewTemperature(w.MinimumTemperature)
    if err != nil {
        return Warehouse{}, err
    }
    return Warehouse{
        ID:                modelID,
        WarehouseCode:     modelWarehouseCode,
        Address:           modelAddress,
        Telephone:         modelTelephone,
        MinimumCapacity:   modelMinimumCapacity,
        MinimumTemperature: modelMinimumTemperature,
    }, nil
}

func (w *Warehouse) ToDTO() WarehouseDTO {
    return WarehouseDTO{
        ID:                 &w.ID.ID,
        WarehouseCode:     w.WarehouseCode.Code,
        Address:           w.Address.Address,
        Telephone:         w.Telephone.Telephone,
        MinimumCapacity:   w.MinimumCapacity.Capacity,
        MinimumTemperature: w.MinimumTemperature.Temperature,
    }
}