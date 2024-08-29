package constant

import "fmt"

var (
	ErrInvalidIdentifierOrPassword = fmt.Errorf("invalid identifier or password")
	ErrInvalidParam                = fmt.Errorf("invalid param")
	ErrInternalServer              = fmt.Errorf("internal server error")
	ErrItemNotFound                = fmt.Errorf("one or more items do not exist or have been removed")
)

const (
	ErrCodeInternalServer = int32(10000)
	ErrCodeInvalidParam   = int32(10001)
)
