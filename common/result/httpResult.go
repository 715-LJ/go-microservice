package result

import (
	"fmt"
	"net/http"

	"go-microservice/common/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
)

func HttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	//fmt.Printf("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa:%+v", err)
	if err == nil {
		//成功返回
		r := Success(r.Context(), resp)
		httpx.WriteJson(w, http.StatusOK, r)
	} else {
		//错误返回
		errCode := xerr.SERVER_COMMON_ERROR
		errMsg := fmt.Sprintf("%+v", err)

		causeErr := errors.Cause(err)                // err类型
		if e, ok := causeErr.(*xerr.CodeError); ok { //自定义错误类型
			errCode = e.GetErrCode()
			errMsg = e.GetErrMsg()
		} else {
			if gStatus, ok := status.FromError(causeErr); ok { // grpc err错误
				grpcCode := uint32(gStatus.Code())
				if xerr.IsCodeErr(grpcCode) { //区分自定义错误跟系统底层、db等错误，底层、db错误不能返回给前端
					errCode = grpcCode
					errMsg = gStatus.Message()
				}
			}
		}

		httpx.WriteJson(w, http.StatusBadRequest, Error(r.Context(), errCode, errMsg))
	}
}

func ParamErrorResult(r *http.Request, w http.ResponseWriter, err error) {
	httpx.WriteJson(w, http.StatusBadRequest, Error(r.Context(), xerr.REUQEST_PARAM_ERROR, err.Error()))
}
