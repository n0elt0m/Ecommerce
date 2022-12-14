package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/noel/ecommerce/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controllers.Signup())
	incomingRoutes.POST("/users/login", controllers.Login())
	incomingRoutes.POST("/admin/login", controllers.AdminLogin())
	incomingRoutes.POST("/admin/addproduct", controllers.ProductViewerAdmin())
	incomingRoutes.DELETE("/admin/deleteuser", controllers.AdminDeleteUser())
	incomingRoutes.GET("/users/productview", controllers.SearchProduct())
	incomingRoutes.GET("/users/search", controllers.SearchProductByQuery())
	incomingRoutes.GET("/users/productminrate", controllers.MinRateFilter())
	incomingRoutes.GET("/users/productmaxrate", controllers.MaxRateFilter())
	incomingRoutes.GET("/users/productrating", controllers.RatingFilter())
	incomingRoutes.GET("/users/productsortdescending", controllers.DescendingSort())
	incomingRoutes.GET("/users/productsortascending", controllers.AscendingSort())

}
