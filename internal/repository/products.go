package repository

import "github.com/nicitapa/firstProgect/internal/models"

func (r *Repository) GetAllProducts() (products []models.Product, err error) {
	if err = r.db.Select(&products, `
		SELECT id, product_name, manufacturer, product_count, price 
		FROM products
		ORDER BY id`); err != nil {
		return nil, err
	}

	return products, nil
}

func (r *Repository) GetProductByID(id int) (product models.Product, err error) {
	if err = r.db.Get(&product, `
		SELECT id, product_name, manufacturer, product_count, price 
		FROM products
		WHERE id = $1`, id); err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func (r *Repository) CreateProduct(product models.Product) (err error) {
	_, err = r.db.Exec(`INSERT INTO products (product_name, manufacturer, price, product_count)
					VALUES ($1, $2, $3, $4)`,
		product.ProductName,
		product.Manufacturer,
		product.Price,
		product.ProductCount)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateProductByID(product models.Product) (err error) {
	_, err = r.db.Exec(`
		UPDATE products SET product_name = $1, 
		                    manufacturer = $2, 
		                    price = $3,
		                    product_count = $4
		                WHERE id = $5`,
		product.ProductName,
		product.Manufacturer,
		product.Price,
		product.ProductCount,
		product.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeleteProductByID(id int) (err error) {
	_, err = r.db.Exec(`DELETE FROM products WHERE id = $1`, id)
	if err != nil {
		return err
	}

	return nil
}
