package xerr

var message map[uint32]string

func init() {
	message = make(map[uint32]string)
	message[OK] = "Success"
	message[SERVER_COMMON_ERROR] = "The server is malfunctioning, please try again later"
	message[REUQEST_PARAM_ERROR] = "parameter error"
	message[TOKEN_EXPIRE_ERROR] = "Token invalid, please log in again"
	message[UNAUTHORIZED_ERROR] = "Unauthorized operation"
	message[TOKEN_GENERATE_ERROR] = "Token generation failed"
}

func MapErrMsg(errcode uint32) string {
	if msg, ok := message[errcode]; ok {
		return msg
	} else {
		return message[SERVER_COMMON_ERROR]
	}
}

func IsCodeErr(errcode uint32) bool {
	if _, ok := message[errcode]; ok {
		return true
	} else {
		return false
	}
}
