package manuscript

import (
	"context"
	"go-microservice/app/oae/cmd/api/internal/svc"
	"go-microservice/app/oae/cmd/api/internal/types"
	"go-microservice/common/logc"
	"math/rand"

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
	customLogger := &logc.CustomLogger{}
	customLogger.Error(l.ctx, "This is an error message", "additional data")
	customLogger.Info(l.ctx, "This is an info message", "additional data")

	return &types.ManuscriptInfo{
		Id:     int64(rand.Int()),
		Status: "OK",
		Title:  "oae-service:1002",
	}, nil
}
