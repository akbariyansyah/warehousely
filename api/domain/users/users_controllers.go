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
	user := new(User)
	if err := c.Bind(user); err != nil {
		c.JSON(400, err.Error())
		return
	}

	result, err := u.UserUsecase.HandleUserLogin(user)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	c.JSON(201, result)
}

func (u *UserController) HandleUserRegister(c *gin.Context) {
	user := new(User)
	if err := c.Bind(user); err != nil {
		c.JSON(400, err.Error())
		return
	}

	result, err := u.UserUsecase.HandleUserRegister(user)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	c.JSON(201, result)
}
