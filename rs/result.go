package rs

type Result struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func Success() Result {
	return Result{
		Status:  200,
		Message: "操作成功",
	}
}
func SuccessWithData[T any](data T) Result {
	return Result{
		Status:  200,
		Message: "操作成功",
		Data:    data,
	}
}

func Failure(msg string) Result {
	return Result{
		Status:  500,
		Message: msg,
	}
}
