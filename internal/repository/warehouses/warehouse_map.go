package warehouses

import (
    "errors"
    "github.com/almarino_meli/grupo-5-wave-15/pkg/models/warehouse"
)

// NewWarehouseMap is a function that returns a new instance of WarehouseMap
func NewWarehouseMap(db map[int]warehouse.Warehouse) *WarehouseMap {
    // default db
    defaultDb := make(map[int]warehouse.Warehouse)
    if db != nil {
        defaultDb = db
    }
    return &WarehouseMap{db: defaultDb}
}

// WarehouseMap is a struct that represents a warehouses repository with a memory map implementation
type WarehouseMap struct {
    // db is a map of warehouses
    db map[int]warehouse.Warehouse
}

// GetAll is a method that returns all the warehouses
func (wm *WarehouseMap) GetAll() (warehouses map[int]warehouse.Warehouse, err error) {
    warehouses = wm.db
    return
}

// GetByID is a method that returns a warehouse by its ID
func (wm *WarehouseMap) GetByID(id int) (warehouse warehouse.Warehouse, err error) {
    warehouse, ok := wm.db[id]
    if !ok {
        err = errors.New("warehouse not found")
        return warehouse, err
    }
    return
}

// Create es un método que crea un nuevo warehouse
func (wm *WarehouseMap) Create(dto warehouse.WarehouseDTO) (warehouse warehouse.Warehouse, err error) {
    newID := len(wm.db) + 1
    dto.ID = &newID
    warehouseModel, err := dto.ToModel()
    if err != nil {
        return warehouse, err
    }
    wm.db[newID] = warehouseModel
    return warehouseModel, nil
}

func (wm *WarehouseMap) Update(id int, warehouse warehouse.Warehouse) (warehouse.Warehouse, error) {
    // Verificar si existe el warehouse
    _, exists := wm.db[id]
    if !exists {
        return warehouse, errors.New("warehouse not found")
    }

    // Actualizar en el map
    wm.db[id] = warehouse
    return warehouse, nil
}

func (wm *WarehouseMap) Delete(id int) error {
    if _, exists := wm.db[id]; !exists {
        return errors.New("warehouse not found")
    }
    delete(wm.db, id)
    return nil
}