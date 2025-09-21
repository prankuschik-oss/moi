package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/nicitapa/firstProgect/internal/errs"
	"github.com/nicitapa/firstProgect/internal/models"
	"time"
)
var (
	defaultTTL = time.Minute * 5)

func (s *Service) GetAllUsers() (users []models.User, err error) {
	ctx := context.Background()
	users, err = s.repository.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *Service) GetUsersByID(id int) (users models.User, err error) {
	ctx := context.Background()

	err = s.cache.Get(ctx, fmt.Sprintf("product_%d", id), &users)
	if err == nil {

		return users, nil
	}

	users, err = s.repository.GetUsersByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.User{}, errs.ErrUsersNotfound
		}
		return models.User{}, err
	}

	if err = s.cache.Set(ctx, fmt.Sprintf("product_%d", users.ID), users, defaultTTL); err != nil {
		fmt.Printf("error during cache set: %v\n", err.Error())
}

func (s *Service) CreateUsers(users models.User) (err error) {
		ctx := context.Background()
		if len(users.ProductName) < 4 {
			return errs.ErrInvalidsersName
		}

		err = s.repository.CreateUsers(ctx, users)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) UpdateUsersByID(users models.User) (err error) {
	ctx := context.Background()
	_, err = s.repository.GetUsersByID(ctx, users.ID)
	if err != nil {
		if errors.Is(err, errs.ErrNotfound) {
			return errs.ErrUsersNotfound
		}
		return err
	}
	err = s.repository.UpdateUsersByID(ctx, users)
	if err != nil {
			return err
		}

	return nil
}

func (s *Service) DeleteUsersByID(id int) (err error) {
		ctx := context.Background()
	_, err = s.repository.GetUsersByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errs.ErrUsersNotfound
		}
		return err
	}

	err = s.repository.DeleteUsersByID(id)
	if err != nil {
		return err
	}

	return nil
}
