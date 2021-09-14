package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
	"kratos-mall/app/front/admin/internal/biz"
	"kratos-mall/app/front/admin/internal/biz/ums"

	pb "kratos-mall/api/front/admin/v1"
)

type UmsService struct {
	pb.UnimplementedUmsServer
	uc                              *biz.GreeterUsecase
	changeHistoryUseCase            *ums.ChangeHistoryUseCase
	integrationChangeHistoryUseCase *ums.IntegrationChangeHistoryUseCase
	consumeSettingUseCase           *ums.ConsumeSettingUseCase
	loginLogUseCase                 *ums.LoginLogUseCase
	memberUseCase                   *ums.MemberUseCase
	memberAddressUseCase            *ums.MemberAddressUseCase
	memberLevelUseCase              *ums.MemberLevelUseCase
	ruleSettingUseCase              *ums.RuleSettingUseCase
	tagUseCase                      *ums.TagUseCase
	taskUseCase                     *ums.TaskUseCase
	log                             *log.Helper
}

func NewUmsService(uc *biz.GreeterUsecase, logger log.Logger, changeHistoryUseCase *ums.ChangeHistoryUseCase,
	integrationChangeHistoryUseCase *ums.IntegrationChangeHistoryUseCase,
	consumeSettingUseCase *ums.ConsumeSettingUseCase,
	loginLogUseCase *ums.LoginLogUseCase,
	memberUseCase *ums.MemberUseCase,
	memberAddressUseCase *ums.MemberAddressUseCase,
	memberLevelUseCase *ums.MemberLevelUseCase,
	ruleSettingUseCase *ums.RuleSettingUseCase,
	tagUseCase *ums.TagUseCase,
	taskUseCase *ums.TaskUseCase) *UmsService {
	return &UmsService{
		uc:                              uc,
		log:                             log.NewHelper(logger),
		changeHistoryUseCase:            changeHistoryUseCase,
		consumeSettingUseCase:           consumeSettingUseCase,
		integrationChangeHistoryUseCase: integrationChangeHistoryUseCase,
		loginLogUseCase:                 loginLogUseCase,
		memberUseCase:                   memberUseCase,
		memberAddressUseCase:            memberAddressUseCase,
		memberLevelUseCase:              memberLevelUseCase,
		ruleSettingUseCase:              ruleSettingUseCase,
		tagUseCase:                      tagUseCase,
		taskUseCase:                     taskUseCase}
}

