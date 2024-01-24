package smslog

import (
	"context"

	"github.com/iot-synergy/synergy-message-center/ent/smslog"
	"github.com/iot-synergy/synergy-message-center/internal/svc"
	"github.com/iot-synergy/synergy-message-center/internal/utils/dberrorhandler"
	"github.com/iot-synergy/synergy-message-center/types/mcms"

	"github.com/iot-synergy/synergy-common/i18n"
	"github.com/iot-synergy/synergy-common/utils/uuidx"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSmsLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteSmsLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSmsLogLogic {
	return &DeleteSmsLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteSmsLogLogic) DeleteSmsLog(in *mcms.UUIDsReq) (*mcms.BaseResp, error) {
	_, err := l.svcCtx.DB.SmsLog.Delete().Where(smslog.IDIn(uuidx.ParseUUIDSlice(in.Ids)...)).Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &mcms.BaseResp{Msg: i18n.DeleteSuccess}, nil
}
