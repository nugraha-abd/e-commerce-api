package web

type MerchCreateRequest struct {
	Name        string
	Description string
	Images      string
	Price       int
	Quantity    int
}