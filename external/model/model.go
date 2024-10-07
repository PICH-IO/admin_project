// pkg_models/models.go
package pkg_models

import "time"

// User represents a user in the system
type User struct {
	UserID       int        `db:"user_id" json:"user_id"`
	Username     string     `db:"username" json:"username"`
	Email        string     `db:"email" json:"email"`
	PasswordHash string     `db:"password_hash" json:"-"`
	RoleID       int        `db:"role_id" json:"role_id"`
	CreatedAt    time.Time  `db:"created_at" json:"created_at"`
	CreatedBy    int        `db:"created_by" json:"created_by,omitempty"`
	UpdatedAt    time.Time  `db:"updated_at" json:"updated_at,omitempty"`
	UpdatedBy    int        `db:"updated_by" json:"updated_by,omitempty"`
	DeletedAt    *time.Time `db:"deleted_at" json:"deleted_at,omitempty"`
	DeletedBy    int        `db:"deleted_by" json:"deleted_by,omitempty"`
}

// Role represents a user role
type Role struct {
	RoleID    int       `db:"role_id" json:"role_id"`
	RoleName  string    `db:"role_name" json:"role_name"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	// Add other fields as necessary
}

// Permission represents a permission
type Permission struct {
	PermissionID   int    `db:"permission_id" json:"permission_id"`
	PermissionName string `db:"permission_name" json:"permission_name"`
}

// Token represents the JWT token payload
type Token struct {
	ID          int      `json:"id"`
	Username    string   `json:"username"`
	Permissions []string `json:"permissions"`
	RoleID      int      `json:"role_id"`
}

// Customer represents a customer
type Customer struct {
	CustomerID int        `db:"customer_id" json:"customer_id"`
	UserID     int        `db:"user_id" json:"user_id"`
	FirstName  string     `db:"first_name" json:"first_name"`
	LastName   string     `db:"last_name" json:"last_name"`
	Phone      string     `db:"phone" json:"phone,omitempty"`
	Address    string     `db:"address" json:"address,omitempty"`
	City       string     `db:"city" json:"city,omitempty"`
	PostalCode string     `db:"postal_code" json:"postal_code,omitempty"`
	Country    string     `db:"country" json:"country,omitempty"`
	CreatedAt  time.Time  `db:"created_at" json:"created_at"`
	CreatedBy  int        `db:"created_by" json:"created_by,omitempty"`
	UpdatedAt  *time.Time `db:"updated_at" json:"updated_at,omitempty"`
	UpdatedBy  int        `db:"updated_by" json:"updated_by,omitempty"`
	DeletedAt  *time.Time `db:"deleted_at" json:"deleted_at,omitempty"`
	DeletedBy  int        `db:"deleted_by" json:"deleted_by,omitempty"`
}

// Product represents a product in the system
type Product struct {
	ProductID   int        `db:"product_id" json:"product_id"`
	Name        string     `db:"name" json:"name"`
	Description string     `db:"description" json:"description,omitempty"`
	Price       float64    `db:"price" json:"price"`
	CategoryID  int        `db:"category_id" json:"category_id"`
	Stock       int        `db:"stock" json:"stock"`
	Status      string     `db:"status" json:"status"` // in_stock or out_of_stock
	CreatedAt   time.Time  `db:"created_at" json:"created_at"`
	CreatedBy   int        `db:"created_by" json:"created_by,omitempty"`
	UpdatedAt   *time.Time `db:"updated_at" json:"updated_at,omitempty"`
	UpdatedBy   int        `db:"updated_by" json:"updated_by,omitempty"`
	DeletedAt   *time.Time `db:"deleted_at" json:"deleted_at,omitempty"`
	DeletedBy   int        `db:"deleted_by" json:"deleted_by,omitempty"`
}

// Category represents a product category
type Category struct {
	CategoryID  int        `db:"category_id" json:"category_id"`
	Name        string     `db:"name" json:"name"`
	Description string     `db:"description" json:"description,omitempty"`
	CreatedAt   time.Time  `db:"created_at" json:"created_at"`
	CreatedBy   int        `db:"created_by" json:"created_by,omitempty"`
	UpdatedAt   *time.Time `db:"updated_at" json:"updated_at,omitempty"`
	UpdatedBy   int        `db:"updated_by" json:"updated_by,omitempty"`
	DeletedAt   *time.Time `db:"deleted_at" json:"deleted_at,omitempty"`
	DeletedBy   int        `db:"deleted_by" json:"deleted_by,omitempty"`
}
