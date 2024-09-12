package web

type UserAddressRequest struct {
	ID              int    `json:"id"`
	Label           string `json:"label" validate:"required,min=3,max=100"`
	AddressLine1    string `json:"address_line_1" validate:"required,min=3,max=100"`
	AddressLine2    string `json:"address_line_2" validate:"required,min=3,max=100"`
	City            string `json:"city" validate:"required,min=3,max=100"`
	State           string `json:"state"`
	PostalCode      string `json:"postal_code" validate:"required,min=3,max=100"`
	Country         string `json:"country" validate:"required,min=3,max=100"`
	AddressableType string `json:"addressable_type" validate:"required,oneof=user store"`
}
