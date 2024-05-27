package exception

// Exception 异常
type Exception struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var (
	Success  = Exception{Code: 0, Message: "success"}
	Operate  = Exception{Code: 101, Message: "operate failed"}
	UnAuth   = Exception{Code: 401, Message: "unauthorized"}
	NotFound = Exception{Code: 404, Message: "not found"}
)
