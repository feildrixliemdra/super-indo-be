package errorcustom

import "fmt"

var (
	ErrCategoryAlreadyExists = fmt.Errorf("category already exists")
	ErrCategoryNotFound      = fmt.Errorf("category not found")
)
