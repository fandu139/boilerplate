package v1

type Pagination struct {
	Limit  uint `form:"limit"`
	Offset uint `form:"offset"`
}
