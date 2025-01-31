package warehouses

import (
    "encoding/json"
    "os"
    "github.com/almarino_meli/grupo-5-wave-15/pkg/models/warehouse"
)

// NewWarehouseJSONFile is a function that returns a new instance of WarehouseJSONFile
func NewWarehouseJSONFile(path string) *WarehouseJSONFile {
    return &WarehouseJSONFile{
        path: path,
    }
}

// WarehouseJSONFile is a struct that implements the WarehouseLoader interface
type WarehouseJSONFile struct {
    // path is the path to the file that contains the Warehouses in JSON format
    path string
}

// Load is a method that loads the warehouses from a JSON file and returns a map of warehouses
func (l *WarehouseJSONFile) Load() (v map[int]warehouse.Warehouse, err error) {
    // open file
    file, err := os.Open(l.path)
    if err != nil {
        return
    }
    defer file.Close()

    // decode file
    var warehousesJSON []warehouse.WarehouseDTO
    err = json.NewDecoder(file).Decode(&warehousesJSON)
    if err != nil {
        return
    }

    // serialize warehouses
    v = make(map[int]warehouse.Warehouse)
    for _, warehouseDTO := range warehousesJSON {
        warehouseModel, err := warehouseDTO.ToModel()
        if err != nil {
            return nil, err
        }
        v[(warehouseModel.ID.ID)] = warehouseModel
    }
    return
}