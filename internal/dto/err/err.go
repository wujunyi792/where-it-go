package err

import "errors"

type ServiceError struct {
	Code int
	Err  error
}

func (e *ServiceError) Error() string {
	return e.Err.Error()
}

var (
	RequestErr = ServiceError{
		Code: 40200,
		Err:  errors.New("请求参数错误"),
	}
	CodeErr = ServiceError{
		Code: 40101,
		Err:  errors.New("验证码校验失败"),
	}
	LoginErr = ServiceError{
		Code: 40102,
		Err:  errors.New("用户名或密码错误"),
	}
	JWTErr = ServiceError{
		Code: 40103,
		Err:  errors.New("登录过期"),
	}
	VerifyErr = ServiceError{
		Code: 40104,
		Err:  errors.New("访问受限"),
	}
	OauthErr = ServiceError{
		Code: 40105,
		Err:  errors.New("state 错误"),
	}
	InternalErr = ServiceError{
		Code: 50000,
		Err:  errors.New("内部错误"),
	}
)
