package server

import (
	"github.com/gin-gonic/gin"
	"github.com/hendrihmwn/go-gorm-clean-architecture/app/domain/service"
	"net/http"
)

func (s *BookServer) ListBookHandler(c *gin.Context) {
	result, err := s.listBookService.Call(c, &service.ListBookParams{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, result)
}
