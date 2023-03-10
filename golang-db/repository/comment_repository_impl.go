package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-db/entity"
	"strconv"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{DB: db}
}

func (repository *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	query := "INSERT INTO comments(name, comment) VALUES ($1, $2) RETURNING id"
	rows := repository.DB.QueryRowContext(ctx, query, comment.Name, comment.Comment)
	err := rows.Scan(&comment.Id)

	if err != nil {
		return comment, err
	}

	return comment, nil
}

func (repository *commentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	query := "SELECT id, name, comment FROM comments WHERE id = $1"
	rows, err := repository.DB.QueryContext(ctx, query, id)
	comment := entity.Comment{}
	if err != nil {
		return comment, err
	}
	defer rows.Close()

	if rows.Next() {
		rows.Scan(&comment.Id, &comment.Name, &comment.Comment)
		return comment, nil
	} else {
		err := errors.New("Commnent dengan id " + strconv.Itoa(int(id)) + " not found")
		return comment, err
	}
}

func (repository *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	query := "SELECT id, name, comment FROM comments"
	rows, err := repository.DB.QueryContext(ctx, query)
	var comments []entity.Comment
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var comment entity.Comment
		rows.Scan(&comment.Id, &comment.Name, &comment.Comment)
		comments = append(comments, comment)
	}
	return comments, nil
}
