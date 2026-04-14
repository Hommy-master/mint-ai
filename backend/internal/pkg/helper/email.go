package helper

import (
	"bytes"
	"context"
	"cozeos/internal/consts"
	"cozeos/util"
	"fmt"
	"html/template"
	"strings"

	"github.com/gogf/gf/v2/os/glog"
)

const paymentSuccessEmailTemplate = `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>新订单支付成功通知</title>
    <style type="text/css">
        /* 内联样式 + 嵌入式样式组合 */
        .ExternalClass { width: 100%; }
        .ExternalClass, .ExternalClass p, .ExternalClass span, .ExternalClass font, .ExternalClass td, .ExternalClass div { line-height: 100%; }
        body { margin: 0; padding: 0; -webkit-text-size-adjust: 100%; -ms-text-size-adjust: 100%; font-family: Arial, sans-serif; line-height: 1.6; color: #333; }
        table { border-collapse: collapse; mso-table-lspace: 0pt; mso-table-rspace: 0pt; }
        img { border: 0; height: auto; line-height: 100%; outline: none; text-decoration: none; -ms-interpolation-mode: bicubic; }
        p { display: block; margin: 13px 0; }
        .long-text { word-break: break-all; overflow-wrap: break-word; }
        .summary-value { font-size: 18px; font-weight: 600; color: #2c3e50; }
        .highlight { color: #27ae60; font-weight: 700; font-size: 20px; }
        .info-label { font-weight: 600; color: #555; padding: 12px 0 12px 10px; vertical-align: top; }
        .info-value { color: #222; font-weight: 500; padding: 12px 10px; background: #f8f9fa; border-left: 3px solid #3498db; }
    </style>
    <!--[if mso]>
    <style type="text/css">
        .outlook-table { width: 100% !important; }
        .summary-value { font-size: 16px !important; }
    </style>
    <![endif]-->
</head>
<body style="margin: 0; padding: 0; background-color: #f0f2f5;">
    <!-- 外层容器 -->
    <table width="100%" cellpadding="0" cellspacing="0" border="0" bgcolor="#f0f2f5">
        <tr>
            <td align="center" valign="top" style="padding: 20px 10px;">
                <!-- 内容容器 -->
                <table class="outlook-table" width="100%" cellpadding="0" cellspacing="0" border="0" bgcolor="#ffffff" style="max-width: 680px; border-radius: 10px; box-shadow: 0 4px 15px rgba(0, 0, 0, 0.1);">
                    <!-- 头部 -->
                    <tr>
                        <td bgcolor="#2c3e50" style="padding: 25px 20px; text-align: center; border-bottom: 5px solid #3498db;">
                            <h1 style="font-size: 22px; font-weight: 600; margin-bottom: 5px; color: white;">新订单支付成功通知</h1>
                            <p style="opacity: 0.9; font-size: 15px; color: #ecf0f1; margin: 0;">业务系统自动通知 - 请查阅订单详情</p>
                        </td>
                    </tr>
                    
                    <!-- 摘要区域 -->
                    <tr>
                        <td bgcolor="#f8f9fa" style="padding: 20px 15px; border-bottom: 1px solid #eee;">
                            <table width="100%" cellpadding="0" cellspacing="0" border="0">
                                <tr>
                                    <td width="100%" align="center" valign="top" style="padding: 10px 0;">
                                        <p style="font-size: 13px; color: #7f8c8d; margin: 0 0 5px 0;">订单金额</p>
                                        <p class="highlight" style="margin: 0; font-size: 20px; font-weight: 700; color: #27ae60;">{{ .Amount }} 元</p>
                                    </td>
                                </tr>
                                <tr>
                                    <td width="100%" align="center" valign="top" style="padding: 15px 0 10px;">
                                        <p style="font-size: 13px; color: #7f8c8d; margin: 0 0 5px 0;">订单号</p>
                                        <p class="long-text summary-value" style="margin: 0; padding: 0 10px; word-break: break-all; overflow-wrap: break-word;">{{ .OutTradeNo }}</p>
                                    </td>
                                </tr>
                                <tr>
                                    <td width="100%" align="center" valign="top" style="padding: 10px 0;">
                                        <p style="font-size: 13px; color: #7f8c8d; margin: 0 0 5px 0;">支付时间</p>
                                        <p class="summary-value" style="margin: 0;">{{ .SuccessTime }}</p>
                                    </td>
                                </tr>
                            </table>
                        </td>
                    </tr>
                    
                    <!-- 内容区域 -->
                    <tr>
                        <td style="padding: 25px 20px;">
                            <h2 style="font-size: 18px; color: #2c3e50; padding-bottom: 10px; border-bottom: 2px solid #3498db; margin-bottom: 20px; font-weight: 600;">订单详细信息</h2>
                            
                            <!-- 订单信息表格 -->
                            <table width="100%" cellpadding="0" cellspacing="0" border="0" style="margin-bottom: 25px;">
                                <tr>
                                    <td width="30%" class="info-label" style="font-weight: 600; color: #555; padding: 12px 0 12px 10px; vertical-align: top;">用户ID</td>
                                    <td class="info-value long-text" style="color: #222; font-weight: 500; padding: 12px 10px; background: #f8f9fa; border-left: 3px solid #3498db; word-break: break-all;">{{ .UserID }}</td>
                                </tr>
                                <tr>
                                    <td width="30%" class="info-label" style="font-weight: 600; color: #555; padding: 12px 0 12px 10px; vertical-align: top;">交易ID</td>
                                    <td class="info-value long-text" style="color: #222; font-weight: 500; padding: 12px 10px; background: #f8f9fa; border-left: 3px solid #3498db; word-break: break-all;">{{ .TransactionID }}</td>
                                </tr>
                                <tr>
                                    <td width="30%" class="info-label" style="font-weight: 600; color: #555; padding: 12px 0 12px 10px; vertical-align: top;">支付状态</td>
                                    <td class="info-value" style="color: #222; font-weight: 500; padding: 12px 10px; background: #f8f9fa; border-left: 3px solid #3498db;">支付成功</td>
                                </tr>
                            </table>
                            
                            <!-- 建议操作区域 -->
                            <!-- 建议操作区域 -->
                            <table width="100%" cellpadding="0" cellspacing="0" border="0" bgcolor="#f8f9fa" style="margin-top: 30px; padding: 20px; border-radius: 8px; border: 1px dashed #3498db;">
                                <tr>
                                    <td>
                                        <h3 style="color: #2c3e50; margin-bottom: 15px; font-size: 16px;">建议操作：</h3>
                                        <ul style="list-style-type: none; margin: 0; padding: 0;">
                                            <li style="margin-bottom: 10px; padding-left: 25px; position: relative;">在管理后台查看完整订单信息</li>
                                            <li style="margin-bottom: 10px; padding-left: 25px; position: relative;">检查库存状态并准备发货</li>
                                            <li style="margin-bottom: 10px; padding-left: 25px; position: relative;">审核支付信息与金额是否匹配</li>
                                            <li style="margin-bottom: 10px; padding-left: 25px; position: relative;">关注该用户的后续订单情况</li>
                                        </ul>
                                    </td>
                                </tr>
                            </table>
                        </td>
                    </tr>
                    
                    <!-- 页脚 -->
                    <tr>
                        <td bgcolor="#f8f9fa" style="text-align: center; padding: 20px; color: #7f8c8d; font-size: 13px; border-top: 1px solid #eee;">
                            <p style="margin: 0;">© 2025 逗赛科技-简创AIGC | 用户支付订单通知</p>
                        </td>
                    </tr>
                </table>
            </td>
        </tr>
    </table>
</body>
</html>`

