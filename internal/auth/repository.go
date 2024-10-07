package auth

import (
	"fmt"

	util_error "github.com/PICH-IO/admin-api/pkg/utils/errors"
	"github.com/jmoiron/sqlx"
)

type AuthRepository interface {
	AuthLogin(username string) (*UserDataRes, *util_error.ErrorResponse)
	FetchPermissionsForRole(roleID int) ([]string, error)
}
type authRepository struct {
	db *sqlx.DB
}

func NewAuthReposity(db *sqlx.DB) AuthRepository {
	return &authRepository{
		db: db,
	}
}

func (r *authRepository) AuthLogin(username string) (*UserDataRes, *util_error.ErrorResponse) {
	var user UserDataRes

	query := `SELECT 
				user_id, username, password, role_id, status 
				FROM tbl_users
				WHERE username=$1`

	fmt.Println(query)
	err := r.db.Get(&user, query, username)

	fmt.Println("query data is :", err)
	if err != nil {
		return nil, util_error.NewError("RecordNotFound", "User not found or inactive")
	}

	if user.Status != "active" {
		return nil, util_error.NewError("UserInactive", "User account is inactive")
	}
	return &user, nil
}

func (r *authRepository) FetchPermissionsForRole(roleID int) ([]string, error) {
	var permissions []string
	query := `
        SELECT p.permission_name
        FROM tbl_permissions p
        JOIN tbl_role_permissions rp ON rp.permission_id = p.permission_id
        WHERE rp.role_id = $1
    `
	err := r.db.Select(&permissions, query, roleID)
	if err != nil {
		return nil, err
	}
	return permissions, nil
}
