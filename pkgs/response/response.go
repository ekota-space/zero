package response

func Success[T any](data T) SuccessDataResponse[T] {
	return SuccessDataResponse[T]{Data: data}
}

func Error[T any](err T) ErrorResponse[T] {
	return ErrorResponse[T]{Error: err}
}
