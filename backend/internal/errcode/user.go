/**
 * @Author: thm
 * @Date: 2024/6/10 11:27
 * @Description: 用户模块错误码定义，模块编号为12
 * @Version: 1.0
 */
package errcode

const (
	CodeInvalidCaptcha      = 12001
	CodeCaptchaNotRequested = 12002
	CodeInvalidPhoneNumber  = 12003
	CodeInvalidPassword     = 12004
	CodeCreateQRCodeFailed  = 12005
	CodeInvalidTicket       = 12006
	CodeQRCodeLoginPending  = 12007
	CodeQRCodeLoginFailed   = 12008
	CodeUserNotFound        = 12009
	CodeUnauthorized        = 12010
	CodeInvalidPayMethod    = 12011
	CodeInvalidProduct      = 12012
	CodeCreateOrderFailed   = 12013
	CodeOrderNotFound       = 12014
	CodeQueryOrderFailed    = 12015
	CodeDuplicateUserName   = 12016
)
