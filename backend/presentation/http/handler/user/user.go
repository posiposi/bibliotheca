package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/posiposi/project/backend/usecase/user"
)

type UserHandler interface {
	GetUsers() gin.HandlerFunc
}

type userHandler struct {
	UserUsecase user.UserUsecase
}

func NewUserHandler(userUsecase user.UserUsecase) UserHandler {
	return &userHandler{
		UserUsecase: userUsecase,
	}
}

func (uh *userHandler) GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := uh.UserUsecase.GetList()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"users": users})
	}
}
