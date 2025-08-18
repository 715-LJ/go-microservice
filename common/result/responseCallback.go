package result

import (
	"go-microservice/common/xerr"
	"net/http"
)

func JwtUnauthorizedCallback(w http.ResponseWriter, r *http.Request, err error) {
	HttpResult(r, w, nil, err)
}

func UnsignedCallback(w http.ResponseWriter, r *http.Request, next http.Handler, strict bool, code int) {
	HttpResult(r, w, nil, xerr.NewErrCode(xerr.UNAUTHORIZED_ERROR))
}
