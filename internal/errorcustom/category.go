package errorcustom

import "fmt"

var (
	ErrCategoryAlreadyExists = fmt.Errorf("category already exists")
)
