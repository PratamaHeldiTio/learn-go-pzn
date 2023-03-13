package repository

import (
	"context"
	"database/sql"
	"errors"
	"go-restapi/helper"
	"go-restapi/model/domain"
)

type CategoryRepositoryImpl struct {
}

func (repo *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	query := "INSERT INTO category(name) VALUES($1) RETURNING id"
	stmt, err := tx.Prepare(query)
	helper.PanicError(err)
	defer stmt.Close()

	err = stmt.QueryRowContext(ctx, category.Name).Scan(&category.Id)
	helper.PanicError(err)

	return category
}

func (repo *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	query := "UPDATE category SET name = $1 WHERE id = $2"
	_, err := tx.ExecContext(ctx, query, category.Name, category.Id)
	helper.PanicError(err)

	return category
}

func (repo *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, categoryId int) {
	query := "DELETE category WHERE id = $1"
	_, err := tx.ExecContext(ctx, query, categoryId)
	helper.PanicError(err)
}

func (repo *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	query := "SELECT (id, name) FROM category WHERE id = $1"
	rows, err := tx.QueryContext(ctx, query, categoryId)
	helper.PanicError(err)

	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicError(err)
		return category, nil
	} else {
		return category, errors.New("category is not found")
	}

}

func (repo *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	query := "SELECT id, name FROM category"
	rows, err := tx.QueryContext(ctx, query)
	helper.PanicError(err)

	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicError(err)
		categories = append(categories, category)
	}

	return categories
}
