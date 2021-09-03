package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	sysV1 "kratos-mall/api/sys/v1"
	"kratos-mall/app/front/admin/internal/biz"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) Login(ctx context.Context, g *biz.Login) (*sysV1.LoginResp, error) {
	login, _ := r.data.sysClient.Login(ctx, &sysV1.LoginReq{
		UserName: g.UserName,
		Password: g.Password,
	})

	fmt.Printf("test%v", login)

	//return err
	return login, nil
}
