package validator

import "github.com/gogf/gf/v2/util/gvalid"

func init() {
	// 注册自定义校验规则（避免名称冲突）
	gvalid.RegisterRule("valid-url", RuleValidURL)
}
