package service

import (
	"context"
	"github.com/hendrihmwn/go-gorm-clean-architecture/app/domain/entity"
	"github.com/hendrihmwn/go-gorm-clean-architecture/app/domain/repository"
)

type ListBookParams struct {
}

type ListBookResponse struct {
	Data []entity.Book
}

type ListBookService interface {
	Call(ctx context.Context, params *ListBookParams) (*ListBookResponse, error)
}

type BookServiceImpl struct {
	bookRepository repository.BookRepository
}

func (s *BookServiceImpl) Call(ctx context.Context, params *ListBookParams) (*ListBookResponse, error) {
	list, err := s.bookRepository.List(ctx)
	if err != nil {
		return nil, err
	}

	return &ListBookResponse{
		Data: list,
	}, nil
}

func NewBookService(bookRepository repository.BookRepository) ListBookService {
	return &BookServiceImpl{
		bookRepository: bookRepository,
	}
}