// 预编译模板（全局变量）
var paymentSuccessTmpl = template.Must(template.New("paymentSuccessEmail").Parse(paymentSuccessEmailTemplate))

const botCustomEmailTemplate = `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>新智能体定制通知</title>
    <style type="text/css">
        /* 内联样式 + 嵌入式样式组合 */
        .ExternalClass { width: 100%; }
        .ExternalClass, .ExternalClass p, .ExternalClass span, .ExternalClass font, .ExternalClass td, .ExternalClass div { line-height: 100%; }
        body { margin: 0; padding: 0; -webkit-text-size-adjust: 100%; -ms-text-size-adjust: 100%; font-family: Arial, sans-serif; line-height: 1.6; color: #333; }
        table { border-collapse: collapse; mso-table-lspace: 0pt; mso-table-rspace: 0pt; }
        img { border: 0; height: auto; line-height: 100%; outline: none; text-decoration: none; -ms-interpolation-mode: bicubic; }
        p { display: block; margin: 13px 0; }
        .long-text { word-break: break-all; overflow-wrap: break-word; }
        .summary-value { font-size: 18px; font-weight: 600; color: #2c3e50; }
        .highlight { color: #27ae60; font-weight: 700; font-size: 20px; }
        .info-label { font-weight: 600; color: #555; padding: 12px 0 12px 10px; vertical-align: top; }
        .info-value { color: #222; font-weight: 500; padding: 12px 10px; background: #f8f9fa; border-left: 3px solid #3498db; }
    </style>
    <!--[if mso]>
    <style type="text/css">
        .outlook-table { width: 100% !important; }
        .summary-value { font-size: 16px !important; }
    </style>
    <![endif]-->
</head>
<body style="margin: 0; padding: 0; background-color: #f0f2f5;">
    <!-- 外层容器 -->
    <table width="100%" cellpadding="0" cellspacing="0" border="0" bgcolor="#f0f2f5">
        <tr>
            <td align="center" valign="top" style="padding: 20px 10px;">
                <!-- 内容容器 -->
                <table class="outlook-table" width="100%" cellpadding="0" cellspacing="0" border="0" bgcolor="#ffffff" style="max-width: 680px; border-radius: 10px; box-shadow: 0 4px 15px rgba(0, 0, 0, 0.1);">
                    <!-- 头部 -->
                    <tr>
                        <td bgcolor="#2c3e50" style="padding: 25px 20px; text-align: center; border-bottom: 5px solid #3498db;">
                            <h1 style="font-size: 22px; font-weight: 600; margin-bottom: 5px; color: white;">新智能体定制通知</h1>
                            <p style="opacity: 0.9; font-size: 15px; color: #ecf0f1; margin: 0;">业务系统自动通知 - 请查阅定制详情</p>
                        </td>
                    </tr>
                    
                    <!-- 摘要区域 -->
                    <tr>
                        <td bgcolor="#f8f9fa" style="padding: 20px 15px; border-bottom: 1px solid #eee;">
                            <table width="100%" cellpadding="0" cellspacing="0" border="0">
                                <tr>
                                    <td width="100%" align="center" valign="top" style="padding: 10px 0;">
                                        <p style="font-size: 13px; color: #7f8c8d; margin: 0 0 5px 0;">预算范围</p>
                                        <p class="highlight" style="margin: 0; font-size: 20px; font-weight: 700; color: #27ae60;">{{ .Budget }}</p>
                                    </td>
                                </tr>
                            </table>
                        </td>
                    </tr>
                    
                    <!-- 内容区域 -->
                    <tr>
                        <td style="padding: 25px 20px;">
                            <h2 style="font-size: 18px; color: #2c3e50; padding-bottom: 10px; border-bottom: 2px solid #3498db; margin-bottom: 20px; font-weight: 600;">定制详细信息</h2>
                            
                            <!-- 定制信息表格 -->
                            <table width="100%" cellpadding="0" cellspacing="0" border="0" style="margin-bottom: 25px;">
                                <tr>
                                    <td width="30%" class="info-label" style="font-weight: 600; color: #555; padding: 12px 0 12px 10px; vertical-align: top;">联系方式类型</td>
                                    <td class="info-value long-text" style="color: #222; font-weight: 500; padding: 12px 10px; background: #f8f9fa; border-left: 3px solid #3498db; word-break: break-all;">{{ .ContactType }}</td>
                                </tr>
                                <tr>
                                    <td width="30%" class="info-label" style="font-weight: 600; color: #555; padding: 12px 0 12px 10px; vertical-align: top;">联系方式</td>
                                    <td class="info-value long-text" style="color: #222; font-weight: 500; padding: 12px 10px; background: #f8f9fa; border-left: 3px solid #3498db; word-break: break-all;">{{ .ContactValue }}</td>
                                </tr>
                                <tr>
                                    <td width="30%" class="info-label" style="font-weight: 600; color: #555; padding: 12px 0 12px 10px; vertical-align: top;">参考视频</td>
                                    <td class="info-value long-text" style="color: #222; font-weight: 500; padding: 12px 10px; background: #f8f9fa; border-left: 3px solid #3498db; word-break: break-all;">{{ .VideoLink }}</td>
                                </tr>
                            </table>
                            
                            <!-- 建议操作区域 -->
                            <table width="100%" cellpadding="0" cellspacing="0" border="0" bgcolor="#f8f9fa" style="margin-top: 30px; padding: 20px; border-radius: 8px; border: 1px dashed #3498db;">
                                <tr>
                                    <td>
                                        <h3 style="color: #2c3e50; margin-bottom: 15px; font-size: 16px;">建议操作：</h3>
                                        <ul style="list-style-type: none; margin: 0; padding: 0;">
                                            <li style="margin-bottom: 10px; padding-left: 25px; position: relative;">及时联系客户确认定制需求</li>
                                            <li style="margin-bottom: 10px; padding-left: 25px; position: relative;">查看参考视频了解客户需求</li>
                                            <li style="margin-bottom: 10px; padding-left: 25px; position: relative;">安排项目团队评估开发周期</li>
                                            <li style="margin-bottom: 10px; padding-left: 25px; position: relative;">跟进定制进度并与客户保持沟通</li>
                                        </ul>
                                    </td>
                                </tr>
                            </table>
                        </td>
                    </tr>
                    
                    <!-- 页脚 -->
                    <tr>
                        <td bgcolor="#f8f9fa" style="text-align: center; padding: 20px; color: #7f8c8d; font-size: 13px; border-top: 1px solid #eee;">
                            <p style="margin: 0;">© 2025 逗赛科技-简创AIGC | 智能体定制通知</p>
                        </td>
                    </tr>
                </table>
            </td>
        </tr>
    </table>
</body>
</html>`

