package users

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
)

type UserController struct {
	UserUsecase UserUsecaseInterface
}

func NewUserController(db *pg.DB) *UserController {
	return &UserController{NewUserUsecase(db)}
}

func (u *UserController) HandleUserLogin(c *gin.Context) {
	c.JSON(200, "Login There")
}

func (u *UserController) HandleUserRegister(c *gin.Context) {
	c.JSON(200, "Register There")
}
