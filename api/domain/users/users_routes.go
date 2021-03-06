package users

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
)

func InitUserRoutes(mainRoute string, g *gin.Engine, db *pg.DB) {
	r := g.Group(mainRoute)
	userController := NewUserController(db)

	//http://localhost:8080/users/login
	r.POST("/login", userController.HandleUserLogin)

	//http://localhost:8080/users/register
	r.POST("/register", userController.HandleUserRegister)

	r.DELETE("/:id", userController.HandleDeleteUser)
}
