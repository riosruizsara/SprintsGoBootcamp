package buyers

import (
	"encoding/json"
	"os"

	"github.com/almarino_meli/grupo-5-wave-15/pkg/models/buyer"
)

// NewVehicleJSONFile is a function that returns a new instance of VehicleJSONFile
func NewBuyerJSONFile(path string) *BuyerJSONFile {
	return &BuyerJSONFile{
		path: path,
	}
}

// VehicleJSONFile is a struct that implements the LoaderVehicle interface
type BuyerJSONFile struct {
	// path is the path to the file that contains the vehicles in JSON format
	path string
}

// Load is a method that loads the vehicles
func (l *BuyerJSONFile) Load() (v map[int]buyer.Buyer, err error) {
	// open file
	file, err := os.Open(l.path)
	if err != nil {
		return
	}
	defer file.Close()

	// decode file
	var buyersJSON []buyer.BuyerDoc
	err = json.NewDecoder(file).Decode(&buyersJSON)
	if err != nil {
		return
	}

	// serialize vehicles
	v = make(map[int]buyer.Buyer)
	for _, by := range buyersJSON {
		v[by.Id] = buyer.Buyer{
			Id:           by.Id,
			CardNumberId: by.CardNumberId,
			FirstName:    by.FirstName,
			LastName:     by.LastName,
		}
	}
	return
}
