package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-mall/app/ums/internal/biz"

	pb "kratos-mall/api/ums/v1"
)

type UmsService struct {
	pb.UnimplementedUmsServer
	uc  *biz.GreeterUsecase
	log *log.Helper
}

func NewUmsService(uc *biz.GreeterUsecase, logger log.Logger) *UmsService {
	return &UmsService{uc: uc, log: log.NewHelper(logger)}
}

func (s *UmsService) MemberAdd(ctx context.Context, req *pb.MemberAddReq) (*pb.MemberAddResp, error) {

	s.log.WithContext(ctx).Infof("SayHello Received: %v", req.GetNickname())

	s.uc.Create(ctx, &biz.Greeter{
		Hello: req.GetNickname(),
	})

	return &pb.MemberAddResp{}, nil
}
func (s *UmsService) MemberList(ctx context.Context, req *pb.MemberListReq) (*pb.MemberListResp, error) {
	return &pb.MemberListResp{}, nil
}
func (s *UmsService) MemberUpdate(ctx context.Context, req *pb.MemberUpdateReq) (*pb.MemberUpdateResp, error) {
	return &pb.MemberUpdateResp{}, nil
}
func (s *UmsService) MemberDelete(ctx context.Context, req *pb.MemberDeleteReq) (*pb.MemberDeleteResp, error) {
	return &pb.MemberDeleteResp{}, nil
}
func (s *UmsService) GrowthChangeHistoryAdd(ctx context.Context, req *pb.GrowthChangeHistoryAddReq) (*pb.GrowthChangeHistoryAddResp, error) {
	return &pb.GrowthChangeHistoryAddResp{}, nil
}
func (s *UmsService) GrowthChangeHistoryList(ctx context.Context, req *pb.GrowthChangeHistoryListReq) (*pb.GrowthChangeHistoryListResp, error) {
	return &pb.GrowthChangeHistoryListResp{}, nil
}
func (s *UmsService) GrowthChangeHistoryUpdate(ctx context.Context, req *pb.GrowthChangeHistoryUpdateReq) (*pb.GrowthChangeHistoryUpdateResp, error) {
	return &pb.GrowthChangeHistoryUpdateResp{}, nil
}
func (s *UmsService) GrowthChangeHistoryDelete(ctx context.Context, req *pb.GrowthChangeHistoryDeleteReq) (*pb.GrowthChangeHistoryDeleteResp, error) {
	return &pb.GrowthChangeHistoryDeleteResp{}, nil
}
func (s *UmsService) IntegrationChangeHistoryAdd(ctx context.Context, req *pb.IntegrationChangeHistoryAddReq) (*pb.IntegrationChangeHistoryAddResp, error) {
	return &pb.IntegrationChangeHistoryAddResp{}, nil
}
func (s *UmsService) IntegrationChangeHistoryList(ctx context.Context, req *pb.IntegrationChangeHistoryListReq) (*pb.IntegrationChangeHistoryListResp, error) {
	return &pb.IntegrationChangeHistoryListResp{}, nil
}
func (s *UmsService) IntegrationChangeHistoryUpdate(ctx context.Context, req *pb.IntegrationChangeHistoryUpdateReq) (*pb.IntegrationChangeHistoryUpdateResp, error) {
	return &pb.IntegrationChangeHistoryUpdateResp{}, nil
}
func (s *UmsService) IntegrationChangeHistoryDelete(ctx context.Context, req *pb.IntegrationChangeHistoryDeleteReq) (*pb.IntegrationChangeHistoryDeleteResp, error) {
	return &pb.IntegrationChangeHistoryDeleteResp{}, nil
}
func (s *UmsService) IntegrationConsumeSettingAdd(ctx context.Context, req *pb.IntegrationConsumeSettingAddReq) (*pb.IntegrationConsumeSettingAddResp, error) {
	return &pb.IntegrationConsumeSettingAddResp{}, nil
}
func (s *UmsService) IntegrationConsumeSettingList(ctx context.Context, req *pb.IntegrationConsumeSettingListReq) (*pb.IntegrationConsumeSettingListResp, error) {
	return &pb.IntegrationConsumeSettingListResp{}, nil
}
func (s *UmsService) IntegrationConsumeSettingUpdate(ctx context.Context, req *pb.IntegrationConsumeSettingUpdateReq) (*pb.IntegrationConsumeSettingUpdateResp, error) {
	return &pb.IntegrationConsumeSettingUpdateResp{}, nil
}
func (s *UmsService) IntegrationConsumeSettingDelete(ctx context.Context, req *pb.IntegrationConsumeSettingDeleteReq) (*pb.IntegrationConsumeSettingDeleteResp, error) {
	return &pb.IntegrationConsumeSettingDeleteResp{}, nil
}
func (s *UmsService) MemberLevelAdd(ctx context.Context, req *pb.MemberLevelAddReq) (*pb.MemberLevelAddResp, error) {
	return &pb.MemberLevelAddResp{}, nil
}
func (s *UmsService) MemberLevelList(ctx context.Context, req *pb.MemberLevelListReq) (*pb.MemberLevelListResp, error) {
	return &pb.MemberLevelListResp{}, nil
}
func (s *UmsService) MemberLevelUpdate(ctx context.Context, req *pb.MemberLevelUpdateReq) (*pb.MemberLevelUpdateResp, error) {
	return &pb.MemberLevelUpdateResp{}, nil
}
func (s *UmsService) MemberLevelDelete(ctx context.Context, req *pb.MemberLevelDeleteReq) (*pb.MemberLevelDeleteResp, error) {
	return &pb.MemberLevelDeleteResp{}, nil
}
func (s *UmsService) MemberLoginLogAdd(ctx context.Context, req *pb.MemberLoginLogAddReq) (*pb.MemberLoginLogAddResp, error) {
	return &pb.MemberLoginLogAddResp{}, nil
}
func (s *UmsService) MemberLoginLogList(ctx context.Context, req *pb.MemberLoginLogListReq) (*pb.MemberLoginLogListResp, error) {
	return &pb.MemberLoginLogListResp{}, nil
}
func (s *UmsService) MemberLoginLogUpdate(ctx context.Context, req *pb.MemberLoginLogUpdateReq) (*pb.MemberLoginLogUpdateResp, error) {
	return &pb.MemberLoginLogUpdateResp{}, nil
}
func (s *UmsService) MemberLoginLogDelete(ctx context.Context, req *pb.MemberLoginLogDeleteReq) (*pb.MemberLoginLogDeleteResp, error) {
	return &pb.MemberLoginLogDeleteResp{}, nil
}
func (s *UmsService) MemberMemberTagRelationAdd(ctx context.Context, req *pb.MemberMemberTagRelationAddReq) (*pb.MemberMemberTagRelationAddResp, error) {
	return &pb.MemberMemberTagRelationAddResp{}, nil
}
func (s *UmsService) MemberMemberTagRelationList(ctx context.Context, req *pb.MemberMemberTagRelationListReq) (*pb.MemberMemberTagRelationListResp, error) {
	return &pb.MemberMemberTagRelationListResp{}, nil
}
func (s *UmsService) MemberMemberTagRelationUpdate(ctx context.Context, req *pb.MemberMemberTagRelationUpdateReq) (*pb.MemberMemberTagRelationUpdateResp, error) {
	return &pb.MemberMemberTagRelationUpdateResp{}, nil
}
func (s *UmsService) MemberMemberTagRelationDelete(ctx context.Context, req *pb.MemberMemberTagRelationDeleteReq) (*pb.MemberMemberTagRelationDeleteResp, error) {
	return &pb.MemberMemberTagRelationDeleteResp{}, nil
}
func (s *UmsService) MemberProductCategoryRelationAdd(ctx context.Context, req *pb.MemberProductCategoryRelationAddReq) (*pb.MemberProductCategoryRelationAddResp, error) {
	return &pb.MemberProductCategoryRelationAddResp{}, nil
}
func (s *UmsService) MemberProductCategoryRelationList(ctx context.Context, req *pb.MemberProductCategoryRelationListReq) (*pb.MemberProductCategoryRelationListResp, error) {
	return &pb.MemberProductCategoryRelationListResp{}, nil
}
func (s *UmsService) MemberProductCategoryRelationUpdate(ctx context.Context, req *pb.MemberProductCategoryRelationUpdateReq) (*pb.MemberProductCategoryRelationUpdateResp, error) {
	return &pb.MemberProductCategoryRelationUpdateResp{}, nil
}
func (s *UmsService) MemberProductCategoryRelationDelete(ctx context.Context, req *pb.MemberProductCategoryRelationDeleteReq) (*pb.MemberProductCategoryRelationDeleteResp, error) {
	return &pb.MemberProductCategoryRelationDeleteResp{}, nil
}
func (s *UmsService) MemberReceiveAddressAdd(ctx context.Context, req *pb.MemberReceiveAddressAddReq) (*pb.MemberReceiveAddressAddResp, error) {
	return &pb.MemberReceiveAddressAddResp{}, nil
}
func (s *UmsService) MemberReceiveAddressList(ctx context.Context, req *pb.MemberReceiveAddressListReq) (*pb.MemberReceiveAddressListResp, error) {
	return &pb.MemberReceiveAddressListResp{}, nil
}
func (s *UmsService) MemberReceiveAddressUpdate(ctx context.Context, req *pb.MemberReceiveAddressUpdateReq) (*pb.MemberReceiveAddressUpdateResp, error) {
	return &pb.MemberReceiveAddressUpdateResp{}, nil
}
func (s *UmsService) MemberReceiveAddressDelete(ctx context.Context, req *pb.MemberReceiveAddressDeleteReq) (*pb.MemberReceiveAddressDeleteResp, error) {
	return &pb.MemberReceiveAddressDeleteResp{}, nil
}
func (s *UmsService) MemberRuleSettingAdd(ctx context.Context, req *pb.MemberRuleSettingAddReq) (*pb.MemberRuleSettingAddResp, error) {
	return &pb.MemberRuleSettingAddResp{}, nil
}
func (s *UmsService) MemberRuleSettingList(ctx context.Context, req *pb.MemberRuleSettingListReq) (*pb.MemberRuleSettingListResp, error) {
	return &pb.MemberRuleSettingListResp{}, nil
}
func (s *UmsService) MemberRuleSettingUpdate(ctx context.Context, req *pb.MemberRuleSettingUpdateReq) (*pb.MemberRuleSettingUpdateResp, error) {
	return &pb.MemberRuleSettingUpdateResp{}, nil
}
func (s *UmsService) MemberRuleSettingDelete(ctx context.Context, req *pb.MemberRuleSettingDeleteReq) (*pb.MemberRuleSettingDeleteResp, error) {
	return &pb.MemberRuleSettingDeleteResp{}, nil
}
func (s *UmsService) MemberStatisticsInfoAdd(ctx context.Context, req *pb.MemberStatisticsInfoAddReq) (*pb.MemberStatisticsInfoAddResp, error) {
	return &pb.MemberStatisticsInfoAddResp{}, nil
}
func (s *UmsService) MemberStatisticsInfoList(ctx context.Context, req *pb.MemberStatisticsInfoListReq) (*pb.MemberStatisticsInfoListResp, error) {
	return &pb.MemberStatisticsInfoListResp{}, nil
}
func (s *UmsService) MemberStatisticsInfoUpdate(ctx context.Context, req *pb.MemberStatisticsInfoUpdateReq) (*pb.MemberStatisticsInfoUpdateResp, error) {
	return &pb.MemberStatisticsInfoUpdateResp{}, nil
}
func (s *UmsService) MemberStatisticsInfoDelete(ctx context.Context, req *pb.MemberStatisticsInfoDeleteReq) (*pb.MemberStatisticsInfoDeleteResp, error) {
	return &pb.MemberStatisticsInfoDeleteResp{}, nil
}
func (s *UmsService) MemberTagAdd(ctx context.Context, req *pb.MemberTagAddReq) (*pb.MemberTagAddResp, error) {
	return &pb.MemberTagAddResp{}, nil
}
func (s *UmsService) MemberTagList(ctx context.Context, req *pb.MemberTagListReq) (*pb.MemberTagListResp, error) {
	return &pb.MemberTagListResp{}, nil
}
func (s *UmsService) MemberTagUpdate(ctx context.Context, req *pb.MemberTagUpdateReq) (*pb.MemberTagUpdateResp, error) {
	return &pb.MemberTagUpdateResp{}, nil
}
func (s *UmsService) MemberTagDelete(ctx context.Context, req *pb.MemberTagDeleteReq) (*pb.MemberTagDeleteResp, error) {
	return &pb.MemberTagDeleteResp{}, nil
}
func (s *UmsService) MemberTaskAdd(ctx context.Context, req *pb.MemberTaskAddReq) (*pb.MemberTaskAddResp, error) {
	return &pb.MemberTaskAddResp{}, nil
}
func (s *UmsService) MemberTaskList(ctx context.Context, req *pb.MemberTaskListReq) (*pb.MemberTaskListResp, error) {
	return &pb.MemberTaskListResp{}, nil
}
func (s *UmsService) MemberTaskUpdate(ctx context.Context, req *pb.MemberTaskUpdateReq) (*pb.MemberTaskUpdateResp, error) {
	return &pb.MemberTaskUpdateResp{}, nil
}
func (s *UmsService) MemberTaskDelete(ctx context.Context, req *pb.MemberTaskDeleteReq) (*pb.MemberTaskDeleteResp, error) {
	return &pb.MemberTaskDeleteResp{}, nil
}