func (s *UmsService) MemberAdd(ctx context.Context, req *pb.MemberAddReq) (*pb.MemberAddResp, error) {
	return &pb.MemberAddResp{}, nil
}
func (s *UmsService) MemberList(ctx context.Context, req *pb.MemberListReq) (*pb.MemberListResp, error) {
	listResp, _ := s.memberUseCase.ListMember(ctx, &ums.MemberListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	list := make([]*pb.MemberListData, 0)

	copier.Copy(&list, &listResp.List)
	return &pb.MemberListResp{
		Code:     "000000",
		Message:  "查询订单信息成功",
		Current:  int32(req.Current),
		PageSize: int32(req.PageSize),
		Total:    int32(listResp.Total),
		Data:     list,
		Success:  true,
	}, nil
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
	listResp, _ := s.changeHistoryUseCase.ListChangeHistory(ctx, &ums.ChangeHistoryListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	list := make([]*pb.GrowthChangeHistoryListData, 0)

	copier.Copy(&list, &listResp.List)
	return &pb.GrowthChangeHistoryListResp{
		Code:     "000000",
		Message:  "查询订单信息成功",
		Current:  int32(req.Current),
		PageSize: int32(req.PageSize),
		Total:    int32(listResp.Total),
		Data:     list,
		Success:  true,
	}, nil
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
	listResp, _ := s.integrationChangeHistoryUseCase.ListIntegrationChangeHistory(ctx, &ums.IntegrationChangeHistoryListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	list := make([]*pb.IntegrationChangeHistoryListData, 0)

	copier.Copy(&list, &listResp.List)
	return &pb.IntegrationChangeHistoryListResp{
		Code:     "000000",
		Message:  "查询订单信息成功",
		Current:  int32(req.Current),
		PageSize: int32(req.PageSize),
		Total:    int32(listResp.Total),
		Data:     list,
		Success:  true,
	}, nil
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
	listResp, _ := s.consumeSettingUseCase.ListConsumeSetting(ctx, &ums.ConsumeSettingListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	list := make([]*pb.IntegrationConsumeSettingListData, 0)

	copier.Copy(&list, &listResp.List)
	return &pb.IntegrationConsumeSettingListResp{
		Code:     "000000",
		Message:  "查询订单信息成功",
		Current:  int32(req.Current),
		PageSize: int32(req.PageSize),
		Total:    int32(listResp.Total),
		Data:     list,
		Success:  true,
	}, nil
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
	listResp, _ := s.memberLevelUseCase.ListMemberLevl(ctx, &ums.MemberLevlListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	list := make([]*pb.MemberLevelListData, 0)

	copier.Copy(&list, &listResp.List)
	return &pb.MemberLevelListResp{
		Code:     "000000",
		Message:  "查询订单信息成功",
		Current:  int32(req.Current),
		PageSize: int32(req.PageSize),
		Total:    int32(listResp.Total),
		Data:     list,
		Success:  true,
	}, nil
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
	listResp, _ := s.loginLogUseCase.ListLoginLog(ctx, &ums.LoginLogListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	list := make([]*pb.MemberLoginLogListData, 0)

	copier.Copy(&list, &listResp.List)
	return &pb.MemberLoginLogListResp{
		Code:     "000000",
		Message:  "查询订单信息成功",
		Current:  int32(req.Current),
		PageSize: int32(req.PageSize),
		Total:    int32(listResp.Total),
		Data:     list,
		Success:  true,
	}, nil
}
func (s *UmsService) MemberLoginLogUpdate(ctx context.Context, req *pb.MemberLoginLogUpdateReq) (*pb.MemberLoginLogUpdateResp, error) {
	return &pb.MemberLoginLogUpdateResp{}, nil
}
func (s *UmsService) MemberLoginLogDelete(ctx context.Context, req *pb.MemberLoginLogDeleteReq) (*pb.MemberLoginLogDeleteResp, error) {
	return &pb.MemberLoginLogDeleteResp{}, nil
}
func (s *UmsService) MemberReceiveAddressAdd(ctx context.Context, req *pb.MemberReceiveAddressAddReq) (*pb.MemberReceiveAddressAddResp, error) {
	return &pb.MemberReceiveAddressAddResp{}, nil
}
func (s *UmsService) MemberReceiveAddressList(ctx context.Context, req *pb.MemberReceiveAddressListReq) (*pb.MemberReceiveAddressListResp, error) {
	listResp, _ := s.memberAddressUseCase.ListMemberAddress(ctx, &ums.MemberAddressListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	list := make([]*pb.MemberReceiveAddressListData, 0)

	copier.Copy(&list, &listResp.List)
	return &pb.MemberReceiveAddressListResp{
		Code:     "000000",
		Message:  "查询订单信息成功",
		Current:  int32(req.Current),
		PageSize: int32(req.PageSize),
		Total:    int32(listResp.Total),
		Data:     list,
		Success:  true,
	}, nil
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
	listResp, _ := s.ruleSettingUseCase.ListRuleSetting(ctx, &ums.RuleSettingListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	list := make([]*pb.MemberRuleSettingListData, 0)

	copier.Copy(&list, &listResp.List)
	return &pb.MemberRuleSettingListResp{
		Code:     "000000",
		Message:  "查询订单信息成功",
		Current:  int32(req.Current),
		PageSize: int32(req.PageSize),
		Total:    int32(listResp.Total),
		Data:     list,
		Success:  true,
	}, nil
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
	listResp, _ := s.tagUseCase.ListTag(ctx, &ums.TagListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	list := make([]*pb.MemberTagListData, 0)

	copier.Copy(&list, &listResp.List)
	return &pb.MemberTagListResp{
		Code:     "000000",
		Message:  "查询订单信息成功",
		Current:  int32(req.Current),
		PageSize: int32(req.PageSize),
		Total:    int32(listResp.Total),
		Data:     list,
		Success:  true,
	}, nil
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
	listResp, _ := s.taskUseCase.ListTask(ctx, &ums.TaskListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	list := make([]*pb.MemberTaskListData, 0)

	copier.Copy(&list, &listResp.List)
	return &pb.MemberTaskListResp{
		Code:     "000000",
		Message:  "查询订单信息成功",
		Current:  int32(req.Current),
		PageSize: int32(req.PageSize),
		Total:    int32(listResp.Total),
		Data:     list,
		Success:  true,
	}, nil
}
func (s *UmsService) MemberTaskUpdate(ctx context.Context, req *pb.MemberTaskUpdateReq) (*pb.MemberTaskUpdateResp, error) {
	return &pb.MemberTaskUpdateResp{}, nil
}
func (s *UmsService) MemberTaskDelete(ctx context.Context, req *pb.MemberTaskDeleteReq) (*pb.MemberTaskDeleteResp, error) {
	return &pb.MemberTaskDeleteResp{}, nil
}
