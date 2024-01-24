package task

import (
	"context"

	"github.com/iot-synergy/synergy-job/ent"
	"github.com/iot-synergy/synergy-job/ent/task"
	"github.com/iot-synergy/synergy-job/ent/tasklog"
	"github.com/iot-synergy/synergy-job/internal/svc"
	"github.com/iot-synergy/synergy-job/internal/utils/dberrorhandler"
	"github.com/iot-synergy/synergy-job/internal/utils/entx"
	"github.com/iot-synergy/synergy-job/types/job"

	"github.com/iot-synergy/synergy-common/i18n"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteTaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTaskLogic {
	return &DeleteTaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteTaskLogic) DeleteTask(in *job.IDsReq) (*job.BaseResp, error) {
	err := entx.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {
		_, err := tx.TaskLog.Delete().Where(tasklog.HasTasksWith(task.IDIn(in.Ids...))).Exec(l.ctx)
		if err != nil {
			return err
		}

		_, err = tx.Task.Delete().Where(task.IDIn(in.Ids...)).Exec(l.ctx)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &job.BaseResp{Msg: i18n.DeleteSuccess}, nil
}
