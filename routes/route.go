package routes

import (
	"net/http"

	"github.com/anasmisbah/backend-inventory/controllers"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoute(e *echo.Echo)  {
	e.Validator = &CustomValidator{validator: validator.New()}
	e.Use(middleware.Logger())
	e.POST("/products",controllers.AddProductController)
	e.GET("/products/:id",controllers.GetDetailProductController)
	e.GET("/products",controllers.GetProductsController)
	e.PUT("/products/:id",controllers.UpdateProductController)
	e.DELETE("/products/:id",controllers.DeleteProductController)

	e.POST("/stocks",controllers.AddStockController)
	e.GET("/stocks/:id",controllers.GetDetailStockController)
	e.GET("/stocks",controllers.GetStocksController)
	e.PUT("/stocks/:id",controllers.UpdateStockController)
	e.DELETE("/stocks/:id",controllers.DeleteStockController)
}

type CustomValidator struct {
    validator *validator.Validate
  }
func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
	  // Optionally, you could return the error to give each route more control over the status code
	  return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
  }