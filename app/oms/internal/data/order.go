package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/oms/internal/biz"
	"kratos-mall/app/oms/internal/data/model"
	"kratos-mall/pkg/util/pagination"
)

type orderRepo struct {
	data *Data
	log  *log.Helper
}

func NewOrderRepo(data *Data, logger log.Logger) biz.OrderRepo {
	return &orderRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (o orderRepo) CreateOrder(ctx context.Context, order *biz.Order) error {
	panic("implement me")
}

func (o orderRepo) GetOrder(ctx context.Context, id int64) error {
	panic("implement me")
}

func (o orderRepo) UpdateOrder(ctx context.Context, order *biz.Order) error {
	panic("implement me")
}

func (o orderRepo) ListOrder(ctx context.Context, req *biz.OrderListReq) ([]*biz.Order, error) {
	var list []model.OmsOrder
	result := o.data.db.WithContext(ctx).
		Limit(int(req.PageSize)).
		Offset(int(pagination.GetPageOffset(req.Current, req.PageSize))).
		Find(&list)

	if result.Error != nil {
		return nil, result.Error
	}

	//rv := make([]*biz.User, 0)
	//for _, item := range list {
	//	rv = append(rv, &biz.User{
	//		Id:             item.Id,
	//		Name:           item.Name,
	//		NickName:       item.NickName,
	//		Avatar:         item.Avatar,
	//		Password:       item.Password,
	//		Salt:           item.Salt,
	//		Email:          item.Email,
	//		Mobile:         item.Mobile,
	//		Status:         item.Status,
	//		DeptId:         item.DeptId,
	//		CreateBy:       item.CreateBy,
	//		CreateTime:     item.CreateTime,
	//		LastUpdateBy:   item.LastUpdateBy,
	//		LastUpdateTime: item.LastUpdateTime,
	//		DelFlag:        item.DelFlag,
	//		JobId:          item.JobId,
	//	})
	//}
	//return rv, nil

	fmt.Printf("%v", list)
	return nil, nil
}

func (o orderRepo) DeleteOrder(ctx context.Context, id int64) error {
	panic("implement me")
}
