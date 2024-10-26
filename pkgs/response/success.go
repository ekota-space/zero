package response

type SuccessDataResponse[T any] struct {
	Data T `json:"data"`
}
