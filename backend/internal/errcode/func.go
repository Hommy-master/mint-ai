package errcode

// 获取本地化错误消息
func LocalizedMessage(code int, lang string) string {
	if messages, ok := errorMessages[code]; ok {
		if msg, ok := messages[lang]; ok {
			return msg
		}

		// 默认返回英文
		return messages["en-US"]
	}

	return "unknown"
}
