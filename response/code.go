/*
*

	@author: taco
	@Date: 2023/8/15
	@Time: 15:38

*
*/
package response

import "errors"

const (
	SUCCESS int32 = iota + 1000
)

const (
	ACCESS_TOKEN_INVALID int32 = iota + 2000
	ACCESS_EXPIRED
	ACCESS_DENY
	ACCESS_NOT_FOUND
	ACCESS_PWD_WRONG
	ACCESS_ACCOUNT_INVALID
	ACCESS_KEY_INVALID
	ACCOUNT_ALREADY_EXISTS
	ACCESS_CODE_WRONG
	GROUP_ALREADY_EXISTS
	ACCESS_TOO_FAST
	DELETE_ADMIN_WRONG
	CANT_CREATE_GROUP
	CANT_CREATE_ACCOUNT
	REFRESH_EXPIRED

	ACCESS_NOT_TRUST_DEVICE

	DATA_EXISTS
	ACCOUNT_NOT_FOUND

	NOT_ALLOW_PARAMS
	SECURITY_TOKEN_INVALID
	SECURITY_TOKEN_EXPIRED
)

const (
	NOT_FOUND int32 = iota + 3000
)

const (
	FAIL int32 = iota + 4000
	WRONG_PARAM
	NOT_FOUND_METHOD
	METADATA_NOT_FOUND
	AUTHORIZATION_NOT_FOUND
	ACCESSKEY_NOT_FOUND
	WRONG_CAPTCHA
	WECHAT_ERR_USERTOKEN_EXPIRED
	MOVIE_EXIST
	CHAPTER_EXIST
	EPISODE_EXIST
	SCENE_EXIST
	DATA_EXIST
	NOT_INVITED
	ACCESS_DENIED
	NOT_ADMIN
)

const (
	SERVER_WRONG int32 = 5000 + iota
)

const (
	OPERATE_ARTICLE_STATUS_ERR int32 = 6000 + iota
	OPERATE_LABEL_STATUS_ERR
)

const (

	// note: sdk error code %5d
	ERR_INIT_SDK_NOT_CLIENT int32 = 10001 + iota
	ERR_LOGININFO_NIL
	ERR_JSON_MARSHAL
	ERR_INIT_SDK_NOT_LOGINED
)

const (
	// note: dbl 游戏相关
	ERR_SCENE_LOCK int32 = 600001 + iota
	USER_HAS_PAYMENT
	USER_NO_PAYMENT
	USER_PAYMENT_SUCCESS
	USER_PAYMENT_TIMEOUT
	USER_PAYMENT_PROCESSING
	USER_PAYMENT_FAIL
	TRANSACTION_ERR_RESOURCES_OCCUPIED
)

