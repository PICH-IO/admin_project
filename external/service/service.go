// pkg_services/services.go
package services

import (
	"github.com/PICH-IO/admin-api/pkg/models"
	"github.com/jmoiron/sqlx"
)

// GetRoleIDByName retrieves the role ID by role name
func GetRoleIDByName(db *sqlx.DB, roleName string) (int, error) {
	var role models.Role
	query := `SELECT * FROM tbl_roles WHERE role_name = $1`
	err := db.Get(&role, query, roleName)
	if err != nil {
		return 0, err
	}
	return role.RoleID, nil
}

// CreateUser inserts a new user into the database
func CreateUser(db *sqlx.DB, user *models.User) (int, error) {
	query := `
        INSERT INTO tbl_users (username, email, password_hash, role_id, created_at)
        VALUES ($1, $2, $3, $4, NOW())
        RETURNING user_id
    `
	var userID int
	err := db.Get(&userID, query, user.Username, user.Email, user.PasswordHash, user.RoleID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

// GetUserByUsername retrieves a user by username
func GetUserByUsername(db *sqlx.DB, username string) (*models.User, error) {
	var user models.User
	query := `SELECT * FROM tbl_users WHERE username = $1 AND deleted_at IS NULL`
	err := db.Get(&user, query, username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByID retrieves a user by user ID
func GetUserByID(db *sqlx.DB, userID string) (*models.User, error) {
	var user models.User
	query := `SELECT * FROM tbl_users WHERE user_id = $1 AND deleted_at IS NULL`
	err := db.Get(&user, query, userID)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUser updates user information
func UpdateUser(db *sqlx.DB, user *models.User) error {
	query := `
        UPDATE tbl_users
        SET email = $1, password_hash = $2, updated_at = NOW(), updated_by = $3
        WHERE user_id = $4
    `
	_, err := db.Exec(query, user.Email, user.PasswordHash, user.UpdatedBy, user.UserID)
	return err
}

// SoftDeleteUser performs a soft delete on a user
func SoftDeleteUser(db *sqlx.DB, userID string) error {
	query := `
        UPDATE tbl_users
        SET deleted_at = NOW()
        WHERE user_id = $1
    `
	_, err := db.Exec(query, userID)
	return err
}

// FetchPermissionsForRole fetches permissions for a given role ID
func FetchPermissionsForRole(db *sqlx.DB, roleID int) ([]string, error) {
	var permissions []string
	query := `
        SELECT p.permission_name
        FROM tbl_permissions p
        JOIN tbl_role_permissions rp ON rp.permission_id = p.permission_id
        WHERE rp.role_id = $1
    `
	err := db.Select(&permissions, query, roleID)
	if err != nil {
		return nil, err
	}
	return permissions, nil
}

// GetUsersByRoleName retrieves all users with a specific role name
func GetUsersByRoleName(db *sqlx.DB, roleName string) ([]models.User, error) {
	var users []models.User
	query := `
        SELECT u.*
        FROM tbl_users u
        JOIN tbl_roles r ON u.role_id = r.role_id
        WHERE r.role_name = $1 AND u.deleted_at IS NULL
    `
	err := db.Select(&users, query, roleName)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// CreateProduct inserts a new product into the database
func CreateProduct(db *sqlx.DB, product *models.Product) (int, error) {
	query := `
        INSERT INTO tbl_products (name, description, price, category_id, stock, status, created_at)
        VALUES ($1, $2, $3, $4, $5, $6, NOW())
        RETURNING product_id
    `
	var productID int
	err := db.Get(&productID, query, product.Name, product.Description, product.Price, product.CategoryID, product.Stock, product.Status)
	if err != nil {
		return 0, err
	}
	return productID, nil
}

// GetAllProducts retrieves all products
func GetAllProducts(db *sqlx.DB) ([]models.Product, error) {
	var products []models.Product
	query := `SELECT * FROM tbl_products WHERE deleted_at IS NULL`
	err := db.Select(&products, query)
	if err != nil {
		return nil, err
	}
	return products, nil
}

// GetProductByID retrieves a product by product ID
func GetProductByID(db *sqlx.DB, productID string) (*models.Product, error) {
	var product models.Product
	query := `SELECT * FROM tbl_products WHERE product_id = $1 AND deleted_at IS NULL`
	err := db.Get(&product, query, productID)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// UpdateProduct updates product information
func UpdateProduct(db *sqlx.DB, product *models.Product) error {
	query := `
        UPDATE tbl_products
        SET name = $1, description = $2, price = $3, category_id = $4, stock = $5, status = $6, updated_at = NOW(), updated_by = $7
        WHERE product_id = $8
    `
	_, err := db.Exec(query, product.Name, product.Description, product.Price, product.CategoryID, product.Stock, product.Status, product.UpdatedBy, product.ProductID)
	return err
}

// SoftDeleteProduct performs a soft delete on a product
func SoftDeleteProduct(db *sqlx.DB, productID string) error {
	query := `
        UPDATE tbl_products
        SET deleted_at = NOW()
        WHERE product_id = $1
    `
	_, err := db.Exec(query, productID)
	return err
}
