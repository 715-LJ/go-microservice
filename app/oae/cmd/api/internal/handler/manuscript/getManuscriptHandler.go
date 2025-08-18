package manuscript

import (
	"net/http"

	"go-microservice/common/result"

	"go-microservice/app/oae/cmd/api/internal/logic/manuscript"
	"go-microservice/app/oae/cmd/api/internal/svc"
)

func GetManuscriptHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := manuscript.NewGetManuscriptLogic(r.Context(), svcCtx)
		resp, err := l.GetManuscript()

		result.HttpResult(r, w, resp, err)
	}
}
