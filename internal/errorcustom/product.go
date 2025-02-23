package errorcustom

import "fmt"

var (
	ErrProductNotFound = fmt.Errorf("product not found")
)
