package pms

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
	pmsV1 "kratos-mall/api/pms/v1"
	"kratos-mall/app/front/admin/internal/biz/pms"
	"kratos-mall/app/front/admin/internal/data"
)

type productRepo struct {
	data *data.Data
	log  *log.Helper
}

func NewProductRepo(data *data.Data, logger log.Logger) pms.ProductRepo {
	return &productRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (p productRepo) CreateProduct(ctx context.Context, product *pms.Product) error {
	panic("implement me")
}

func (p productRepo) GetProduct(ctx context.Context, id int64) error {
	panic("implement me")
}

func (p productRepo) UpdateProduct(ctx context.Context, product *pms.Product) error {
	panic("implement me")
}

func (p productRepo) ListProduct(ctx context.Context, req *pms.ProductListReq) (*pms.ProductListResp, error) {
	list, _ := p.data.PmsClient.ProductList(ctx, &pmsV1.ProductListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	orders := make([]*pms.Product, 0)
	copier.Copy(&orders, &list.List)

	return &pms.ProductListResp{
		Total: list.Total,
		List:  orders,
	}, nil
}

func (p productRepo) DeleteProduct(ctx context.Context, id int64) error {
	panic("implement me")
}
