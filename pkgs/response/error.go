package response

type ErrorResponse[T any] struct {
	Error T `json:"error"`
}
