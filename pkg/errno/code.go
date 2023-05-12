// This file is designed to define any error code
package errno

const (
	SuccessCode = 10000
	SuccessMsg  = "ok"

	ServiceErrorCode           = 10001 // 未知错误
	ParamErrorCode             = 10002 // 参数错误
	AuthorizationFailedErrCode = 10003 // 鉴权失败
	AuthorizationExpiredCode   = 10004 // 鉴权过期
	UnexpectedTypeErrorCode    = 10005 // 未知类型
	NotImplementErrorCode      = 10006 // 未实装
)

var (
	Success = NewErrNo(SuccessCode, SuccessMsg)

	ServiceError         = NewErrNo(ServiceErrorCode, "service is unable to start successfully")
	ServiceInternalError = NewErrNo(ServiceErrorCode, "service internal error")
	ParamError           = NewErrNo(ParamErrorCode, "parameter error")

	// Auth
	AuthorizationFailError    = NewErrNo(AuthorizationFailedErrCode, "authorization failed")
	AuthorizationExpiredError = NewErrNo(AuthorizationExpiredCode, "token expired")

	// User
	ErrUsernameAlreadyExists = NewErrNo(ParamErrorCode, "user existed")
)