var WrongMessageEn = map[int32]string{
	SUCCESS:                            "success",
	ACCESS_TOKEN_INVALID:               "invalid token",
	ACCESS_EXPIRED:                     "user licence expired",
	REFRESH_EXPIRED:                    "refresh licence expired",
	ACCESS_NOT_TRUST_DEVICE:            "not a trusted device, need secondary verification",
	ACCESS_DENY:                        "permission denied",
	ACCESS_NOT_FOUND:                   "account does not exist",
	ACCESS_PWD_WRONG:                   "incorrect username or password",
	ACCESS_ACCOUNT_INVALID:             "account is disable",
	ACCESS_KEY_INVALID:                 "AccessKey is invalid",
	ACCOUNT_ALREADY_EXISTS:             "user already exists",
	ACCESS_CODE_WRONG:                  "verification code error",
	ACCESS_TOO_FAST:                    "Access too fast",
	GROUP_ALREADY_EXISTS:               "user group already exists",
	DATA_EXISTS:                        "data exists",
	CANT_CREATE_GROUP:                  "Super administrator cannot create groups",
	CANT_CREATE_ACCOUNT:                "unable to create sub-account, please use root account to create one",
	NOT_FOUND:                          "record not found",
	FAIL:                               "fail",
	NOT_FOUND_METHOD:                   "request method not found",
	WRONG_PARAM:                        "param error",
	METADATA_NOT_FOUND:                 "metadata not found",
	AUTHORIZATION_NOT_FOUND:            "authorization not found",
	ACCESSKEY_NOT_FOUND:                "accesskey not found",
	ACCOUNT_NOT_FOUND:                  "account not found",
	WRONG_CAPTCHA:                      "wrong captcha",
	WECHAT_ERR_USERTOKEN_EXPIRED:       "wechat user_token is expired",
	MOVIE_EXIST:                        "movie already exists",
	CHAPTER_EXIST:                      "chapter already exists",
	EPISODE_EXIST:                      "episode already exists",
	SCENE_EXIST:                        "scene already exists",
	DATA_EXIST:                         "data already exists",
	NOT_INVITED:                        "not invited user",
	NOT_ADMIN:                          "not admin user",
	ACCESS_DENIED:                      "access denied",
	DELETE_ADMIN_WRONG:                 "super administrator cannot be deleted",
	SERVER_WRONG:                       "Internal Server Error",
	OPERATE_ARTICLE_STATUS_ERR:         "The article is on the shelf and cannot be operated",
	OPERATE_LABEL_STATUS_ERR:           "Tab is open and not operable",
	ERR_INIT_SDK_NOT_CLIENT:            "sdk client is nil",
	ERR_LOGININFO_NIL:                  "reset time, logininfo is nil",
	ERR_JSON_MARSHAL:                   "json marshal err",
	ERR_INIT_SDK_NOT_LOGINED:           "sdk client isn't logined",
	ERR_SCENE_LOCK:                     "the scene not unlock",
	USER_PAYMENT_SUCCESS:               "payment success",
	USER_PAYMENT_TIMEOUT:               "payment timeout",
	USER_PAYMENT_PROCESSING:            "payment processing",
	USER_PAYMENT_FAIL:                  "payment fail",
	TRANSACTION_ERR_RESOURCES_OCCUPIED: "transaction resources occupied",
	NOT_ALLOW_PARAMS:                   "not allow params",
	SECURITY_TOKEN_INVALID:             "security token invalid",
	SECURITY_TOKEN_EXPIRED:             "security token expired",
}

type ApiResponse struct {
	Code    int32       `json:"code"`    // 状态码
	Message string      `json:"message"` // 状态短语
	Result  interface{} `json:"result"`  // 数据结果集
}

