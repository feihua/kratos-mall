package data

import (
	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"kratos-mall/app/pms/internal/conf"
)

var ProviderSet = wire.NewSet(
	NewData,
	NewDB,
	NewGreeterRepo,
	NewBrandRepo,
	NewCategoryRepo,
	NewCommentRepo,
	NewCommentReplayRepo,
	NewFeightTemplateRepo,
	NewMemberPriceRepo,
	NewOperateHistoryRepo,
	NewProductRepo,
	NewSkuStockRepo,
	NewVertifyRecordRepo,
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

	//if err := db.AutoMigrate(&Users{}); err != nil {
	//	log.Fatal(err)
	//}
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
