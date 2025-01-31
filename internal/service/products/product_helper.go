package products

import (
	"reflect"

	"github.com/almarino_meli/grupo-5-wave-15/pkg/models"
	"github.com/almarino_meli/grupo-5-wave-15/pkg/models/products"
)

// This file contains helper functions for the products service

// Helper function to apply patch data to the existing product
func applyPatch(product *products.Product, patch products.ProductPatch) {
	productValue := reflect.ValueOf(product).Elem()
	patchValue := reflect.ValueOf(patch)

	for i := 0; i < patchValue.NumField(); i++ {
		patchField := patchValue.Field(i)
		if !patchField.IsNil() {
			fieldName := patchValue.Type().Field(i).Name
			productField := productValue.FieldByName(fieldName)
			if productField.IsValid() && productField.CanSet() {
				if fieldName == "ProductTypeID" || fieldName == "SellerID" {
					// Handle models.ID fields separately
					productField.Set(reflect.ValueOf(models.ID{ID: patchField.Elem().Interface().(int)}))
				} else {
					productField.Set(reflect.Indirect(patchField))
				}
			} else if fieldName == "Width" || fieldName == "Height" || fieldName == "Length" {
				// Handle Dimentions fields separately
				dimentionsField := productValue.FieldByName("Dimentions")
				if dimentionsField.IsValid() && dimentionsField.CanSet() {
					dimentions := dimentionsField.Interface().(products.Dimentions)
					if fieldName == "Width" {
						dimentions.Width = patchField.Elem().Interface().(float64)
					} else if fieldName == "Height" {
						dimentions.Height = patchField.Elem().Interface().(float64)
					} else if fieldName == "Length" {
						dimentions.Length = patchField.Elem().Interface().(float64)
					}
					dimentionsField.Set(reflect.ValueOf(dimentions))
				}
			}
		}
	}
}
