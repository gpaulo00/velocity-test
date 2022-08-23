package models

// Warehouse data struct
type Warehouse struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Address     string  `json:"address"`
	AddressLine string  `json:"address_line"`
	City        string  `json:"city"`
	Country     string  `json:"country"`
	Lat         float64 `json:"lat"`
	Lng         float64 `json:"lng"`
}

func (Warehouse) TableName() string {
	return "warehouse"
}
