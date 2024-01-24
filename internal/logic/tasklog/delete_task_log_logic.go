package tasklog

import (
	"context"

	"github.com/iot-synergy/synergy-job/ent/tasklog"
	"github.com/iot-synergy/synergy-job/internal/svc"
	"github.com/iot-synergy/synergy-job/internal/utils/dberrorhandler"
	"github.com/iot-synergy/synergy-job/types/job"

	"github.com/iot-synergy/synergy-common/i18n"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteTaskLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteTaskLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTaskLogLogic {
	return &DeleteTaskLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteTaskLogLogic) DeleteTaskLog(in *job.IDsReq) (*job.BaseResp, error) {
	_, err := l.svcCtx.DB.TaskLog.Delete().Where(tasklog.IDIn(in.Ids...)).Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &job.BaseResp{Msg: i18n.DeleteSuccess}, nil
}
