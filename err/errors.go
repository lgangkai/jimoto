package err

import "github.com/asim/go-micro/v3/errors"

const (
	SUCCESS = 0

	ERR_PARAMETER_PARSE_FAILED = 1
	ERR_INTERNAL_ERROR         = 2
	ERR_BAD_REQUEST            = 3

	ERR_GET_PROFILE_FAILED    = 100001
	ERR_UPDATE_PROFILE_FAILED = 100002
	ERR_DELETE_PROFILE_FAILED = 100003
	ERR_CREATE_PROFILE_FAILED = 100004

	ERR_EMAIL_IS_REGISTERED = 200001
	ERR_REGISTER_INTERNAL   = 200002
	ERR_REGISTER_REQUEST    = 200003
	ERR_LOGIN_NO_USER       = 200004
	ERR_LOGIN_INTERNAL      = 200005
	ERR_LOGIN_REQUEST       = 200006
	ERR_PASSWORD_MISMATCH   = 200007
	ERR_AUTH_FAILED         = 200008
	ERR_TOKEN_EXPIRED       = 200009

	ERR_GET_COMMODITY_FAILED        = 300001
	ERR_UPDATE_COMMODITY_FAILED     = 300002
	ERR_DELETE_COMMODITY_FAILED     = 300003
	ERR_CREATE_COMMODITY_FAILED     = 300004
	ERR_GET_COMMODITY_LIST_FAILED   = 300005
	ERR_GET_COMMODITY_IMAGES_FAILED = 300006
)

var errMsg = map[int32]string{
	SUCCESS: "Success.",

	ERR_PARAMETER_PARSE_FAILED: "Parse parameter failed.",
	ERR_INTERNAL_ERROR:         "Internal error.",
	ERR_BAD_REQUEST:            "Bad request.",

	ERR_GET_PROFILE_FAILED:    "Get profile failed.",
	ERR_UPDATE_PROFILE_FAILED: "Update profile failed.",
	ERR_DELETE_PROFILE_FAILED: "Delete profile failed.",
	ERR_CREATE_PROFILE_FAILED: "Create profile failed.",

	ERR_EMAIL_IS_REGISTERED: "Register failed, email has been registered.",
	ERR_REGISTER_INTERNAL:   "Register failed, internal server error.",
	ERR_REGISTER_REQUEST:    "Register failed, bad request.",
	ERR_LOGIN_NO_USER:       "Login failed, no such user.",
	ERR_LOGIN_INTERNAL:      "Login failed, internal server error.",
	ERR_PASSWORD_MISMATCH:   "Login failed, password mismatch.",
	ERR_LOGIN_REQUEST:       "Login failed, bad request.",
	ERR_AUTH_FAILED:         "Auth failed, invalid token.",
	ERR_TOKEN_EXPIRED:       "Auth failed, login status expired",

	ERR_GET_COMMODITY_FAILED:      "Get commodity failed.",
	ERR_GET_COMMODITY_LIST_FAILED: "Get commodity list failed",
}

func New(code int32) error {
	return errors.New("", errMsg[code], code)
}

func GetMsg(code int32) string {
	return errMsg[code]
}

//type Err struct {
//	Code uint32
//}
//
//func (e Err) Error() string {
//	return errMsg[e.Code]
//}
