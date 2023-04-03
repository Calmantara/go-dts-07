package response

const (
	InvalidParam       = "invalid param request"
	InvalidBody        = "invalid body request"
	InvalidQuery       = "invalid query request"
	SomethingWentWrong = "something went wrong"
	Unauthorized       = "unauthorized request"
)

type SuccessResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}
