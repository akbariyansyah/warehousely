package users

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
)

func InitUserRoutes(mainRoute string, g *gin.Engine, db *pg.DB) {
	r := g.Group(mainRoute)
	userController := NewUserController(db)

	//http://localhost:8080/users/login
	r.GET("/login", userController.HandleUserLogin)

	//http://localhost:8080/users/register
	r.GET("/register", userController.HandleUserRegister)
}
