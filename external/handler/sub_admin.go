// pkg_handlers/sub_admin_handlers.go
package pkg_handlers

import (
	"github.com/PICH-IO/admin-api/pkg/models"
	"github.com/PICH-IO/admin-api/pkg/services"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

// AddProduct handles adding a new product
func AddProduct(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req struct {
			Name        string  `json:"name" validate:"required"`
			Description string  `json:"description"`
			Price       float64 `json:"price" validate:"required,gt=0"`
			CategoryID  int     `json:"category_id" validate:"required"`
			Stock       int     `json:"stock" validate:"required,gte=0"`
		}

		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		// Create product
		product := &models.Product{
			Name:        req.Name,
			Description: req.Description,
			Price:       req.Price,
			CategoryID:  req.CategoryID,
			Stock:       req.Stock,
			Status:      "in_stock",
		}

		productID, err := services.CreateProduct(db, product)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to add product",
			})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message":    "Product added successfully",
			"product_id": productID,
		})
	}
}

// GetAllProducts retrieves all products
func GetAllProducts(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		products, err := services.GetAllProducts(db)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to retrieve products",
			})
		}

		return c.JSON(products)
	}
}

// UpdateProduct handles updating a product's information
func UpdateProduct(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		productID := c.Params("id")
		var req struct {
			Name        string  `json:"name"`
			Description string  `json:"description"`
			Price       float64 `json:"price" validate:"omitempty,gt=0"`
			CategoryID  int     `json:"category_id"`
			Stock       int     `json:"stock" validate:"omitempty,gte=0"`
			Status      string  `json:"status" validate:"omitempty,oneof=in_stock out_of_stock"`
		}

		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		// Fetch existing product
		product, err := services.GetProductByID(db, productID)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Product not found",
			})
		}

		// Update fields if provided
		if req.Name != "" {
			product.Name = req.Name
		}
		if req.Description != "" {
			product.Description = req.Description
		}
		if req.Price > 0 {
			product.Price = req.Price
		}
		if req.CategoryID > 0 {
			product.CategoryID = req.CategoryID
		}
		if req.Stock >= 0 {
			product.Stock = req.Stock
		}
		if req.Status != "" {
			product.Status = req.Status
		}

		// Update product in database
		if err := services.UpdateProduct(db, product); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to update product",
			})
		}

		return c.JSON(fiber.Map{
			"message": "Product updated successfully",
		})
	}
}

// DeleteProduct handles deleting or archiving a product
func DeleteProduct(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		productID := c.Params("id")

		// Delete or archive logic (soft delete example)
		if err := services.SoftDeleteProduct(db, productID); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to delete product",
			})
		}

		return c.JSON(fiber.Map{
			"message": "Product deleted successfully",
		})
	}
}
