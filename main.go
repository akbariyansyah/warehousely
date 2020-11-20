package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)
const (
	ServerHost = "localhost"
	ServerPort = "8000"
)
type Goods struct {
	ID int
	Name string
	Type string
	Stock int

}
func main() {
	g := gin.Default()

	g.GET("/", func(c *gin.Context) {
		c.JSON(200, "Hello There")
	})
	log.Printf("Starting server at %s:%s",ServerHost,ServerPort)
	source := fmt.Sprintf("%s:%s",ServerHost,ServerPort)
	g.Run(source)
}
