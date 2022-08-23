package models

import (
	"time"
)

// Order data struct
type Order struct {
	ID                    uint                `json:"id"`
	CreatedAt             time.Time           `json:"created_at"`
	WarehouseID           int                 `json:"warehouse_id"`
	Warehouse             Warehouse           `json:"warehouse"`
	ShippingInformationID int                 `json:"shipping_information_id"`
	ShippingInformation   ShippingInformation `json:"shipping_information"`
	TotalKms              float64             `json:"total_kms"`
}

func (Order) TableName() string {
	return "order"
}

func (h Order) GetOrder(id string) (*Order, error) {
	Init()
	db := GetDB()
	var item Order

	// get from database
	err := db.
		Preload("Warehouse").Preload("ShippingInformation").
		First(&item, id).Error
	if err != nil {
		return nil, err
	}

	// return
	return &item, nil
}

func (h Order) UpdateOrder(item *Order) error {
	Init()
	db := GetDB()

	// Proccess Update
	if err := db.Save(item).Error; err != nil {
		return err
	}
	return nil
}
