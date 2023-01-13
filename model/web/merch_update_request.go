package web

type MerchUpdateRequest struct {
	Id          int
	Name        string
	Description string
	Images      string
	Price       int
	Quantity    int
}