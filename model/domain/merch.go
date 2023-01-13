package domain

type Merch struct {
	Id          int    `json:"id"`
	SellerId    int    `json:"seller_id"`
	Name        string `json:"name"`
	Images      string `json:"images"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Quantity    int    `json:"quantity"`
}