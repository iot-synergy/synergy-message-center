package smsprovider

import (
	"context"

	"github.com/iot-synergy/synergy-message-center/ent/smsprovider"
	"github.com/iot-synergy/synergy-message-center/internal/svc"
	"github.com/iot-synergy/synergy-message-center/internal/utils/dberrorhandler"
	"github.com/iot-synergy/synergy-message-center/types/mcms"

	"github.com/iot-synergy/synergy-common/i18n"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSmsProviderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteSmsProviderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSmsProviderLogic {
	return &DeleteSmsProviderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteSmsProviderLogic) DeleteSmsProvider(in *mcms.IDsReq) (*mcms.BaseResp, error) {
	_, err := l.svcCtx.DB.SmsProvider.Delete().Where(smsprovider.IDIn(in.Ids...)).Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &mcms.BaseResp{Msg: i18n.DeleteSuccess}, nil
}
