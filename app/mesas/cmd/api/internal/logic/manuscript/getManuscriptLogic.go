package manuscript

import (
	"context"
	"go-microservice/app/mesas/model"
	"go-microservice/common/xerr"
	"math/rand"

	"go-microservice/app/mesas/cmd/api/internal/svc"
	"go-microservice/app/mesas/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetManuscriptLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取稿件信息
func NewGetManuscriptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetManuscriptLogic {
	return &GetManuscriptLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetManuscriptLogic) GetManuscript() (resp *types.ManuscriptInfo, err error) {
	manuscript, err := model.NewManuscriptModel().GetItem(12)
	if err != nil {
		return nil, xerr.NewErrCodeMsg(xerr.SERVER_COMMON_ERROR, xerr.MapErrMsg(xerr.SERVER_COMMON_ERROR))
	}
	return &types.ManuscriptInfo{
		Id:     int64(rand.Int()),
		Status: "OK",
		Title:  "mesas-service:1001-" + manuscript.Title,
	}, nil
}
