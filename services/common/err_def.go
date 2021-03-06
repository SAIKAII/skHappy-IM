package common

const COMMON_ERROR = 1000
const CLIENT_ERROR = 2000
const INTERNEL_ERROR = 3000

// 一般错误
const (
	COMMON_UNKNOWN_ERROR = COMMON_ERROR + iota
	COMMON_USER_NOT_FOUND_ERROR
	COMMON_PWD_NOT_MATCH_ERROR
	COMMON_JWT_VERIFY_ERROR
)

// 客户端错误
const (
	CLIENT_UNKNOWN_ERROR = CLIENT_ERROR + iota
	CLIENT_REQUEST_PARAMS_ERROR
)

// 服务端错误
const (
	INTERNEL_UNKNOWN_ERROR = INTERNEL_ERROR + iota
	INTERNEL_UNMARSHAL_ERROR
	INTERNEL_GENERATE_JWT_ERROR
)
