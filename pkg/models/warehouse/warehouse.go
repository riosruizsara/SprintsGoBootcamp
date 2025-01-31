package warehouse

import (
    "github.com/almarino_meli/grupo-5-wave-15/pkg/errors"
    "github.com/almarino_meli/grupo-5-wave-15/pkg/models"
    "github.com/go-playground/validator/v10"
)

type Warehouse struct {
    ID                models.ID    `validate:"required"`
    WarehouseCode     WarehouseCode `validate:"required"`
    Address           Address      `validate:"required,max=512"`
    Telephone         Telephone    `validate:"required"`
    MinimumCapacity   Capacity     `validate:"required"`
    MinimumTemperature Temperature `validate:"required"`
}

type WarehouseCode struct {
    Code string
}

func NewWarehouseCode(code string) (WarehouseCode, error) {
    if code == "" {
        return WarehouseCode{}, &errors.ValidationError{Message: "WarehouseCode must not be empty"}
    }
    return WarehouseCode{Code: code}, nil
}

type Address struct {
    Address string
}

const MaxAddressLength = 512

func NewAddress(address string) (Address, error) {
    if address == "" {
        return Address{}, &errors.ValidationError{Message: "Address must not be empty"}
    }
    if len(address) > MaxAddressLength {
        return Address{}, &errors.ValidationError{Message: "Address must not be longer than 512 characters"}
    }
    return Address{Address: address}, nil
}

type Telephone struct {
    Telephone string
}

const MaxTelephoneLength = 15

func NewTelephone(telephone string) (Telephone, error) {
    if telephone == "" {
        return Telephone{}, &errors.ValidationError{Message: "Telephone must not be empty"}
    }
    if len(telephone) > MaxTelephoneLength {
        return Telephone{}, &errors.ValidationError{Message: "Telephone must not be longer than 15 characters"}
    }
    return Telephone{Telephone: telephone}, nil
}

type Capacity struct {
    Capacity int
}

func NewCapacity(capacity int) (Capacity, error) {
    if capacity <= 0 {
        return Capacity{}, &errors.ValidationError{Message: "Capacity must be greater than 0"}
    }
    return Capacity{Capacity: capacity}, nil
}

type Temperature struct {
    Temperature float64
}

func NewTemperature(temperature float64) (Temperature, error) {
    if temperature < -273.15 {
        return Temperature{}, &errors.ValidationError{Message: "Temperature must not be less than -273.15°C"}
    }
    return Temperature{Temperature: temperature}, nil
}

func (w Warehouse) GetID() int {
    return w.ID.ID
}

func ValidateWarehouse(warehouse Warehouse) error {
    validate := validator.New()
    return validate.Struct(warehouse)
}