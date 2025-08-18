package {{.PkgName}}

import (
    "net/http"

    "go-microservice/common/result"
    {{if .HasRequest}}"go-microservice/common/translator"{{end}}
    {{if .HasRequest}}"github.com/zeromicro/go-zero/rest/httpx"{{end}}
    {{.ImportPackages}}
)

func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        {{if .HasRequest}}var req types.{{.RequestType}}
        if err := httpx.Parse(r, &req); err != nil {
            result.ParamErrorResult(r, w, err)
            return
        }

        validateErr := translator.Validate(&req)
        if validateErr != nil {
            result.ParamErrorResult(r, w, validateErr)
            return
        }

        {{end}}l := {{.LogicName}}.New{{.LogicType}}(r.Context(), svcCtx)
        {{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})

        result.HttpResult(r, w, {{if .HasResp}}resp{{else}}nil{{end}}, err)
    }
}