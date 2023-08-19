package controllers

import (
	"net/http"
	"strconv"

	"github.com/anasmisbah/backend-inventory/config"
	"github.com/anasmisbah/backend-inventory/models"
	"github.com/labstack/echo/v4"
)

func AddProductController(c echo.Context) error {
	var requestProduct models.Product
	c.Bind(&requestProduct)
	if err := c.Validate(requestProduct); err != nil {
		return err
	}
	result := config.DB.Create(&requestProduct)	
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError,models.BaseResponse{
			Status: false,
			Message: "Failed insert data product",
			Data: nil,
		})
	}
	return c.JSON(http.StatusOK,models.BaseResponse{
		Status: true,
		Message: "Success insert data product",
		Data: requestProduct,
	})
}

func GetDetailProductController(c echo.Context) error {
	id,_ := strconv.Atoi(c.Param("id"))
	var product models.Product = models.Product{Id: uint(id)}
	result := config.DB.Model(&models.Product{}).Preload("Stock").First(&product)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound,models.BaseResponse{
			Status: false,
			Message: "Failed get data product",
			Data: nil,
		})
	}
	return c.JSON(http.StatusOK, models.BaseResponse{
		Status: true,
		Message: "success get data product",
		Data: product,
	})
}

func DeleteProductController(c echo.Context) error {
	id,_ := strconv.Atoi(c.Param("id"))
	var product models.Product = models.Product{Id: uint(id)}
	result := config.DB.Delete(&product)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound,models.BaseResponse{
			Status: false,
			Message: "Failed delete data product",
			Data: nil,
		})
	}
	return c.JSON(http.StatusOK, models.BaseResponse{
		Status: true,
		Message: "success delete data product",
		Data: nil,
	})
}
func GetProductsController(c echo.Context) error {
	var products []models.Product
	result := config.DB.Model(&models.Product{}).Preload("Stock").Find(&products)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound,models.BaseResponse{
			Status: false,
			Message: "Failed get data products",
			Data: nil,
		})
	}
	return c.JSON(http.StatusOK, models.BaseResponse{
		Status: true,
		Message: "success get data products",
		Data: products,
	})
}

func UpdateProductController(c echo.Context) error {
	var product models.Product
	result := config.DB.First(&product)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound,models.BaseResponse{
			Status: false,
			Message: "Failed get data product",
			Data: nil,
		})
	}
	var requestUpdatedProduct models.Product
	c.Bind(&requestUpdatedProduct)
	if err := c.Validate(requestUpdatedProduct); err != nil {
		return err
	}
	result = config.DB.Model(&product).Updates(&requestUpdatedProduct)	
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError,models.BaseResponse{
			Status: false,
			Message: "Failed update data product",
			Data: nil,
		})
	}
	return c.JSON(http.StatusOK,models.BaseResponse{
		Status: true,
		Message: "Success update data product",
		Data: product,
	})
}