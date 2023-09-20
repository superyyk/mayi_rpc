package logic

import (
	"context"
	"time"

	"github.com/superyyk/mayi_rpc/app/kehu/api/internal/svc"
	"github.com/superyyk/mayi_rpc/app/kehu/api/internal/types"
	"github.com/superyyk/mayi_rpc/model"
	"github.com/superyyk/mayi_rpc/tool"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddKehuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddKehuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddKehuLogic {
	return &AddKehuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddKehuLogic) AddKehu(req *types.AddReq) (resp *types.Resp, err error) {
	// todo: add your logic here and delete this line
	//kehu := new(kehu_model.Kehu)
	uid, _ := tool.GetUuid(12)
	kehu := new(model.Kehu)
	kehu.Name = req.Name
	kehu.Age = req.Age
	kehu.Tel = req.Tel
	kehu.Puber = req.Puber
	kehu.Uid = uid
	kehu.Level = req.Level
	kehu.Time = time.Now().Unix()
	kehu.Tip = req.Tip
	//var kehu kehu_model.Kehu
	//l.svcCtx.KehuModel.Insert(l.ctx, kehu)
	if err := l.svcCtx.Db.Table("kehu").Create(&kehu).Error; err != nil {
		logx.Errorf("创建客户失败：", err)
		return &types.Resp{
			Res: "400",
			Msg: "添加失败",
		}, nil
	}
	return &types.Resp{
		Res: "200",
		Msg: "添加成功",
	}, nil
}
