package web

type StoreResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	OwnerID     int    `json:"owner_id"`
	AddressID   int    `json:"address_id"`
	Description string `json:"description"`
	Photo       string `json:"photo"`
	CreatedAt   string `json:"created_at"`
}

type StoreCreateResponse struct {
	StoreResponse   StoreResponse       `json:"store"`
	AddressResponse UserAddressResponse `json:"address"`
}
