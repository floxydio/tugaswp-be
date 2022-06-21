package main

import (
	"io"
	"net/http"
	"os"
	database "tugaswp/Database"
	models "tugaswp/Models"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	database.Connect()
	e.Static("/products", "products")
	e.POST("/login", func(c echo.Context) error {
		var user models.User
		err := database.DB.Where("email = ? AND password = ?", c.FormValue("email"), c.FormValue("password")).First(&user).Error
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Login Success",
			"data":    user,
		})
	})

	e.POST("/register", func(c echo.Context) error {
		var user models.User

		if err := c.Bind(&user); err != nil {
			return err
		}
		database.DB.Create(&user)
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Successfully Register your account",
		})

	})

	e.GET("/products", func(c echo.Context) error {
		var product []models.Product
		search := c.QueryParam("search")
		var err error

		if search != "" {
			err = database.DB.Where("nama_produk LIKE ?", "%"+search+"%").Find(&product).Error
		} else {
			err = database.DB.Find(&product).Error
		}

		if err != nil {
			return c.JSON(500, echo.Map{
				"message": err.Error(),
			})
		}
		return c.JSON(http.StatusOK, echo.Map{
			"data": product,
		})
	})
	e.GET("/product-history/:id", func(c echo.Context) error {
		var productHistory []models.DBGetProduct
		var err error
		id := c.Param("id")
		err = database.DB.Raw("SELECT product_history.id, user.nama, product.nama_produk, product.harga FROM product_history LEFT JOIN user ON product_history.user_id = user.id LEFT JOIN product ON product_history.product_id = product.id WHERE user.id = ?", id).Find(&productHistory).Error
		if err != nil {
			return c.JSON(500, echo.Map{
				"message": err.Error(),
			})
		}
		return c.JSON(http.StatusOK, echo.Map{
			"data": productHistory,
		})
	})

	e.POST("/product-history", func(c echo.Context) error {
		var productHistory models.ProductHistory

		if err := c.Bind(&productHistory); err != nil {
			return err
		}

		err := database.DB.Create(&productHistory).Error
		if err != nil {
			return c.JSON(500, echo.Map{
				"message": err.Error(),
			})
		}
		return c.JSON(http.StatusOK, echo.Map{
			"data": productHistory,
		})
	})

	e.POST("/create-products", func(c echo.Context) error {
		file, err := c.FormFile("image")
		if err != nil {
			return err
		}
		if file != nil {

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
			"message": "Product created",
		})
	})
	e.Start(":2000")

}
