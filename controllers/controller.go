package controllers

//basic function golang example

// import (
// 	"fmt"
// 	"golang_API_yaya/common"
// 	"golang_API_yaya/model"
// 	"io"
// 	"log"
// 	"net/http"
// 	"os"
// 	"strconv"
// 	"time"

// 	"github.com/gin-gonic/gin"
// )

// func (idb *InDB) InsertCategory(c *gin.Context) {
// 	var category model.Category
// 	var res model.Response
// 	now := time.Now()

// 	category_name := c.PostForm("name")
// 	category_enable := true

// 	if category_name == "" {
// 		res.Success = false
// 		res.Message = "request not valid"
// 		c.JSON(http.StatusBadRequest, res)
// 		return
// 	}

// 	category.Name = category_name
// 	category.Enable = &category_enable
// 	category.CreatedAt = now
// 	category.UpdatedAt = now

// 	err := idb.DB.Create(&category)

// 	if err.Error != nil {
// 		common.Error(err.Error, "Failed to insert category")
// 		res.Success = false
// 		res.Message = "Failed to insert category"
// 		c.JSON(http.StatusBadRequest, res)
// 		return
// 	}

// 	res.Success = true
// 	res.Message = "Success to insert category"
// 	res.Data = category
// 	c.JSON(http.StatusOK, res)
// }

// func (idb *InDB) GetCategoryList(c *gin.Context) {
// 	var categories []*model.Category
// 	var res model.Response

// 	enableParam := c.Query("enable")

// 	enable := true
// 	if enableParam == "true" {
// 		enable = true
// 	} else if enableParam == "false" {
// 		enable = false
// 	}

// 	err := idb.DB.Where("enable = ? AND deleted_at IS NULL", enable).Find(&categories)

// 	if err.Error != nil {
// 		common.Error(err.Error, "Failed get category list")
// 		res.Success = false
// 		res.Message = "Failed get category list"
// 		res.Data = categories
// 		c.JSON(http.StatusBadRequest, res)
// 		return
// 	}

// 	if len(categories) == 0 {
// 		res.Success = true
// 		res.Message = "category is empty"
// 		c.JSON(http.StatusOK, res)
// 		return
// 	} else {
// 		res.Success = true
// 		res.Message = "Success get category list"
// 		res.Data = categories
// 		c.JSON(http.StatusOK, res)
// 	}
// }

// func (idb *InDB) UpdateCategory(c *gin.Context) {
// 	var res model.Response

// 	id := c.Param("id")
// 	name := c.PostForm("name")
// 	enable := c.PostForm("enable")
// 	e := true
// 	if enable == "true" {
// 		e = true
// 	} else if enable == "false" {
// 		e = false
// 	}

// 	if enable == "" {
// 		res.Success = false
// 		res.Message = "request not valid"
// 		c.JSON(http.StatusBadRequest, res)
// 		return
// 	}

// 	var category model.Category
// 	var newCategory model.Category

// 	err := idb.DB.First(&category, id)

// 	if err.Error != nil {
// 		res.Success = false
// 		res.Message = "can't get data from database"
// 		c.JSON(http.StatusBadGateway, res)
// 		return
// 	}
// 	//newProduct.ID = product.ID
// 	if name != "" {
// 		newCategory.Name = name
// 	}
// 	newCategory.Name = name
// 	newCategory.Enable = &e
// 	newCategory.UpdatedAt = time.Now()

// 	err = idb.DB.Model(&category).Update(newCategory)

// 	if err.Error != nil {
// 		res.Success = false
// 		res.Message = "can't update data to database"
// 		c.JSON(http.StatusBadGateway, res)
// 		return
// 	}

// 	res.Success = true
// 	res.Message = "Success update data to database"
// 	res.Data = err.Value
// 	c.JSON(http.StatusOK, res)

// }

// func (idb *InDB) DeleteCategory(c *gin.Context) {
// 	var res model.Response

// 	id := c.Param("id")

// 	now := time.Now()

// 	var category model.Category
// 	var newCategory model.Category

// 	err := idb.DB.First(&category, id)

