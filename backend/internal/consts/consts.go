package consts

type ContextKey string

const (
	CtxMaxDownloadFileSize ContextKey = "MaxDownloadFileSize" // 最大下载文件限制
)

const (
	CtxAccessUserKey     = "user"
	CtxAccessUserMailKey = "mail"

	WebRootPrefix = "/app"
)

const (
	CliLoggerName = "cli"
)
