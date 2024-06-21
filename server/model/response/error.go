package response

// Error 异常
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var (
	Success = Error{Code: 0, Message: "success"}

	ErrOperate  = Error{Code: -1, Message: "operate failed"}
	ErrUnAuth   = Error{Code: -2, Message: "unauthorized"}
	ErrNotFound = Error{Code: -3, Message: "not found"}
)