// 	if err.Error != nil {
// 		res.Success = false
// 		res.Message = "data doesn't exist"
// 		c.JSON(http.StatusBadGateway, res)
// 		return
// 	}

// 	newCategory.DeletedAt = &now

// 	err = idb.DB.Model(&category).Update(newCategory)

// 	if err.Error != nil {
// 		res.Success = false
// 		res.Message = "can't delete data from database"
// 		c.JSON(http.StatusBadGateway, res)
// 		return
// 	}
// 	err = idb.DB.Exec("DELETE FROM category_products WHERE category_id=?", id)

// 	res.Success = true
// 	res.Message = "Success delete data from database"
// 	//res.Data = err.Value
// 	c.JSON(http.StatusOK, res)

// }

// func (idb *InDB) InsertProduct(c *gin.Context) {
// 	var product model.ProductModel
// 	var res model.Response
// 	now := time.Now()

// 	product_name := c.PostForm("name")
// 	product_description := c.PostForm("description")
// 	product_enable := true

// 	if product_name == "" || product_description == "" {
// 		res.Success = false
// 		res.Message = "request not valid"
// 		c.JSON(http.StatusBadRequest, res)
// 		return
// 	}

// 	product.Name = product_name
// 	product.Enable = &product_enable
// 	product.Description = product_description
// 	product.CreatedAt = now
// 	product.UpdatedAt = now

// 	err := idb.DB.Table("products").Create(&product)

// 	if err.Error != nil {
// 		common.Error(err.Error, "Failed to insert product")
// 		res.Success = false
// 		res.Message = "Failed to insert product"
// 		c.JSON(http.StatusBadRequest, res)
// 		return
// 	}

// 	res.Success = true
// 	res.Message = "Success to insert product"
// 	res.Data = product
// 	c.JSON(http.StatusOK, res)
// }

// func (idb *InDB) GetProductList(c *gin.Context) {
// 	var products []*model.ProductModel
// 	var res model.Response

// 	enableParam := c.Query("enable")

// 	enable := true
// 	if enableParam == "true" {
// 		enable = true
// 	} else if enableParam == "false" {
// 		enable = false
// 	}

// 	//idb.DB.SingularTable(true)

// 	err := idb.DB.Table("products").Where("enable = ?", enable).Find(&products)

// 	if err.Error != nil {
// 		common.Error(err.Error, "Failed get product list")
// 		res.Success = false
// 		res.Message = "Failed get product list"
// 		res.Data = products
// 		c.JSON(http.StatusBadRequest, res)
// 		return
// 	}

// 	if len(products) == 0 {
// 		res.Success = true
// 		res.Message = "product is empty"
// 		c.JSON(http.StatusOK, res)
// 		return
// 	} else {
// 		res.Success = true
// 		res.Message = "Success get product list"
// 		res.Data = products
// 		c.JSON(http.StatusOK, res)
// 	}
// }

// func (idb *InDB) UpdateProduct(c *gin.Context) {
// 	var res model.Response

// 	id := c.Param("id")
// 	name := c.PostForm("name")
// 	description := c.PostForm("description")
// 	enable := c.PostForm("enable")

// 	// if name == "" || description == "" || enable == "" {
// 	// 	res.Success = false
// 	// 	res.Message = "request not valid"
// 	// 	c.JSON(http.StatusBadRequest, res)
// 	// 	return
// 	// }
// 	e := true
// 	if enable == "true" {
// 		e = true
// 	} else if enable == "false" {
// 		e = false
// 	}

// 	var product model.Product
// 	var newProduct model.Product

// 	err := idb.DB.First(&product, id)

// 	if err.Error != nil {
// 		res.Success = false
// 		res.Message = "can't get data from database"
// 		c.JSON(http.StatusBadGateway, res)
// 		return
// 	}
// 	//newProduct.ID = product.ID
// 	if name != "" {
// 		newProduct.Name = name
// 	}
// 	if description != "" {
// 		newProduct.Description = description
// 	}
// 	newProduct.Enable = &e
// 	newProduct.UpdatedAt = time.Now()

// 	err = idb.DB.Model(&product).Update(newProduct)

