package request

type StatusResponse struct {
	Status string `json:"status"`
}

func OKStatusResponse() StatusResponse {
	return StatusResponse{Status: "OK"}
}

type Pagination struct {
	Total    int `json:"total" example:"527"`
	PageSize int `json:"page_size" example:"100"`
	Page     int `json:"page" example:"6"`
	Count    int `json:"count" example:"27"`
}
