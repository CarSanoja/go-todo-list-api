package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(CORSMiddleware())
	r.Use(LoggingMiddleware())

	r.LoadHTMLGlob("handlers/views/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", nil)
	})

	r.GET("/todos", GetToDos)
	r.GET("/todos/:id", GetToDoByID)
	r.POST("/todos", CreateToDo)
	r.PUT("/todos/:id", UpdateToDo)
	r.DELETE("/todos/:id", DeleteToDo)

	return r
}
