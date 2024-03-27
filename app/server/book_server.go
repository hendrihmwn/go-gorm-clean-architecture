package server

import (
	"github.com/hendrihmwn/go-gorm-clean-architecture/app/domain/service"
)

type BookServer struct {
	listBookService service.ListBookService
}

type BookServerSetter func(*BookServer)

func NewBookServer(setters ...BookServerSetter) *BookServer {
	s := &BookServer{}

	for _, set := range setters {
		set(s)
	}

	return s
}

func WithListBookService(listBookService service.ListBookService) BookServerSetter {
	return func(as *BookServer) {
		as.listBookService = listBookService
	}
}
