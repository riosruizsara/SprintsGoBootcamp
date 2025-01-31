package warehouses

import (
	"github.com/almarino_meli/grupo-5-wave-15/internal/repository/warehouses"
	"github.com/almarino_meli/grupo-5-wave-15/pkg/errors"
	"github.com/almarino_meli/grupo-5-wave-15/pkg/models/warehouse"
)

type WarehouseServiceDefault struct {
	rp warehouses.WarehouseRepository
}

func NewWarehouseServiceDefault(rp warehouses.WarehouseRepository) WarehouseServiceDefault {
	return WarehouseServiceDefault{rp: rp}
}

func (ws WarehouseServiceDefault) GetAll() (warehouses map[int]warehouse.Warehouse, err error) {
	warehouses, err = ws.rp.GetAll()
	if err != nil {
		// currently database in memory can't return an error
	}
	if len(warehouses) == 0 {
		err = &errors.NotFoundError{Message: "No warehouses were found"}
	}
	return
}

func (ws WarehouseServiceDefault) GetByID(id int) (warehouse warehouse.Warehouse, err error) {
	warehouse, err = ws.rp.GetByID(id)
	if err != nil {
		return warehouse, err
	}
	return
}

func (ws WarehouseServiceDefault) Create(dto warehouse.WarehouseDTO) (warehouse.Warehouse, error) {
	// Obtener todos los warehouses existentes
	warehouses, err := ws.rp.GetAll()
	if err != nil {
		return warehouse.Warehouse{}, err
	}

	// Verificar si ya existe un warehouse con el mismo código
	for _, w := range warehouses {
		if w.WarehouseCode.Code == dto.WarehouseCode {
			return warehouse.Warehouse{}, &errors.DuplicateError{
				Message: "Ya existe un warehouse con ese código",
			}
		}
	}

	return ws.rp.Create(dto)
}

func (ws WarehouseServiceDefault) Update(id int, dto warehouse.WarehouseDTO) (warehouse.Warehouse, error) {
	// Verificar si existe
	existingWarehouse, err := ws.rp.GetByID(id)
	if err != nil {
		return warehouse.Warehouse{}, &errors.NotFoundError{Message: "Warehouse not found"}
	}

	warehouses, err := ws.rp.GetAll()
	if err != nil {
		return warehouse.Warehouse{}, err
	}

	for _, w := range warehouses {
		if w.WarehouseCode.Code == dto.WarehouseCode {
			return warehouse.Warehouse{}, &errors.DuplicateError{
				Message: "Ya existe un warehouse con ese código",
			}
		}
	}

	// Actualizar solo los campos que vienen en el DTO
	if dto.Address != "" {
		var err error
		existingWarehouse.Address, err = warehouse.NewAddress(dto.Address)
		if err != nil {
			return warehouse.Warehouse{}, err
		}
	}
	if dto.Telephone != "" {
		var err error
		existingWarehouse.Telephone, err = warehouse.NewTelephone(dto.Telephone)
		if err != nil {
			return warehouse.Warehouse{}, err
		}
	}
	if dto.WarehouseCode != "" {
		var err error
		existingWarehouse.WarehouseCode, err = warehouse.NewWarehouseCode(dto.WarehouseCode)
		if err != nil {
			return warehouse.Warehouse{}, err
		}
	}
	if dto.MinimumCapacity != 0 {
		var err error
		existingWarehouse.MinimumCapacity, err = warehouse.NewCapacity(dto.MinimumCapacity)
		if err != nil {
			return warehouse.Warehouse{}, err
		}
	}
	if dto.MinimumTemperature != 0 {
		var err error
		existingWarehouse.MinimumTemperature, err = warehouse.NewTemperature(dto.MinimumTemperature)
		if err != nil {
			return warehouse.Warehouse{}, err
		}
	}

	// Guardar cambios en el repositorio
	return ws.rp.Update(id, existingWarehouse)
}

func (ws WarehouseServiceDefault) Delete(id int) error {
	// Verificar si existe
	_, err := ws.rp.GetByID(id)
	if err != nil {
		return &errors.NotFoundError{Message: "Warehouse not found"}
	}

	return ws.rp.Delete(id)
}
