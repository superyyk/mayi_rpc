package logic

import (
	"context"
	"time"

	"mayi/app/product/api/internal/svc"
	"mayi/app/product/api/internal/types"
	"mayi/model"
	"mayi/tool"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProductLogic {
	return &AddProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddProductLogic) AddProduct(req *types.AddReq) (resp *types.Resp, err error) {
	// todo: add your logic here and delete this line
	//res, err := l.svcCtx.UserModel.FindOne(l.ctx, tool.String_int64(req.Uid))
	uid, _ := tool.GetUuid(12)
	product := &model.Products{
		Name:      req.Name,
		Unit:      req.Unit,
		Guige:     req.Guige,
		Suply:     req.Suply,
		Pertime:   req.PerTime,
		Perhaocai: req.PerHaocai,
		Yuanliao:  req.Yuanliao,
		Price:     req.Price,
		Tip:       req.Tip,
		Imgs:      req.Imgs,
		Puber:     req.Puber,
		Time:      tool.Int64_string(time.Now().Unix()),
		Uid:       uid,
	}
	tx := l.svcCtx.Db.Begin()

	//检查计量单位是否存在
	var unit []model.Unit
	tx.Table("units").Where("name=?", req.Unit).First(&unit)
	if len(unit) == 0 {
		u := &model.Unit{
			Name: req.Unit,
		}
		if err := tx.Table("units").Create(&u).Error; err != nil {
			tx.Rollback()
			logx.Info("add units err:", err.Error)
		}
	}
	//检查规格是否存在
	var guige []model.Guige
	tx.Table("guige").Where("name=? AND puber=?", req.Guige, req.Puber).First(&guige)
	if len(guige) == 0 {
		u := &model.Guige{
			Name:  req.Guige,
			Puber: req.Uid,
		}
		if err := tx.Table("guige").Create(&u).Error; err != nil {
			tx.Rollback()
			logx.Info("add guige err:", err.Error)
		}
	}

	if err := tx.Table("products").Create(&product).Error; err != nil {
		tx.Rollback()
		return &types.Resp{
			Res:  "400",
			Msg:  "系统繁忙",
			Data: 400,
		}, nil
	}
	tx.Commit()
	return &types.Resp{
		Res:  "200",
		Msg:  "添加成功！",
		Data: 200,
	}, nil
}
