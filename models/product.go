package models

type Product struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Price float64 `json:"price"`
}

type ProductRequest struct{
	Name string `json:"name"`
	Price float64 `json:"price"`
}