package pkg_models

// token response
type Token struct {
	Id       float64 `json:"id"`
	Username string  `json:"user_name"`
	// MembershipId   float64 `json:"membership_id"`
	// MembershipRole string  `json:"membership_role"`
	// RoleId         float64 `json:"role_id"`
}

// paging option request
type PagingOption struct {
	PerPage int `json:"perpage" query:"per_page" validate:"required"`
	Page    int `json:"page" query:"page" validete:"required"`
}

// filter request
type Filter struct {
	Property string      `json:"property" query:"property"`
	Value    interface{} `json:"value" query:"value"`
}

// sort for request
type Sort struct {
	Property  string `json:"property" query:"property"`
	Direction string `json:"direction" query:"direction"`
}
