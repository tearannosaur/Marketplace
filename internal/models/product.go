package models

import (
	"errors"
	er "mp/internal/errors"

	"github.com/google/uuid"
)

type Product struct {
	ProductId    uuid.UUID `json:"product_id" db:"product_id"`
	ProductName  string    `json:"product_name" db:"product_name"`
	Price        float64   `json:"product_price" db:"product_price"`
	CategoryName string    `json:"product_category" db:"product_category"`
	Description  string    `json:"product_description" db:"product_description"`
}

type ProductResponce struct {
	ProductName  string  `json:"product_name"`
	Price        float64 `json:"product_price"`
	CategoryName string  `json:"product_category"`
	Description  string  `json:"product_description"`
}

func NewProduct(p ProductResponce) (Product, error) {
	err := ValidateProductData(p)
	if err != nil {
		return Product{}, err
	}
	product := Product{
		ProductId:    uuid.New(),
		Price:        p.Price,
		CategoryName: p.CategoryName,
		Description:  p.Description,
	}
	return product, nil
}

func ValidateProductData(p ProductResponce) error {
	if p.CategoryName == "" || p.Price <= 0 {
		return errors.New(er.IncorrectProductDataErr)
	}
	return nil
}

func (p *Product) ProductChangePrice(amount float64) error {
	if amount < 0 {
		return errors.New(er.IncorrectProductPriceErr)
	}
	p.Price = amount
	return nil
}

func (p *Product) ProductChangeName(newName string) error {
	if newName == "" {
		return errors.New(er.IncorrectProductDataErr)
	}
	p.ProductName = newName
	return nil
}

func (p *Product) ProductChangeCategory(newCategory string) error {
	if newCategory == "" {
		return errors.New(er.IncorrectProductDataErr)
	}
	p.CategoryName = newCategory
	return nil
}