// 	if err.Error != nil {
// 		res.Success = false
// 		res.Message = "can't update data to database"
// 		c.JSON(http.StatusBadGateway, res)
// 		return
// 	}

// 	res.Success = true
// 	res.Message = "Success update data to database"
// 	res.Data = err.Value
// 	c.JSON(http.StatusOK, res)

// }

// func (idb *InDB) DeleteProduct(c *gin.Context) {
// 	var res model.Response

// 	id := c.Param("id")
// 	now := time.Now()

// 	var product model.Product
// 	var newProduct model.Product

// 	err := idb.DB.First(&product, id)

// 	if err.Error != nil {
// 		res.Success = false
// 		res.Message = "data doesn't exist"
// 		c.JSON(http.StatusBadGateway, res)
// 		return
// 	}

// 	newProduct.DeletedAt = &now

// 	err = idb.DB.Model(&product).Update(newProduct)

// 	if err.Error != nil {
// 		res.Success = false
// 		res.Message = "can't delete data from database"
// 		c.JSON(http.StatusBadGateway, res)
// 		return
// 	}

// 	err = idb.DB.Exec("DELETE FROM category_products WHERE product_id=?", id)

// 	err = idb.DB.Exec("DELETE FROM product_images WHERE product_id=?", id)

// 	res.Success = true
// 	res.Message = "Success delete data from database"
// 	//res.Data = err.Value
// 	c.JSON(http.StatusOK, res)

// }

// func (idb *InDB) InsertCategoryProduct(c *gin.Context) {
// 	var categoryproduct model.CategoryProduct
// 	var result gin.H
// 	now := time.Now()

// 	product_id := c.PostForm("product_id")
// 	category_id := c.PostForm("category_id")
// 	p_id, _ := strconv.ParseUint(product_id, 10, 32)
// 	c_id, _ := strconv.ParseUint(category_id, 10, 32)

// 	categoryproduct.CategoryID = uint(c_id)
// 	categoryproduct.ProductID = uint(p_id)
// 	categoryproduct.CreatedAt = now
// 	categoryproduct.UpdatedAt = now

// 	if categoryproduct.CategoryID == 0 || categoryproduct.ProductID == 0 {
// 		result = gin.H{
// 			"success": false,
// 			"message": "Request not valid",
// 		}
// 		c.JSON(http.StatusBadRequest, result)
// 		return
// 	}

// 	err := idb.DB.Create(&categoryproduct)

// 	if err.Error != nil {
// 		common.Error(err.Error, "Failed to insert category product")
// 		result = gin.H{
// 			"success": false,
// 			"message": err.Error,
// 		}
// 		c.JSON(http.StatusBadRequest, result)
// 		return
// 	}

// 	result = gin.H{
// 		"success": true,
// 		"message": "Success insert category product",
// 		"data":    categoryproduct,
// 	}
// 	c.JSON(http.StatusOK, result)
// }

// func (idb *InDB) UpdateCategoryProduct(c *gin.Context) {
// 	var res model.Response

// 	id := c.Param("id")
// 	product_id := c.PostForm("product_id")
// 	category_id := c.PostForm("category_id")

// 	if product_id == "" || category_id == "" {
// 		res.Success = false
// 		res.Message = "request not valid"
// 		c.JSON(http.StatusBadRequest, res)
// 		return
// 	}

// 	p_id, _ := strconv.ParseUint(product_id, 10, 32)
// 	c_id, _ := strconv.ParseUint(category_id, 10, 32)

// 	var pc model.CategoryProduct
// 	var newPc model.CategoryProduct

// 	err := idb.DB.First(&pc, id)

// 	if err.Error != nil {
// 		res.Success = false
// 		res.Message = "can't get data from database"
// 		c.JSON(http.StatusBadGateway, res)
// 		return
// 	}
// 	//newProduct.ID = product.ID
// 	newPc.ProductID = uint(p_id)
// 	newPc.CategoryID = uint(c_id)
// 	newPc.UpdatedAt = time.Now()

// 	err = idb.DB.Model(&pc).Update(newPc)

// 	if err.Error != nil {
// 		res.Success = false
// 		res.Message = "can't update data to database"
// 		c.JSON(http.StatusBadGateway, res)
// 		return
// 	}

