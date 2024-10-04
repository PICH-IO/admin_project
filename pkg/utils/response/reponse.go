package util_response

// struct for normal response
type BodyResponse struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data"`
}

// struct for paging response
type BodyResponseWithPaging struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data"`
	Page       int         `json:"page"`
	PerPage    int         `json:"perpage"`
	Total      int         `json:"total"`
}

// function for normal response
func HttpResponse(
	success bool,
	message string,
	status_code int,
	data interface{},

) *BodyResponse {
	return &BodyResponse{
		Success:    success,
		Message:    message,
		StatusCode: status_code,
		Data:       data,
	}
}

// function for paging response
func HttpResponsePaging(
	success bool,
	message string,
	status_code int,
	data interface{},
	page int,
	perpage int,
	total int,
) *BodyResponseWithPaging {
	return &BodyResponseWithPaging{
		Success:    success,
		Message:    message,
		StatusCode: status_code,
		Data:       data,
		Page:       page,
		PerPage:    perpage,
		Total:      total,
	}
}
