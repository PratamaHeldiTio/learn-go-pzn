package request

// add validate tag for validation in service
type CategoryCreatedRequest struct {
	Name string `validate:"required,min=1,max=200"`
}