// 	res.Success = true
// 	res.Message = "Success update data to database"
// 	res.Data = err.Value
// 	c.JSON(http.StatusOK, res)

// }

// func (idb *InDB) DeleteCategoryProduct(c *gin.Context) {
// 	var res model.Response

// 	id := c.Param("id")

// 	var categoryProduct model.CategoryProduct
// 	err := idb.DB.First(&categoryProduct, id)

// 	if err.Error != nil {
// 		res.Success = false
// 		res.Message = "data doesn't exist"
// 		c.JSON(http.StatusBadGateway, res)
// 		return
// 	}

// 	err = idb.DB.Delete(&categoryProduct)

// 	if err.Error != nil {
// 		res.Success = false
// 		res.Message = "can't delete data from database"
// 		c.JSON(http.StatusBadGateway, res)
// 		return
// 	}

// 	res.Success = true
// 	res.Message = "Success delete data from database"
// 	//res.Data = err.Value
// 	c.JSON(http.StatusOK, res)

// }

// func (idb *InDB) UploadImage(c *gin.Context) {
// 	var res model.Response
// 	var image model.Image

// 	now := time.Now()

// 	file, header, err := c.Request.FormFile("file")
// 	filename := header.Filename
// 	url := "./images/" + filename + ".png"
// 	fmt.Println(file)
// 	out, err := os.Create(url)
// 	if err != nil {
// 		log.Fatal(err)
// 		return
// 	}

// 	defer out.Close()
// 	_, err = io.Copy(out, file)
// 	if err != nil {
// 		log.Fatal(err)
// 		return
// 	}

// 	enable := true

// 	image.Name = filename
// 	image.File = url
// 	image.Enable = &enable
// 	image.CreatedAt = now
// 	image.UpdatedAt = now

// 	save := idb.DB.Create(&image)

// 	if save.Error != nil {
// 		res.Success = false
// 		res.Message = "can't save data to database"
// 		c.JSON(http.StatusBadGateway, res)
// 		return
// 	}

// 	res.Success = true
// 	res.Message = "Success save data to database"
// 	res.Data = save.Value
// 	c.JSON(http.StatusOK, res)
// }

// func (idb *InDB) GetImageList(c *gin.Context) {
// 	var images []*model.Image
// 	var res model.Response

// 	enableParam := c.Query("enable")

// 	enable := true
// 	if enableParam == "true" {
// 		enable = true
// 	} else if enableParam == "false" {
// 		enable = false
// 	}

// 	err := idb.DB.Where("enable = ? AND deleted_at IS NULL", enable).Find(&images)

// 	if err.Error != nil {
// 		common.Error(err.Error, "Failed get image list")
// 		res.Success = false
// 		res.Message = "Failed get image list"
// 		c.JSON(http.StatusBadRequest, res)
// 		return
// 	}

// 	if len(images) == 0 {
// 		res.Success = true
// 		res.Message = "image is empty"
// 		c.JSON(http.StatusOK, res)
// 		return
// 	} else {
// 		res.Success = true
// 		res.Message = "Success get image list"
// 		res.Data = images
// 		c.JSON(http.StatusOK, res)
// 	}
// }

// func (idb *InDB) EnableImage(c *gin.Context) {
// 	var res model.Response

// 	id := c.Param("id")
// 	enable := c.PostForm("enable")
// 	e := true
// 	if enable == "true" {
// 		e = true
// 	} else if enable == "false" {
// 		e = false
// 	}

// 	if enable == "" {
// 		res.Success = false
// 		res.Message = "request not valid"
// 		c.JSON(http.StatusBadRequest, res)
// 		return
// 	}

// 	var image model.Image
// 	var newImage model.Image

// 	err := idb.DB.First(&image, id)

// 	if err.Error != nil {
// 		res.Success = false
// 		res.Message = "can't get data from database"
// 		c.JSON(http.StatusBadGateway, res)
// 		return
// 	}
// 	//newProduct.ID = product.ID

// 	newImage.Enable = &e
// 	newImage.UpdatedAt = time.Now()

