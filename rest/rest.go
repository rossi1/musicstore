package rest

import (
	"os"

	"github.com/gin-gonic/gin"
)

func RunAPI(address string) error {
	h, err := NewHandler(os.Getenv("dbURL"))
	if err != nil {
		return err
	}
	return RunAPIWithHandler(address, h)
}

func RunAPIWithHandler(address string, h HandlerInterface) error {
	r := gin.Default()

	//get products
	r.GET("/products", h.GetProducts)
	//get promos
	r.GET("/promos", h.GetPromos)
	//post user sign in
	//r.POST("/users/signin", h.SignIn)
	//add a user
	//r.POST("/users", h.AddUser)
	//post user sign out
	//r.POST("/user/:id/signout", h.SignOut)
	//get user orders
	//r.GET("/user/:id/orders", h.GetOrders)
	//post purchase charge
	//r.POST("/users/charge", h.Charge)
	//run the server

	userGroup := r.Group("/user")
	{
		userGroup.POST("/:id/signout", h.SignOut)
		userGroup.GET("/:id/orders", h.GetOrders)
	}

	usersGroup := r.Group("/users")
	{
		usersGroup.POST("/charge", h.Charge)
		usersGroup.POST("/signin", h.SignIn)
		usersGroup.POST("", h.AddUser)
	}

	return r.Run(address)
}
