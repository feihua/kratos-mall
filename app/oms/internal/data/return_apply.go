package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/oms/internal/biz"
	"kratos-mall/app/oms/internal/data/model"
	"kratos-mall/pkg/util/pagination"
)

type returnApplyRepo struct {
	data *Data
	log  *log.Helper
}

func NewReturnApplyRepo(data *Data, logger log.Logger) biz.ReturnApplyRepo {
	return &returnApplyRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r returnApplyRepo) CreateReturnApply(ctx context.Context, apply *biz.ReturnApply) error {
	panic("implement me")
}

func (r returnApplyRepo) GetReturnApply(ctx context.Context, id int64) error {
	panic("implement me")
}

func (r returnApplyRepo) UpdateReturnApply(ctx context.Context, apply *biz.ReturnApply) error {
	panic("implement me")
}

func (r returnApplyRepo) ListReturnApply(ctx context.Context, req *biz.ReturnApplyListReq) (*biz.ReturnApplyListResp, error) {
	var all []model.OmsOrderReturnApply
	result := r.data.db.WithContext(ctx).
		Limit(int(req.PageSize)).
		Offset(int(pagination.GetPageOffset(req.Current, req.PageSize))).
		Find(&all)

	var count int64
	r.data.db.WithContext(ctx).Model(&all).Count(&count)

	if result.Error != nil {
		return nil, result.Error
	}

	list := make([]*biz.ReturnApply, 0)

	for _, item := range all {
		list = append(list, &biz.ReturnApply{
			Id:               item.Id,
			OrderId:          item.OrderId,
			CompanyAddressId: item.CompanyAddressId,
			ProductId:        item.ProductId,
			OrderSn:          item.OrderSn,
			CreateTime:       item.CreateTime.Format("2006-01-02 15:04:05"),
			MemberUserName:   item.MemberUsername,
			ReturnAmount:     item.ReturnAmount,
			ReturnName:       item.ReturnName,
			ReturnPhone:      item.ReturnPhone,
			Status:           item.Status,
			HandleTime:       item.HandleTime.Format("2006-01-02 15:04:05"),
			ProductPic:       item.ProductPic,
			ProductName:      item.ProductName,
			ProductBrand:     item.ProductBrand,
			ProductAttr:      item.ProductAttr,
			ProductCount:     item.ProductCount,
			ProductPrice:     item.ProductPrice,
			ProductRealPrice: item.ProductRealPrice,
			Reason:           item.Reason,
			Description:      item.Description,
			ProofPics:        item.ProofPics,
			HandleNote:       item.HandleNote,
			HandleMan:        item.HandleMan,
			ReceiveMan:       item.ReceiveMan,
			ReceiveTime:      item.ReceiveTime.Format("2006-01-02 15:04:05"),
			ReceiveNote:      item.ReceiveNote,
		})
	}

	return &biz.ReturnApplyListResp{
		Total: count,
		List:  list,
	}, nil
}

func (r returnApplyRepo) DeleteReturnApply(ctx context.Context, id int64) error {
	panic("implement me")
}
