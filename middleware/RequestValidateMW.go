package middleware

import (
	context "context"
	"github.com/cloudwego/kitex/pkg/endpoint"
)

func RequestValidateMW(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		return authToken(ctx, req, resp, next)
	}
}

func authToken(ctx context.Context, req, resp interface{}, next endpoint.Endpoint) error {
	// TODO parse token & set context
	//ctx = context.WithValue(ctx, "key", "value")
	err := next(ctx, req, resp)
	return err
}
