package web

type StoreRequest struct {
	Name        string `json:"name" validate:"required,min=3,max=100"`
	OwnerID     int    `json:"owner_id"`
	AddressID   int    `json:"address_id"`
	Description string `json:"description" validate:"required,min=3,max=1000"`
	Photo       string `json:"photo"`
}

type StoreCreateRequest struct {
	StoreRequest   StoreRequest       `json:"store"`
	AddressRequest UserAddressRequest `json:"address"`
}
