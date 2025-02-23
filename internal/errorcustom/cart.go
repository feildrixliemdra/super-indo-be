package errorcustom

import "fmt"

var (
	ErrCartNotFound = fmt.Errorf("cart not found")
)
