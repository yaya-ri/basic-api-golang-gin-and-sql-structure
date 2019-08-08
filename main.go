package main

import (
	"golang_API_yaya/config"
	"golang_API_yaya/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}

	router := gin.Default()

	// router.GET("/category", inDB.GetCategoryList)
	// router.POST("/category", inDB.InsertCategory)
	// router.PUT("/category/:id", inDB.UpdateCategory)
	// router.DELETE("category/:id", inDB.DeleteCategory)

	// router.POST("/product", inDB.InsertProduct)
	// router.GET("/product", inDB.GetProductList)
	// router.PUT("/product/:id", inDB.UpdateProduct)
	// router.DELETE("product/:id", inDB.DeleteProduct)

	// router.POST("/category_product", inDB.InsertCategoryProduct)
	// router.PUT("/category_product/:id", inDB.UpdateCategoryProduct)
	// router.DELETE("category_product/:id", inDB.DeleteCategoryProduct)

	// router.POST("/image", inDB.UploadImage)
	// router.GET("image", inDB.GetImageList)
	// router.DELETE("image/:id", inDB.DeleteImage)
	// router.PUT("image/:id", inDB.EnableImage)

	// router.POST("/product_image", inDB.InsertProductImage)
	// router.DELETE("product_image/:id", inDB.DeleteProductImage)

	// router.GET("/getAll", inDB.GetAll)

	router.Run(":3000")
}
