package result

import (
	"context"
	"github.com/zeromicro/go-zero/core/trace"
)

type ResponseSuccessBean struct {
	Code      uint32      `json:"code"`
	Message   string      `json:"message"`
	RequestId string      `json:"request_id"`
	Data      interface{} `json:"data"`
}
type NullJson struct{}

func Success(ctx context.Context, data interface{}) *ResponseSuccessBean {
	return &ResponseSuccessBean{200, "OK", trace.TraceIDFromContext(ctx), data}
}

type ResponseErrorBean struct {
	Code      uint32 `json:"code"`
	Message   string `json:"message"`
	RequestId string `json:"request_id"`
}

func Error(ctx context.Context, errCode uint32, errMsg string) *ResponseErrorBean {
	return &ResponseErrorBean{errCode, errMsg, trace.TraceIDFromContext(ctx)}
}
