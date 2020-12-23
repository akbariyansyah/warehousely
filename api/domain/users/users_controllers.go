package users

import (
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUsecase UserUsecaseInterface
}

func NewUserController(usecase UserUsecaseInterface) *UserController {
	return &UserController{UserUsecase: usecase}
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

func (u *UserController) HandleDeleteUser(c *gin.Context) {
	id := c.Param("id")
	err := u.UserUsecase.HandleDeleteUser(id)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	c.JSON(201, "Delete User with ID "+id+" Success")
}
