package i

import (
	"context"
	"cozeos/internal/errcode"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

func T(ctx context.Context, code int) error {
	lang := ctx.Value("lang").(string)
	message := errcode.LocalizedMessage(code, lang)
	return gerror.NewCode(gcode.New(code, message, nil))
}
