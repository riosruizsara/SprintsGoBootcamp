package warehouses

import (
    "github.com/almarino_meli/grupo-5-wave-15/pkg/models/warehouse"
)

// WarehouseService is an interface that represents a service that can be used to interact with warehouses
type WarehouseService interface {
    GetAll() (warehouses map[int]warehouse.Warehouse, err error)
    GetByID(id int) (warehouse warehouse.Warehouse, err error)
    Create(dto warehouse.WarehouseDTO) (warehouse.Warehouse, error)
    Update(id int, dto warehouse.WarehouseDTO) (warehouse.Warehouse, error)
    Delete(id int) error
}