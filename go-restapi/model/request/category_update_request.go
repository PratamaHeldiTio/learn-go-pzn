package request

type CategoryUpdateRequest struct {
	Id   int    `validate:"required"`
	Name string `validate:"required,min=1,max=200"`
}
