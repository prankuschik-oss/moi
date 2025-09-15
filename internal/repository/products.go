package repository

import "github.com/nicitapa/firstProgect/internal/models"

func (r *Repository) GetAllProducts() (products []models.Product, err error) {
	if err = r.db.Select(&products, `
		SELECT id, name, email, age, 
		FROM products
		ORDER BY id`); err != nil {
		return nil, err
	}

	return products, nil
}

func (r *Repository) GetProductByID(id int) (product models.Product, err error) {
	if err = r.db.Get(&product, `
		SELECT id, name, email, age, 
		FROM products
		WHERE id = $1`, id); err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func (r *Repository) CreateProduct(product models.Product) (err error) {
	_, err = r.db.Exec(`INSERT INTO products (name, email, age)
					VALUES ($1, $2, $3,)`,
		product.name,
		product.email,
		product.age)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateProductByID(product models.Product) (err error) {
	_, err = r.db.Exec(`
		UPDATE products SET name = $1, 
		                    email = $2, 
		                    age = $3,
		                    		                WHERE id = $5`,
		product.name,
		product.email,
		product.age,
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
