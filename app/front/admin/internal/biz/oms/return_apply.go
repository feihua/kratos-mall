package oms

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type ReturnApplyListReq struct {
	Current  int64
	PageSize int64
}

type ReturnApply struct {
	Id               int64
	OrderId          int64  // 订单id
	CompanyAddressId int64  // 收货地址表id
	ProductId        int64  // 退货商品id
	OrderSn          string // 订单编号
	CreateTime       string // 申请时间
	MemberUsername   string // 会员用户名
	ReturnAmount     string // 退款金额
	ReturnName       string // 退货人姓名
	ReturnPhone      string // 退货人电话
	Status           int    // 申请状态：0->待处理；1->退货中；2->已完成；3->已拒绝
	HandleTime       string // 处理时间
	ProductPic       string // 商品图片
	ProductName      string // 商品名称
	ProductBrand     string // 商品品牌
	ProductAttr      string // 商品销售属性：颜色：红色；尺码：xl;
	ProductCount     int    // 退货数量
	ProductPrice     string // 商品单价
	ProductRealPrice string // 商品实际支付单价
	Reason           string // 原因
	Description      string // 描述
	ProofPics        string // 凭证图片，以逗号隔开
	HandleNote       string // 处理备注
	HandleMan        string // 处理人员
	ReceiveMan       string // 收货人
	ReceiveTime      string // 收货时间
	ReceiveNote      string // 收货备注
}

type ReturnApplyRepo interface {
	CreateReturnApply(context.Context, *ReturnApply) error
	GetReturnApply(ctx context.Context, id int64) error
	UpdateReturnApply(context.Context, *ReturnApply) error
	ListReturnApply(ctx context.Context, req *ReturnApplyListReq) ([]*ReturnApply, error)
	DeleteReturnApply(ctx context.Context, id int64) error
}

type ReturnApplyUseCase struct {
	repo ReturnApplyRepo
	log  *log.Helper
}

func NewReturnApplyUseCase(repo ReturnApplyRepo, logger log.Logger) *ReturnApplyUseCase {
	return &ReturnApplyUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (r *ReturnApplyUseCase) CreateReturnApply(ctx context.Context, user *ReturnApply) error {
	panic("implement me")
}

func (r *ReturnApplyUseCase) GetReturnApply(ctx context.Context, id int64) error {
	panic("implement me")
}

func (r *ReturnApplyUseCase) UpdateReturnApply(ctx context.Context, user *ReturnApply) error {
	panic("implement me")
}

func (r *ReturnApplyUseCase) ListReturnApply(ctx context.Context, pageNum, pageSize int64) ([]*ReturnApply, error) {
	panic("implement me")
}

func (r *ReturnApplyUseCase) DeleteReturnApply(ctx context.Context, id int64) error {
	panic("implement me")
}
