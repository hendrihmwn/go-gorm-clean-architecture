package repository

import (
	"context"
	"github.com/hendrihmwn/go-gorm-clean-architecture/app/domain/entity"
	"gorm.io/gorm"
)

type BookRepository interface {
	List(ctx context.Context) ([]entity.Book, error)
}

type BookRepositoryImpl struct {
	db *gorm.DB
}

func (r *BookRepositoryImpl) List(ctx context.Context) ([]entity.Book, error) {
	var book []entity.Book
	err := r.db.WithContext(ctx).Find(&book).Error
	if err != nil {
		return nil, err
	}
	return book, nil
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &BookRepositoryImpl{
		db: db,
	}
}
