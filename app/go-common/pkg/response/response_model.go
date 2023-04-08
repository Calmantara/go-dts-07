package response

const (
	InvalidParam       = "invalid param request"
	InvalidBody        = "invalid body request"
	InvalidPayload     = "invalid payload request"
	InvalidQuery       = "invalid query request"
	InternalServer     = "internal server error"
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
