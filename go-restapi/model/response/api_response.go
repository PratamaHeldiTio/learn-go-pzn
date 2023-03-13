package response

type APIResponse struct {
	Code   int
	Status string
	Data   interface{}
}