// 	err = idb.DB.Model(&image).Update(newImage)

// 	if err.Error != nil {
// 		res.Success = false
// 		res.Message = "can't update data to database"
// 		c.JSON(http.StatusBadGateway, res)
// 		return
// 	}

// 	res.Success = true
// 	res.Message = "Success update data to database"
// 	res.Data = err.Value
// 	c.JSON(http.StatusOK, res)

// }

// func (idb *InDB) InsertProductImage(c *gin.Context) {
// 	var productImage model.ProductImage
// 	var res model.Response
// 	now := time.Now()

// 	product_id := c.PostForm("product_id")
// 	image_id := c.PostForm("image_id")
// 	p_id, _ := strconv.ParseUint(product_id, 10, 32)
// 	i_id, _ := strconv.ParseUint(image_id, 10, 32)

// 	productImage.ProductID = uint(p_id)
// 	productImage.ImageID = uint(i_id)
// 	productImage.CreatedAt = now
// 	productImage.UpdatedAt = now

// 	if productImage.ImageID == 0 || productImage.ProductID == 0 {
// 		res.Success = false
// 		res.Message = "request not valid"
// 		c.JSON(http.StatusBadRequest, res)
// 		return
// 	}

// 	err := idb.DB.Create(&productImage)

// 	if err.Error != nil {
// 		common.Error(err.Error, "Failed to insert product image")
// 		res.Success = false
// 		res.Message = "can't save to database"
// 		c.JSON(http.StatusBadRequest, res)
// 		return
// 	}

// 	res.Success = true
// 	res.Message = "Success save data to database"
// 	res.Data = err.Value
// 	c.JSON(http.StatusOK, res)
// }

// func (idb *InDB) DeleteProductImage(c *gin.Context) {
// 	var res model.Response

// 	id := c.Param("id")

// 	var productImage model.ProductImage

// 	err := idb.DB.First(&productImage, id)

// 	if err.Error != nil {
// 		res.Success = false
// 		res.Message = "data doesn't exist"
// 		c.JSON(http.StatusBadGateway, res)
// 		return
// 	}

// 	err = idb.DB.Delete(&productImage)

// 	if err.Error != nil {
// 		res.Success = false
// 		res.Message = "can't delete data from database"
// 		c.JSON(http.StatusBadGateway, res)
// 		return
// 	}

// 	res.Success = true
// 	res.Message = "Success delete data from database"
// 	//res.Data = err.Value
// 	c.JSON(http.StatusOK, res)

// }

// func (idb *InDB) DeleteImage(c *gin.Context) {
// 	var res model.Response

// 	id := c.Param("id")

// 	var image model.Image
// 	var newImage model.Image

// 	err := idb.DB.First(&image, id)

// 	if err.Error != nil {
// 		res.Success = false
// 		res.Message = "data doesn't exist"
// 		c.JSON(http.StatusBadGateway, res)
// 		return
// 	}
// 	now := time.Now()

// 	newImage.DeletedAt = &now

// 	err = idb.DB.Model(&image).Update(newImage)

// 	if err.Error != nil {
// 		res.Success = false
// 		res.Message = "can't delete data from database"
// 		c.JSON(http.StatusBadGateway, res)
// 		return
// 	}

// 	err = idb.DB.Exec("DELETE FROM product_images WHERE image_id=?", id)

// 	res.Success = true
// 	res.Message = "Success delete data from database"
// 	//res.Data = err.Value
// 	c.JSON(http.StatusOK, res)

// }

// func (idb *InDB) GetAll(c *gin.Context) {
// 	var res model.Response
// 	var all []model.Product

// 	err := idb.DB.Preload("Category", "enable=?", true).Preload("Image", "enable=?", true).Find(&all)

// 	if err.Error != nil {
// 		common.Error(err.Error, "Failed to get product category: ")
// 		res.Success = false
// 		res.Message = "can't get data from database"
// 		c.JSON(http.StatusBadRequest, res)
// 		return
// 	}

// 	res.Success = true
// 	res.Message = "Success get product with category and image from database"
// 	res.Data = all
// 	c.JSON(http.StatusOK, res)
// }
