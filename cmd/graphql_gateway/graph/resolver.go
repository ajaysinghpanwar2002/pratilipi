package graph

import "github.com/ajaysinghpanwar2002/pratilipi/cmd/graphql_gateway/graph/model"

type Resolver struct {
	ProductService ProductService
	OrderService   OrderService
}

type ProductService interface {
	GetProducts() ([]*model.Product, error)
	GetProductByID(id string) (*model.Product, error)
	CreateProduct(input model.ProductInput) (*model.Product, error)
}

type OrderService interface {
	GetOrders() ([]*model.Order, error)
	GetOrderByID(id string) (*model.Order, error)
	PlaceOrder(input model.OrderInput) (*model.Order, error)
}
