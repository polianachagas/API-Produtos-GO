package usecase

import "go-api/model"

type ProductUsecase struct {
	//repository
}

// func que inicializa a estrutura
func NewProductUseCase() ProductUsecase {
	return ProductUsecase{}
}

// func responsável por tratar das regras de negócio da rota de GetProducts
func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	return []model.Product{}, nil
}