var WrongMessageZh = map[int32]string{
	SUCCESS:                            "请求成功",
	ACCESS_TOKEN_INVALID:               "无效token",
	ACCESS_EXPIRED:                     "用户凭证过期",
	REFRESH_EXPIRED:                    "刷新凭证过期",
	ACCESS_NOT_TRUST_DEVICE:            "不是可信设备，需要二次验证",
	ACCESS_DENY:                        "权限不足",
	ACCESS_NOT_FOUND:                   "账户不存在",
	ACCESS_PWD_WRONG:                   "用户名或密码不正确",
	ACCESS_ACCOUNT_INVALID:             "账户已经被冻结",
	ACCESS_KEY_INVALID:                 "AccessKey无效",
	ACCOUNT_ALREADY_EXISTS:             "用户已存在",
	ACCESS_TOO_FAST:                    "太频繁了",
	ACCESS_CODE_WRONG:                  "验证码错误",
	DELETE_ADMIN_WRONG:                 "超级管理员不可删除",
	GROUP_ALREADY_EXISTS:               "用户组已存在",
	DATA_EXISTS:                        "数据已存在",
	CANT_CREATE_GROUP:                  "超级管理员不可创建组",
	CANT_CREATE_ACCOUNT:                "无法创建子账号,请用根账号创建",
	MOVIE_EXIST:                        "该标题的影剧已经存在",
	CHAPTER_EXIST:                      "该影剧下此标题的章节已经存在",
	EPISODE_EXIST:                      "该章节下此标题的剧集已经存在",
	SCENE_EXIST:                        "该章节下此标题的场景已经存在",
	DATA_EXIST:                         "该标题的数据已经存在",
	NOT_INVITED:                        "不是受邀用户",
	NOT_ADMIN:                          "不是管理员用户",
	ACCESS_DENIED:                      "访问的资源没有足够的权限",
	NOT_FOUND:                          "记录未找到",
	FAIL:                               "请求失败",
	WRONG_PARAM:                        "参数错误",
	NOT_FOUND_METHOD:                   "未找到请求方法",
	METADATA_NOT_FOUND:                 "没找到metadata",
	AUTHORIZATION_NOT_FOUND:            "没找到验证头",
	ACCESSKEY_NOT_FOUND:                "没找到用户appid",
	ACCOUNT_NOT_FOUND:                  "没找到用户头信息",
	WRONG_CAPTCHA:                      "验证码错误",
	WECHAT_ERR_USERTOKEN_EXPIRED:       "微信授权中用户的token已过期",
	SERVER_WRONG:                       "服务器错误",
	OPERATE_ARTICLE_STATUS_ERR:         "文章处于上架状态，不可操作",
	OPERATE_LABEL_STATUS_ERR:           "标签处于开放状态，不可操作",
	ERR_INIT_SDK_NOT_CLIENT:            "客户端尚未完成初始化",
	ERR_LOGININFO_NIL:                  "重置过期时间时，返回的登录信息为空",
	ERR_JSON_MARSHAL:                   "json序列化错误",
	ERR_INIT_SDK_NOT_LOGINED:           "sdk尚未登录",
	ERR_SCENE_LOCK:                     "该场景尚未解锁，请通关相关剧情",
	USER_PAYMENT_SUCCESS:               "支付成功",
	USER_PAYMENT_TIMEOUT:               "支付超时",
	USER_PAYMENT_PROCESSING:            "支付处理中",
	USER_PAYMENT_FAIL:                  "支付失败",
	TRANSACTION_ERR_RESOURCES_OCCUPIED: "该资源被其他资源占用",
	NOT_ALLOW_PARAMS:                   "不被允许的参数",
	SECURITY_TOKEN_INVALID:             "错误的安全令牌",
	SECURITY_TOKEN_EXPIRED:             "安全令牌已过期",
}

func StatusToErr(code int32, v ...any) error {
	if len(v) > 0 {
		lang := v[0].(string)
		if lang == "" || len(lang) <= 0 {
			lang = "zh"
		}
		if lang == "zh" {
			return toErr(WrongMessageZh[code])
		}
		return toErr(WrongMessageEn[code])
	} else {
		return toErr(WrongMessageZh[code])
	}
}

func toErr(str string) error {
	return errors.New(str)
}

func StatusText(code int32, v ...any) string {
	if len(v) > 0 {
		lang := v[0].(string)
		if lang == "" || len(lang) <= 0 {
			lang = "zh"
		}
		if lang == "zh" {
			return WrongMessageZh[code]
		}
		return WrongMessageEn[code]
	} else {
		return WrongMessageZh[code]
	}
}

func InvalidParametersError(lang ...string) ApiResponse {
	return responseError(WRONG_PARAM, "", lang[0])
}

func InternalServiceError(message string, lang ...string) ApiResponse {
	return responseError(FAIL, message, lang[0])
}

func ResponseError(code int32, message string, lang ...string) ApiResponse {
	return responseError(code, message, lang[0])
}

func ResponseSuccess(result interface{}, msg string, lang ...string) ApiResponse {
	if len(msg) == 0 {
		msg = getResponseMsgWithLang(SUCCESS, lang[0])
	}
	return responseOutput(SUCCESS, msg, result)
}

func responseError(code int32, message string, lang ...string) ApiResponse {
	var msg string
	if len(message) == 0 {
		msg = getResponseMsgWithLang(code, lang[0])
	}
	msg = message

	return responseOutput(code, msg, nil)
}

func responseOutput(code int32, message string, result interface{}) ApiResponse {
	if result == nil {
		result = ""
	}
	return ApiResponse{
		Code:    code,
		Message: message,
		Result:  result,
	}
}

func getResponseMsgWithLang(code int32, la ...string) string {
	lang := "zh"
	if la[0] != "" || len(la) >= 0 {
		lang = "en"
	}
	var msg string
	switch lang {
	case "zh":
		msg = WrongMessageZh[code]
	default:
		msg = WrongMessageEn[code]
	}
	return msg
}
