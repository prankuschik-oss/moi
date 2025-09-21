package repository

import (
	"context"
	"github.com/nicitapa/firstProgect/internal/models"
)

func (r *Repository) GetAllUsers(ctx context.Context) (users []models.User, err error) {
	if err = r.db.SelectContext(ctx, &users, `
		SELECT id, name, email, age, 
		FROM users
		ORDER BY id`); err != nil {
		return nil, r.translateError(err)
	}

	return users, nil
}

func (r *Repository) GetUsersByID(ctx context.Context, id int) (user models.User, err error) {
	if err = r.db.GetContext(ctx, &user, `
		SELECT id, name, email, age
		FROM users
		WHERE id = $1`, id); err != nil {
		return models.User{}, r.translateError(err)
	}

	return user, nil
}

func (r *Repository) CreateUsers(ctx context.Context, users models.User) (err error) {
	_, err = r.db.ExecContext(ctx, `INSERT INTO users (name, email, age)
					VALUES ($1, $2, $3,)`,
		users.Name,
		users.Email,
		users.Age)
	if err != nil {
		return r.translateError(err)
	}

	return nil
}

func (r *Repository) UpdateUsersByID(ctx context.Context, users models.User) (err error) {
	_, err = r.db.ExecContext(ctx, `
		UPDATE users SET name = $1, 
		                    email = $2, 
		                    age = $3,
		                    		                WHERE id = $5`,
		users.Name,
		users.Email,
		users.Age,
		users.ID)
	if err != nil {
		return r.translateError(err)
	}

	return nil
}

func (r *Repository) DeleteUsersByID(ctx context.Context, id int) (err error) {
	_, err = r.db.ExecContext(ctx, `DELETE FROM users WHERE id = $1`, id)
	if err != nil {
		return r.translateError(err)
	}

	return nil
}
