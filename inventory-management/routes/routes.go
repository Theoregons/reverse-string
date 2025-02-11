package routes

import (
	"inventory-management/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, produkController *controllers.ProdukController, inventarisController *controllers.InventarisController, pesananController *controllers.PesananController) {
	r.POST("/products", produkController.AddProduct)
	r.GET("/products", produkController.GetProducts)
	r.GET("/products/:id", produkController.GetProductByID)
	r.PUT("/products/:id", produkController.UpdateProduct)
	r.DELETE("/products/:id", produkController.DeleteProduct)

	r.GET("/inventory/:id", inventarisController.GetInventory)
	r.PUT("/inventory/:id", inventarisController.UpdateInventory)

	r.POST("/orders", pesananController.CreateOrder)
	r.GET("/orders", pesananController.GetOrders)
	r.GET("/orders/:id", pesananController.GetOrderByID)
	r.DELETE("/orders/:id", pesananController.DeleteOrder)
}
