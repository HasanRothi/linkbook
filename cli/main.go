package main

import (
	"linkbook/cli/db"
	"linkbook/cli/middlewares"
	"linkbook/cli/route_handlers"

	"github.com/gin-gonic/gin"
)

// var DB string

func require() {
	// middlewares.Logger()
	db.Connect()
}
func main() {
	require()

	server := gin.New()
	server.Use(gin.Recovery())
	server.Use(middlewares.Recover)
	server.GET("/", route_handlers.Home)
	server.GET("/guest", route_handlers.GuestHome("ROtHi"))

	//link routes
	server.GET("/links", route_handlers.GetAllLink)
	server.GET("/link/:id", route_handlers.GetSingleLink)
	server.GET("/links/pending", route_handlers.GetPendingLinks)
	server.GET("/links/popular", route_handlers.GetPopularLinks)
	server.POST("/link", route_handlers.PostSingleLink)
	server.POST("/link/published", route_handlers.PublishedSingleLink)

	//user routes
	server.GET("/users", route_handlers.GetAllUser)
	server.GET("/user/:id", route_handlers.GetSingleUser)
	server.POST("/user", route_handlers.PostSingleUser)

	server.Run()
}
