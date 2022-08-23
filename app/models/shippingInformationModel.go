package models

// ShippingInformation data struct
type ShippingInformation struct {
	ID            uint   `json:"id"`
	RecipientName string `json:"recipient_name"`
	PhoneNumber   string `json:"phone_number"`
	Address       string `json:"address"`
	AddressLine   string `json:"address_line"`
	City          string `json:"city"`
	Country       string `json:"country"`
}

func (ShippingInformation) TableName() string {
	return "shipping_information"
}
