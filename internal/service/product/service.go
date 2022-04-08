package product

import "fmt"

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Product {
	return allProduct
}

func (s *Service) Get(idx int) (*Product, error) {

	if idx < 0 || idx >= len(allProduct) {
		return nil, fmt.Errorf("Index %v out of bounds", idx)
	}

	return &allProduct[idx], nil
}
