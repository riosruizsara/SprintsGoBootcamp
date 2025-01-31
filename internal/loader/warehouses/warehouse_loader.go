package warehouses

import (
    "github.com/almarino_meli/grupo-5-wave-15/pkg/models/warehouse"
)

// WarehouseLoader is an interface that represents the loader for warehouses
type WarehouseLoader interface {
    // Load is a method that loads the warehouses
    Load() (v map[int]warehouse.Warehouse, err error)
}