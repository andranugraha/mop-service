package constant

import "fmt"

var (
	ErrInvalidIdentifierOrPassword = fmt.Errorf("invalid identifier or password")
	ErrInvalidParam                = fmt.Errorf("invalid param")
	ErrInternalServer              = fmt.Errorf("internal server error")
	ErrItemNotFound                = fmt.Errorf("one or more items may not exist or have been removed")
	ErrMerchantNotFound            = fmt.Errorf("merchant not found")
	ErrTableNotFound               = fmt.Errorf("table not found")
)

var (
	ErrOrderTotalPriceMismatch       = fmt.Errorf("order total price mismatch")
	ErrOrderTableAndMerchantMismatch = fmt.Errorf("order table and merchant mismatch")
)

const (
	ErrCodeInternalServer = int32(10000)
	ErrCodeInvalidParam   = int32(10001)
	ErrCodeTimeout        = int32(10002)

	ErrCodeOrderPriceMismatch            = int32(11000)
	ErrCodeOrderTableAndMerchantMismatch = int32(11001)
)

var errToCode = map[error]int32{
	ErrInternalServer:                ErrCodeInternalServer,
	ErrInvalidParam:                  ErrCodeInvalidParam,
	ErrInvalidIdentifierOrPassword:   ErrCodeInvalidParam,
	ErrItemNotFound:                  ErrCodeInvalidParam,
	ErrMerchantNotFound:              ErrCodeInvalidParam,
	ErrTableNotFound:                 ErrCodeInvalidParam,
	ErrOrderTotalPriceMismatch:       ErrCodeOrderPriceMismatch,
	ErrOrderTableAndMerchantMismatch: ErrCodeOrderTableAndMerchantMismatch,
}

func GetErrorCode(err error) int32 {
	if code, ok := errToCode[err]; ok {
		return code
	}
	return ErrCodeInternalServer
}
