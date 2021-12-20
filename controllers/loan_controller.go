package controllers

import (
	"context"
	"gocrud/controllers/dto"
	"gocrud/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoanControllerInterface interface {
	Create(c *gin.Context)
	ReadById(c *gin.Context)
	UpdateById(c *gin.Context)
	DeleteById(c *gin.Context)
}

type LoanController struct {
	loanService services.LoanServiceInterface
}

func newLoanController(loanService services.LoanServiceInterface) LoanControllerInterface {
	return &LoanController{
		loanService,
	}
}

func (ctr *LoanController) Create(c *gin.Context) {
	u := dto.CreateLoanRequest{}
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := ctr.userService.Create(context.Background(), u); err != nul {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
}
