package warehouses

import (
    "github.com/almarino_meli/grupo-5-wave-15/pkg/models/warehouse"
)

// WarehouseRepository is an interface that represents a repository that can be used to interact with warehouses
type WarehouseRepository interface {
    GetAll() (warehouses map[int]warehouse.Warehouse, err error)
    GetByID(id int) (warehouse warehouse.Warehouse, err error)
    Create(dto warehouse.WarehouseDTO) (warehouse warehouse.Warehouse, err error)
    Update(id int, warehouse warehouse.Warehouse) (warehouse.Warehouse, error)
    Delete(id int) error
}