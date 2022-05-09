package main

import (
	"io"
	"net/http"
	"os"
	"tugaaswp/Database"
	"tugaaswp/Models"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	database.Connect()	
	e.Static("/products", "products")
	e.POST("/login", func(c echo.Context) error {
		var user models.User
		err := database.DB.Where("email = ? AND password = ?", c.FormValue("email"), c.FormValue("password")).First(&user).Error
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, echo.Map{
			"message" : "Login Success",
			"data" : user,
		})
	})

	e.POST("/register", func(c echo.Context) error {
		var user models.User

		if err := c.Bind(&user); err != nil {
			return err
		}
		database.DB.Create(&user)
		return c.JSON(http.StatusOK, echo.Map{
			"message" : "Successfully Register your account",
		})


	})

	e.GET("/products", func(c echo.Context) error {
		var product []models.Product

		err := database.DB.Find(&product).Error

		if err != nil {
			return c.JSON(500, err)
		}
		return c.JSON(http.StatusOK, echo.Map{
			"data": product,
		})
	})
	e.GET("/products/:id", func(c echo.Context) error {
		var product models.Product
		err := database.DB.Where("id = ?", c.Param("id")).First(&product).Error

		if err !=nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"error": "Product not found",
			})
		}
		return c.JSON(http.StatusOK, echo.Map{
			"data": product,
		})
	})
	e.POST("/create-products", func(c echo.Context) error {
		file, err := c.FormFile("image")
		if err != nil {
			return err
		}
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()
	
		// Destination
		dst, err := os.Create("products/" + file.Filename)
	
		if err != nil {
			return err
		}
		defer dst.Close()
	
		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}
	
		 product := models.Product{
			 Image: file.Filename,
		 }

		if err := c.Bind(&product); err != nil {
			return err
		}
		errinput := database.DB.Create(&product).Error

		if errinput != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"error": "Product cannot create",
			})
		}
		return c.JSON(http.StatusCreated, echo.Map{
			"message" : "Product created",
		})
	})
	

}