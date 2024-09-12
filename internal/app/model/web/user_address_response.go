package web

type UserAddressResponse struct {
	ID              int    `json:"id"`
	Label           string `json:"label"`
	AddressLine1    string `json:"address_line_1"`
	AddressLine2    string `json:"address_line_2"`
	City            string `json:"city"`
	State           string `json:"state"`
	PostalCode      string `json:"postal_code"`
	Country         string `json:"country"`
	AddressableID   int    `json:"addressable_id"`
	AddressableType string `json:"addressable_type"`
}
