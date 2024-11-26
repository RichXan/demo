package controller

import (
	"net/http"
	"x-project/internal/access/entity"
	"x-project/internal/access/repository"
	"x-project/internal/http/dto"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AccountController struct {
	accountRepo repository.AccountRepository
}

func NewAccountController(db *gorm.DB) *AccountController {
	accountRepo := repository.NewAccountRepo(db)
	return &AccountController{
		accountRepo: accountRepo,
	}
}

func (c *AccountController) Create(ctx *gin.Context) {
	var dto dto.CreateAccountDto
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	account := &entity.Account{
		Username: dto.Username,
		Password: dto.Password,
	}
	err := c.accountRepo.Create(account)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": account})
}
