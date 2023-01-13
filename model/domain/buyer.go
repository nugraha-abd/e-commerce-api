package domain

type Buyer struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}