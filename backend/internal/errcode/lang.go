package errcode

// 错误码映射表，支持中英文
var errorMessages = map[int]map[string]string{
	// 成功状态码
	CodeSuccess: {
		"zh-CN": "成功",
		"en-US": "success",
	},

	// 通用错误码 (10xxx)
	CodeInternalServerError: {
		"zh-CN": "服务器内部错误",
		"en-US": "Internal erver error",
	},
	CodeInvalidParameter: {
		"zh-CN": "参数无效",
		"en-US": "Invalid parameter",
	},
	CodeDownloadFailed: {
		"zh-CN": "下载失败，请确认下载链接是否有效",
		"en-US": "Download failed, please confirm whether the download link is valid",
	},
	CodeMakeDirFailed: {
		"zh-CN": "创建目录失败",
		"en-US": "Make dir failed",
	},
	CodeServiceBusy: {
		"zh-CN": "服务繁忙，请稍后重试",
		"en-US": "Service busy, please try again later",
	},
	CodeTimeout: {
		"zh-CN": "服务超时",
		"en-US": "Service timeout",
	},
	CodeInvalidAPIKey: {
		"zh-CN": "API密钥无效",
		"en-US": "Invalid apiKey",
	},
	CodeInsufficientPoints: {
		"zh-CN": "积分不足，请充值",
		"en-US": "Insufficient credits. please recharge",
	},
	CodeNotPermission: {
		"zh-CN": "没有权限",
		"en-US": "No permission",
	},
	CodeUnknown: {
		"zh-CN": "未知错误",
		"en-US": "unknown error",
	},

	// 用户认证与订单模块错误码 (12xxx)
	CodeInvalidCaptcha: {
		"zh-CN": "验证码无效",
		"en-US": "Invalid captcha",
	},
	CodeCaptchaNotRequested: {
		"zh-CN": "未请求验证码",
		"en-US": "Captcha not requested",
	},
	CodeInvalidPhoneNumber: {
		"zh-CN": "手机号码无效",
		"en-US": "Invalid phone number",
	},
	CodeInvalidPassword: {
		"zh-CN": "密码无效",
		"en-US": "Invalid password",
	},
	CodeCreateQRCodeFailed: {
		"zh-CN": "创建二维码失败",
		"en-US": "Failed to create QR code",
	},
	CodeInvalidTicket: {
		"zh-CN": "票据无效",
		"en-US": "Invalid ticket",
	},
	CodeQRCodeLoginPending: {
		"zh-CN": "二维码登录待处理，请等待用户扫码",
		"en-US": "QR code login pending, waiting for user scan",
	},
	CodeQRCodeLoginFailed: {
		"zh-CN": "二维码登录失败",
		"en-US": "QR code login failed",
	},
	CodeUserNotFound: {
		"zh-CN": "用户未找到",
		"en-US": "User not found",
	},
	CodeUnauthorized: {
		"zh-CN": "用户未授权，请重新登录",
		"en-US": "User unauthorized, please login again",
	},
	CodeInvalidPayMethod: {
		"zh-CN": "支付方式无效",
		"en-US": "Invalid payment method",
	},
	CodeInvalidProduct: {
		"zh-CN": "商品无效",
		"en-US": "Invalid product",
	},
	CodeCreateOrderFailed: {
		"zh-CN": "创建订单失败",
		"en-US": "Failed to create order",
	},
	CodeOrderNotFound: {
		"zh-CN": "订单未找到",
		"en-US": "Order not found",
	},
	CodeQueryOrderFailed: {
		"zh-CN": "查询订单失败",
		"en-US": "Failed to query order",
	},
	CodeDuplicateUserName: {
		"zh-CN": "用户名重复，请换个名称重试",
		"en-US": "Username duplicate, please try another name",
	},

	// 通知模块错误码（15xxx）
	CodeSendMailFailed: {
		"zh-CN": "发送邮件失败",
		"en-US": "Send mail failed",
	},

	// 文件上传模块错误码（16xxx）
	CodeUploadFileNotFound: {
		"zh-CN": "上传文件未找到",
		"en-US": "Upload file not found",
	},
	CodeUploadFileFormatNotSupported: {
		"zh-CN": "上传文件格式不支持",
		"en-US": "Upload file format not supported",
	},

}
