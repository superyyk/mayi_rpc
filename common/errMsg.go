package common

var message map[uint32]string

// 成功返回
const OK uint32 = 200

/**(前3位代表业务,后三位代表具体功能)**/

// 全局错误码
const SERVER_COMMON_ERROR uint32 = 100
const REUQEST_PARAM_ERROR uint32 = 101
const TOKEN_ERROR uint32 = 401
const TOKEN_GENERATE_ERROR uint32 = 402
const DB_ERROR uint32 = 404
const UPDATE_ERR uint32 = 405
const OPTION_ERR uint32 = 400
const CHANGE_ERR uint32 = 406
const SAVE_ERR uint32 = 407
const SHEARCH_ERR uint32 = 408
const CODE_ERR uint32 = 409
const USER_ERR uint32 = 410

// 用户模块
func init() {
	message = make(map[uint32]string)
	message[OK] = "成功！"
	message[SERVER_COMMON_ERROR] = "服务器开小差啦,稍后再来试一试"
	message[REUQEST_PARAM_ERROR] = "参数错误"
	message[TOKEN_ERROR] = "token失效，请重新登陆"
	message[TOKEN_GENERATE_ERROR] = "生成token失败"
	message[DB_ERROR] = "数据库繁忙,请稍后再试"
	message[UPDATE_ERR] = "更新数据影响行数为0"
	message[OPTION_ERR] = "操作失败"
	message[CHANGE_ERR] = "修改失败"
	message[SAVE_ERR] = "保存失败"
	message[SHEARCH_ERR] = "查询失败"
	message[USER_ERR] = "未找到该用户信息，用户名或密码错误！"
	message[CODE_ERR] = "验证码信息不符"

}

func GetErrMsg(errcode uint32) string {
	if msg, ok := message[errcode]; ok {
		return msg
	} else {
		return "服务器开小差啦,稍后再来试一试"
	}
}

func IsCodeErr(errcode uint32) bool {
	if _, ok := message[errcode]; ok {
		return true
	} else {
		return false
	}
}
