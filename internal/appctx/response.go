package appctx

const (
	ResponseNameSuccess = "SUCCESS"
	ResponseBadRequest  = "BAD_REQUEST"
)

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Errors  any    `json:"errors,omitempty"`
	Meta    any    `json:"_meta,omitempty"`
}

type Metadata struct {
	Limit       int `json:"limit"`
	CurrentPage int `json:"page"`
	TotalPage   int `json:"total_page"`
	TotalCount  int `json:"total_count"`
}
