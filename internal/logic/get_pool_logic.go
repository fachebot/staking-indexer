package logic

import (
	"context"

	"staking-indexer/internal/svc"
	"staking-indexer/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPoolLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPoolLogic {
	return &GetPoolLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPoolLogic) GetPool(req *types.GetPoolRequest) (resp *types.GetPoolResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
