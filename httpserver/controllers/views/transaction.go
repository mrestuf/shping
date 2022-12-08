package views

import "time"

type CreateTransaction struct {
	TotalPrice   uint   `json:"total_price"`
	Quantity     uint   `json:"quantity"`
	ProductTitle string `json:"product_title"`
}

type GetMyTransaction struct {
	Id         uint `json:"id"`
	ProductId  uint `json:"product_id"`
	UserId     uint `json:"user_id"`
	Quantity   uint `json:"quantity"`
	TotalPrice uint `json:"total_price"`
	Product    []ProductTransaction
}

type ProductTransaction struct {
	Id         uint      `json:"id"`
	Title      string    `json:"title"`
	Price      uint      `json:"price"`
	Stock      uint      `json:"stock"`
	CategoryId uint      `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
