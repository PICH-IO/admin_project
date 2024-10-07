package auth

import (
	pkg_jwt "github.com/PICH-IO/admin-api/pkg/jwt"
	pkg_models "github.com/PICH-IO/admin-api/pkg/models"
	util_error "github.com/PICH-IO/admin-api/pkg/utils/errors"
)

type AuthService interface {
	Login(req UserReq) (*UserRes, *util_error.ErrorResponse)
}

type authService struct {
	repo AuthRepository
}

func NewAuthService(repo AuthRepository) AuthService {
	return &authService{
		repo: repo,
	}
}

func (s *authService) Login(req UserReq) (*UserRes, *util_error.ErrorResponse) {

	user, errlogin := s.repo.AuthLogin(req.Username)
	if errlogin != nil {
		return nil, errlogin
	}

	permission, errPerm := s.repo.FetchPermissionsForRole(user.RoleID)
	if errPerm != nil {
		return nil, util_error.NewError("Failed", "Failed to fetch permissions")
	}

	// Create token payload
	var usertoken = &pkg_models.Token{
		Id:          float64(user.Id),
		Username:    user.Username,
		RoleId:      float64(user.RoleID),
		Permissions: permission,
	}
	// fmt.Println(usertoken)

	// Generate JWT token
	token, err := pkg_jwt.GenerateJWT(usertoken)
	if err != nil {
		return nil, util_error.NewError("ErrorGeneratingJWT", "The JWT could not be generated. Please try again later.")
	}

	return &UserRes{
		Token: token,
	}, nil
}