// 预编译模板（全局变量）
var botCustomTmpl = template.Must(template.New("botCustomEmail").Parse(botCustomEmailTemplate))

func init() {
	// 模板字段校验
	required := []string{"UserID", "OutTradeNo", "Amount", "TransactionID", "SuccessTime"}
	for _, field := range required {
		if !strings.Contains(paymentSuccessEmailTemplate, field) {
			panic("Email template missing required field: " + field)
		}
	}
	// 验证botCustomEmailTemplate的字段
	requiredBotCustom := []string{"ContactType", "ContactValue", "VideoLink", "Budget"}
	for _, field := range requiredBotCustom {
		if !strings.Contains(botCustomEmailTemplate, field) {
			panic("BotCustomEmail template missing required field: " + field)
		}
	}
}

func SendPaymentSuccessEmail(ctx context.Context, userID uint, outTradeNo string, amount float64, transactionID string, successTime string) error {
	var tpl bytes.Buffer
	err := paymentSuccessTmpl.Execute(&tpl, map[string]string{
		"UserID":        fmt.Sprintf("%d", userID),
		"OutTradeNo":    outTradeNo,
		"Amount":        fmt.Sprintf("%.2f", amount),
		"TransactionID": transactionID,
		"SuccessTime":   successTime,
	})
	if err != nil {
		glog.Warningf(ctx, "Failed to render payment success email template: %v", err)
		return err
	}

	err = util.SendMail(
		consts.SmtpHost,
		consts.SmtpPort,
		consts.SmtpUser,
		consts.SmtpPass,
		consts.SmtpUser,
		[]string{"taohongmin@sina.cn"},
		"订单支付成功通知",
		tpl.String(),
		util.HTML,
	)
	if err != nil {
		glog.Warningf(ctx, "Failed to send payment success email: %v", err)
	}

	return err
}

// SendBotCustomEmail 发送智能体定制通知邮件
func SendBotCustomEmail(ctx context.Context, contactType, contactValue, videoLink, budget string) error {
	var tpl bytes.Buffer
	err := botCustomTmpl.Execute(&tpl, map[string]string{
		"ContactType":  contactType,
		"ContactValue": contactValue,
		"VideoLink":    videoLink,
		"Budget":       budget,
	})
	if err != nil {
		glog.Warningf(ctx, "Failed to render bot custom email template: %v", err)
		return err
	}

	err = util.SendMail(
		consts.SmtpHost,
		consts.SmtpPort,
		consts.SmtpUser,
		consts.SmtpPass,
		consts.SmtpUser,
		[]string{"taohongmin@sina.cn"},
		"新智能体定制通知",
		tpl.String(),
		util.HTML,
	)
	if err != nil {
		glog.Warningf(ctx, "Failed to send bot custom email: %v", err)
	}

	return err
}
