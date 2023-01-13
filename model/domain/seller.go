package domain

type Seller struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Products string `json:"products"`
	Address  string `json:"address"`
	Balance  int    `json:"balance"`
}