package logic

import (
	"context"
	"fmt"
	"github.com/superyyk/mayi_rpc/app/upimg/api/internal/svc"
	"github.com/superyyk/mayi_rpc/app/upimg/api/internal/types"
	"github.com/superyyk/mayi_rpc/common"
	"github.com/superyyk/mayi_rpc/model"
	"github.com/superyyk/mayi_rpc/tool"
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpFileLogic {
	return &UpFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpFileLogic) UpFile(req *types.UpReq, r *http.Request) (resp *types.Resp, err error) {
	// todo: add your logic here and delete this line
	_, FileHeader, err := r.FormFile("file")
	res := make(map[string]interface{})

	//2. 获取文件后缀名
	extName := path.Ext(FileHeader.Filename)
	allowExtMap := map[string]bool{
		".jpg":  true,
		".png":  true,
		".gif":  true,
		".jpeg": true,
		".webp": true,
	} //通过map指定能接受的文件的后缀名
	if !allowExtMap[extName] {
		res["src"] = ""
		return &types.Resp{
			Res:  "400",
			Msg:  "不是有效图片文件",
			Data: res,
		}, nil
	}
	//3. 创建图片保存的目录
	day := tool.GetDay()
	dir := "./static/upload/" + day
	err = os.MkdirAll(dir, 0666)
	if err != nil {
		fmt.Println(err)
		//c.String(http.StatusOK, "MkdirAll失败")
		res["src"] = ""
		return &types.Resp{
			Res:  "400",
			Msg:  "存储失败",
			Data: res,
		}, nil
	}
	//4. 生成文件名称和文件保存的目录
	fileName := strconv.FormatInt(tool.GetUnix(), 10) + extName
	//5. 执行上传
	dst := path.Join(dir, fileName)
	var t *tool.Img
	err = t.SaveUploadedFile(FileHeader, dst)
	src := common.Base.BaseUrl + tool.Int_string(common.Base.Port) + "/" + dst
	res["src"] = src
	uid, _ := tool.GetUuid(12)
	img := &model.Images{
		Userid: req.UserUid,
		Uid:    uid,
		Src:    src,
		Time:   tool.GetUnix(),
	}
	if err = l.svcCtx.Db.Table("images").Create(&img).Error; err != nil {
		return &types.Resp{
			Res:  "400",
			Msg:  "存库失败",
			Data: res,
		}, nil
	}
	return &types.Resp{
		Res:  "200",
		Msg:  "存储成功",
		Data: res,
	}, nil
}
