package data

import (
	"github.com/feihua/kratos-mall/app/sms/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewDB,
	NewGreeterRepo,
	NewCouponRepo,
	NewCouponHistoryRepo,
	NewFlashPromotionRepo,
	NewFlashPromotionHistoryRepo,
	NewFlashPromotionSessionRepo,
	NewHomeAdvertiseRepo,
	NewHomeBrandRepo,
	NewHomeNewProductRepo,
	NewHomeRecommendProductRepo,
	NewHomeRecommendSubjectRepo,
)

// Data .
type Data struct {
	db  *gorm.DB
	log *log.Helper
}

func NewDB(conf *conf.Data, logger1 log.Logger) *gorm.DB {
	log := log.NewHelper(log.With(logger1, "module", "ums-service/data/gorm"))

	db, err := gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}

	return db
}

// NewData .
func NewData(db *gorm.DB, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "order-service/data"))

	d := &Data{
		db:  db,
		log: log,
	}
	return d, func() {

	}, nil
}
