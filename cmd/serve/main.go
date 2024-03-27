package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/hendrihmwn/go-gorm-clean-architecture/app/config"
	"github.com/hendrihmwn/go-gorm-clean-architecture/app/domain/repository"
	"github.com/hendrihmwn/go-gorm-clean-architecture/app/domain/service"
	"github.com/hendrihmwn/go-gorm-clean-architecture/app/server"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
)

var r *gin.Engine

func init() {
	db, err := gorm.Open(postgres.Open(config.GetDBConnectionString()), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to initialize gorm: %v", err)
	}

	r = gin.Default()

	// initiate repository
	bookRepository := repository.NewBookRepository(db)

	// Initiate Service
	listBookService := service.NewBookService(bookRepository)

	// Initiate Server
	bookServer := server.NewBookServer(
		server.WithListBookService(listBookService),
	)

	// Define your handler
	r.GET("/books", bookServer.ListBookHandler)
}

func run() error {
	// NOTE: Setup context, so the requets can be cancelled
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	srv := &http.Server{
		Addr:    ":8000",
		Handler: r,
	}

	return srv.ListenAndServe()
}

func main() {

	if err := run(); err != nil {
		log.Fatal(err)
	}
}
