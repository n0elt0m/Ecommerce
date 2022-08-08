package main

import (
	"github.com/gin-gonic/gin"
	"github.com/noel/ecommerce/controllers"
	"github.com/noel/ecommerce/database"
	"github.com/noel/ecommerce/middleware"
	"github.com/noel/ecommerce/routes"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))
	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/listcart", controllers.GetItemFromCart())
	router.POST("/addaddress", controllers.AddAddress())
	router.PUT("/edithomeaddress", controllers.EditHomeAddress())
	router.PUT("/editworkaddress", controllers.EditWrorkAddress())
	router.GET("/deleteaddresses", controllers.DeleteAddress())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))
}
