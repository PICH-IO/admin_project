package auth

type UserRes struct {
	Token string `json:"token"`
}

type UserReq struct {
	Username string `json:"username" validate:"required,username"`
	Password string `json:"password" validate:"required"`
}

type UserDataRes struct {
	Id       int    `db:"user_id" json:"id"`
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"-"`
	RoleID   int    `db:"role_id" json:"role_id"`
	Status   string `db:"status" json:"status"`
}
