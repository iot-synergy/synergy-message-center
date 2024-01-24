package emaillog

import (
	"context"

	"github.com/iot-synergy/synergy-message-center/internal/svc"
	"github.com/iot-synergy/synergy-message-center/internal/utils/dberrorhandler"
	"github.com/iot-synergy/synergy-message-center/types/mcms"

	"github.com/iot-synergy/synergy-common/i18n"

	"github.com/iot-synergy/synergy-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateEmailLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateEmailLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateEmailLogLogic {
	return &CreateEmailLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateEmailLogLogic) CreateEmailLog(in *mcms.EmailLogInfo) (*mcms.BaseUUIDResp, error) {
	query := l.svcCtx.DB.EmailLog.Create().
		SetNotNilTarget(in.Target).
		SetNotNilSubject(in.Subject).
		SetNotNilContent(in.Content)

	if in.SendStatus != nil {
		query.SetNotNilSendStatus(pointy.GetPointer(uint8(*in.SendStatus)))
	}

	result, err := query.Save(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &mcms.BaseUUIDResp{Id: result.ID.String(), Msg: i18n.CreateSuccess}, nil
}
