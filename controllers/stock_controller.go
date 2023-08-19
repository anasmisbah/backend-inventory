package controllers

import (
	"net/http"
	"strconv"

	"github.com/anasmisbah/backend-inventory/config"
	"github.com/anasmisbah/backend-inventory/models"
	"github.com/labstack/echo/v4"
)

func AddStockController(c echo.Context) error {
	var requestStock models.Stock
	c.Bind(&requestStock)
	if err := c.Validate(requestStock); err != nil {
		return err
	}
	result := config.DB.Create(&requestStock)	
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError,models.BaseResponse{
			Status: false,
			Message: "Failed insert data stock",
			Data: nil,
		})
	}
	return c.JSON(http.StatusOK,models.BaseResponse{
		Status: true,
		Message: "Success insert data stock",
		Data: requestStock,
	})
}

func GetDetailStockController(c echo.Context) error {
	id,_ := strconv.Atoi(c.Param("id"))
	var product models.Stock = models.Stock{Id: uint(id)}
	result := config.DB.First(&product)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound,models.BaseResponse{
			Status: false,
			Message: "Failed get data stock",
			Data: nil,
		})
	}
	return c.JSON(http.StatusOK, models.BaseResponse{
		Status: true,
		Message: "success get data stock",
		Data: product,
	})
}

func DeleteStockController(c echo.Context) error {
	id,_ := strconv.Atoi(c.Param("id"))
	var product models.Stock = models.Stock{Id: uint(id)}
	result := config.DB.Delete(&product)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound,models.BaseResponse{
			Status: false,
			Message: "Failed delete data stock",
			Data: nil,
		})
	}
	return c.JSON(http.StatusOK, models.BaseResponse{
		Status: true,
		Message: "success delete data stock",
		Data: nil,
	})
}
func GetStocksController(c echo.Context) error {
	var products []models.Stock
	result := config.DB.Find(&products)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound,models.BaseResponse{
			Status: false,
			Message: "Failed get data stocks",
			Data: nil,
		})
	}
	return c.JSON(http.StatusOK, models.BaseResponse{
		Status: true,
		Message: "success get data stocks",
		Data: products,
	})
}

func UpdateStockController(c echo.Context) error {
	var product models.Stock
	result := config.DB.First(&product)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound,models.BaseResponse{
			Status: false,
			Message: "Failed get data stock",
			Data: nil,
		})
	}
	var requestUpdatedProduct models.Stock
	c.Bind(&requestUpdatedProduct)
	if err := c.Validate(requestUpdatedProduct); err != nil {
		return err
	}
	result = config.DB.Model(&product).Updates(&requestUpdatedProduct)	
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError,models.BaseResponse{
			Status: false,
			Message: "Failed update data stock",
			Data: nil,
		})
	}
	return c.JSON(http.StatusOK,models.BaseResponse{
		Status: true,
		Message: "Success update data stock",
		Data: product,
	})
}