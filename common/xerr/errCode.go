package xerr

// OK 成功返回
const OK uint32 = 200

/**(前3位代表业务,后三位代表具体功能)**/

// 全局错误码
const SERVER_COMMON_ERROR uint32 = 500
const REUQEST_PARAM_ERROR uint32 = 400
const TOKEN_EXPIRE_ERROR uint32 = 401
const UNAUTHORIZED_ERROR uint32 = 403

// 用户模块
const TOKEN_GENERATE_ERROR uint32 = 100004
