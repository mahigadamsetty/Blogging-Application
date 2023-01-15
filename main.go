package main

import (
	handler "blog1/handler"
	"blog1/middleware"

	"github.com/gin-gonic/gin"
)

func main() {

	g := gin.Default()

	protected := g.Group("/")
	protected.Use(middleware.JwtAuthMiddleware())

	// register the protected routes that needs authorization
	protected.POST("/post", handler.HandleAddPost)
	protected.DELETE("/post/:id", handler.HandleDeletePost)
	protected.PATCH("/post/:id", handler.HandleUpdatePost)

	public := g.Group("/")

	// register the public routes that doesn't need authorization
	public.GET("/post", handler.HandleGetAllPosts)
	public.GET("/post/:id", handler.HandleGetPostByID)
	public.POST("/login", handler.HandleLogin)

	if err := g.Run("localhost:8080"); err != nil {
		panic(err)
	}
}
