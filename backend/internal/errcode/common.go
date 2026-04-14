/**
 * @Author: thm
 * @Date: 2024/6/10 11:27
 * @Description: 错误码定义，用于返回给客户端的错误信息，前2位表示模块名称，后3位表示错误码，如：10001表示用户模块的第一个错误码，10002表示用户模块的第二个错误码，以此类推
 * @Version: 1.0
 */
package errcode

const (
	CodeSuccess             = 0
	CodeInternalServerError = 10001
	CodeInvalidParameter    = 10002
	CodeDownloadFailed      = 10003
	CodeMakeDirFailed       = 10004
	CodeServiceBusy         = 10005
	CodeTimeout             = 10006
	CodeInvalidAPIKey       = 10007 // 这个值不能修改，影响太大
	CodeInsufficientPoints  = 10008
	CodeNotPermission       = 10009
	CodeUnknown             = 10100
)
