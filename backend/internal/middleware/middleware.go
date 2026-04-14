package middleware

type (
	sMiddleware struct{}
)

func New() *sMiddleware {
	return &sMiddleware{}
}
