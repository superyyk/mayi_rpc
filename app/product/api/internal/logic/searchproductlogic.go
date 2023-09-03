package logic

import (
	"context"

	"mayi/app/product/api/internal/svc"
	"mayi/app/product/api/internal/types"
	"mayi/model"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
)

type SearchProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchProductLogic {
	return &SearchProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchProductLogic) SearchProduct(req *types.SearchReq) (resp *types.Resp, err error) {
	// todo: add your logic here and delete this line
	var products []model.Products
	var unit []model.Unit
	var guige []model.Guige
	var yuanliao []model.Yuanliao

	var count int
	//tx := l.svcCtx.Db
	mr.Finish(func() error {
		l.svcCtx.Db.Table("products").Where("puber=? AND status=?", req.Uid, 0).Order("id desc").Find(&products).Count(&count)
		if len(products) > 0 {
			for i := 0; i < len(products); i++ {
				var liucheng []model.Liucheng
				l.svcCtx.Db.Table("make_rule").Where("productid=? AND status=?", products[i].Uid, 0).Order("id asc").Find(&liucheng)
				products[i].Liucheng = liucheng
			}
		}
		return nil
	},
		func() error {
			l.svcCtx.Db.Table("units").Where("status=?", 0).Find(&unit)
			return nil
		}, func() error {
			l.svcCtx.Db.Table("guige").Where("puber=? AND status=?", req.Uid, 0).Order("id desc").Find(&guige)
			return nil
		},
		func() error {
			l.svcCtx.Db.Table("yuanliao").Where("puber=?  AND status=?", req.Uid, 0).Order("id desc").Find(&yuanliao)
			return nil
		},
	)

	res := make(map[string]interface{})
	res["count"] = count
	res["products"] = products
	res["unit"] = unit
	res["guige"] = guige
	res["yuanliao"] = yuanliao
	//.Commit()
	if len(products) > 0 {
		return &types.Resp{
			Res:  "200",
			Msg:  "success",
			Data: res,
		}, nil
	}

	return &types.Resp{
		Res:  "200",
		Msg:  "未上传，请上传产品",
		Data: res,
	}, nil
}
