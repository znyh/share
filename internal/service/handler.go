package service

import (
	"context"

	pb "github.com/znyh/proto/share"
)

// 共享数据
func (s *Service) UpdateShare(ctx context.Context, req *pb.UpdateKeyValuesReq) (resp *pb.CommonResp, err error) {
	return
}

// 获取共享数据
func (s *Service) LoadShare(ctx context.Context, req *pb.LoadValuesReq) (resp *pb.LoadValuesResp, err error) {
	return
}

func (s *Service) UpdateGameArena(ctx context.Context, req *pb.UpdateGameArenaReq) (resp *pb.UpdateGameArenaResp, err error) {
	return
}

func (s *Service) DeleteGameArena(ctx context.Context, req *pb.DeleteGameArenaReq) (resp *pb.DeleteGameArenaResp, err error) {
	return
}

func (s *Service) LoadGameAvoidInjury(ctx context.Context, req *pb.LoadGameAvoidInjuryReq) (resp *pb.LoadGameAvoidInjuryResp, err error) {
	return
}

func (s *Service) UpdateKeyValues(ctx context.Context, req *pb.UpdateKeyValuesReq) (resp *pb.CommonResp, err error) {
	return
}

func (s *Service) LoadValues(ctx context.Context, req *pb.LoadValuesReq) (resp *pb.LoadValuesResp, err error) {
	return
}

func (s *Service) SetExpire(ctx context.Context, req *pb.SetExpireReq) (resp *pb.CommonResp, err error) {
	return
}

func (s *Service) LoadExpire(ctx context.Context, req *pb.LoadExpireReq) (resp *pb.LoadExpireResp, err error) {
	return
}

func (s *Service) IncrGameData(ctx context.Context, req *pb.IncrGameDataReq) (resp *pb.CommonResp, err error) {
	return
}

func (s *Service) LoadGameData(ctx context.Context, req *pb.LoadGameDataReq) (resp *pb.LoadGameDataResp, err error) {
	return
}

func (s *Service) SetAdd(ctx context.Context, req *pb.SetReq) (resp *pb.CommonResp, err error) {
	return
}

func (s *Service) SetDel(ctx context.Context, req *pb.SetReq) (resp *pb.CommonResp, err error) {
	return
}

func (s *Service) SetMembers(ctx context.Context, req *pb.SetReq) (resp *pb.SetMembersResp, err error) {
	return
}

func (s *Service) SetRandom(ctx context.Context, req *pb.SetRandomReq) (resp *pb.SetRandomResp, err error) {
	return
}

func (s *Service) SetIsMember(ctx context.Context, req *pb.SetReq) (resp *pb.SetIsMemberResp, err error) {
	return
}

func (s *Service) SetDelAll(ctx context.Context, req *pb.SetReq) (resp *pb.CommonResp, err error) {
	return
}

func (s *Service) HSetUpdate(ctx context.Context, req *pb.HSetUpdateReq) (resp *pb.CommonResp, err error) {
	return
}

func (s *Service) HSetLoad(ctx context.Context, req *pb.HSetLoadReq) (resp *pb.HSetLoadResp, err error) {
	return
}

func (s *Service) HSetDel(ctx context.Context, req *pb.HSetDelReq) (resp *pb.CommonResp, err error) {
	return
}

func (s *Service) HSetUpdateValues(ctx context.Context, req *pb.HSetUpdateValuesReq) (resp *pb.CommonResp, err error) {
	return
}

func (s *Service) HSetLoadValues(ctx context.Context, req *pb.HSetLoadValuesReq) (resp *pb.HSetLoadValuesResp, err error) {
	return
}

func (s *Service) HSetDelValues(ctx context.Context, req *pb.HSetDelValuesReq) (resp *pb.CommonResp, err error) {
	return
}

func (s *Service) HSetDelHash(ctx context.Context, req *pb.HSetDelHashReq) (resp *pb.CommonResp, err error) {
	return
}

func (s *Service) HSetAdd(ctx context.Context, req *pb.HSetAddReq) (resp *pb.CommonResp, err error) {
	return
}

func (s *Service) HSetIncr(ctx context.Context, req *pb.HSetIncrReq) (resp *pb.CommonResp, err error) {
	return
}

func (s *Service) HListAdd(ctx context.Context, req *pb.HListReq) (resp *pb.CommonResp, err error) {
	return
}

func (s *Service) HListAddLast(ctx context.Context, req *pb.HListReq) (resp *pb.CommonResp, err error) {
	return
}

func (s *Service) HListDel(ctx context.Context, req *pb.HListReq) (resp *pb.CommonResp, err error) {
	return
}

func (s *Service) HListDelAll(ctx context.Context, req *pb.HListReq) (resp *pb.CommonResp, err error) {
	return
}

func (s *Service) HListLoadAll(ctx context.Context, req *pb.HListReq) (resp *pb.HListLoadResp, err error) {
	return
}

func (s *Service) HListLoadRange(ctx context.Context, req *pb.HListLoadRangeReq) (resp *pb.HListLoadResp, err error) {
	return
}

func (s *Service) HListPopFirst(ctx context.Context, req *pb.HListReq) (resp *pb.HListPopResp, err error) {
	return
}

func (s *Service) HListPopLast(ctx context.Context, req *pb.HListReq) (resp *pb.HListPopResp, err error) {
	return
}

func (s *Service) ZSetLoadRange(ctx context.Context, req *pb.ZSetLoadRangeReq) (resp *pb.ZSetLoadRangeResp, err error) {
	return
}

func (s *Service) ZSetLoadRevRange(ctx context.Context, req *pb.ZSetLoadRangeReq) (resp *pb.ZSetLoadRangeResp, err error) {
	return
}

func (s *Service) ZSetGetScore(ctx context.Context, req *pb.ZSetReq) (resp *pb.ZSetGetScoreResp, err error) {
	return
}

func (s *Service) ZSetUpdateScore(ctx context.Context, req *pb.ZSetUpdateScoreReq) (resp *pb.CommonResp, err error) {
	return
}

func (s *Service) ZSetDel(ctx context.Context, req *pb.ZSetReq) (resp *pb.CommonResp, err error) {
	return
}

func (s *Service) ZSetDelAll(ctx context.Context, req *pb.ZSetReq) (resp *pb.CommonResp, err error) {
	return
}

func (s *Service) ZSetGetRank(ctx context.Context, req *pb.ZSetReq) (resp *pb.ZSetGetRankResp, err error) {
	return
}

func (s *Service) ZSetGetRevRank(ctx context.Context, req *pb.ZSetReq) (resp *pb.ZSetGetRankResp, err error) {
	return
}

func (s *Service) ZSetRefreshRank(ctx context.Context, req *pb.ZSetRefreshRankReq) (resp *pb.CommonResp, err error) {
	return
}
