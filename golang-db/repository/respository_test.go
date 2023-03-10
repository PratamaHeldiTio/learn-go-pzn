package repository

import (
	"context"
	"fmt"
	_ "github.com/lib/pq"
	golang_db "golang-db"
	"golang-db/entity"
	"testing"
)

func TestInsertComment(t *testing.T) {
	commentRepository := NewCommentRepository(golang_db.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Name:    "heldi",
		Comment: "loreemm",
	}
	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindIdComment(t *testing.T) {
	commentRepository := NewCommentRepository(golang_db.GetConnection())

	ctx := context.Background()
	result, err := commentRepository.FindById(ctx, 11)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindAllComment(t *testing.T) {
	commentRepository := NewCommentRepository(golang_db.GetConnection())

	ctx := context.Background()
	result, err := commentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	for _, comment := range result {
		fmt.Println(comment)
	}
}

