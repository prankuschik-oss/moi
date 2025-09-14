package sevice

import "github.com/nicitapa/firstProgect/internal/models"

func (s *Service) GetAllProducts() (products []models.Product, err error) {
	products, err = s.repository.GetAllProducts()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s *Service) GetProductByID(id int) (product models.Product, err error) {
	product, err = s.repository.GetProductByID(id)
	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func (s *Service) CreateProduct(product models.Product) (err error) {
	err = s.repository.CreateProduct(product)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) UpdateProductByID(product models.Product) (err error) {
	err = s.repository.UpdateProductByID(product)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteProductByID(id int) (err error) {
	err = s.repository.DeleteProductByID(id)
	if err != nil {
		return err
	}

	return nil
}
