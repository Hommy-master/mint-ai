package validator

import (
	"context"
	"net/url"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gvalid"
)

// RuleValidURL 自定义URL校验实现
func RuleValidURL(ctx context.Context, in gvalid.RuleFuncInput) error {
	str := in.Value.String()

	// 空值检查（含字段名）
	if str == "" {
		return gerror.Newf("field `%s` cannot be empty", in.Field)
	}

	// 标准库解析URL
	u, err := url.Parse(str)
	if err != nil {
		return gerror.Newf("field `%s` has invalid format: %v", in.Field, err)
	}

	// 关键组件验证
	if u.Scheme == "" {
		return gerror.Newf("field `%s` is missing protocol (http/https)", in.Field)
	}

	if u.Host == "" {
		return gerror.Newf("field `%s` is missing hostname", in.Field)
	}

	// 协议白名单（可扩展）
	if u.Scheme != "http" && u.Scheme != "https" {
		return gerror.Newf("field `%s` only supports http/https protocols", in.Field)
	}

	return nil
}
