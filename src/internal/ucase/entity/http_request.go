package entity

// Pagination ...
type Pagination struct {
	Page  int `json:"page" form:"page" validate:"min=0"`
	Limit int `json:"limit" form:"limit" validate:"min=1"`
}

// ExampleRequest ...
// payload := &entity.ExampleRequest{}
// context.ShouldBind(payload)
// err := handler.Validator.Request(payload)
// if err != nil {
// 	handler.Rest.ErrorResponse(err)
// 	return
// }
type ExampleRequest struct {
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}
