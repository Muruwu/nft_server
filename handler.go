package main

import (
	"context"

	credits "github.com/mc_nft/kitex_gen/credits"
	"github.com/mc_nft/kitex_gen/ping"
	props "github.com/mc_nft/kitex_gen/props"
	user "github.com/mc_nft/kitex_gen/user"
	virtualCurrency "github.com/mc_nft/kitex_gen/virtual_currency"
)

// NftServiceImpl implements the last service interface defined in the IDL.
type NftServiceImpl struct{}

type BaseResponse interface {
}

// Ping implements the NftServiceImpl interface.
func (s *NftServiceImpl) Ping(ctx context.Context, req *ping.PingRequest) (resp *ping.PingResponse, err error) {
	resp = &ping.PingResponse{Msg: "pong"}
	return
}

// UserLogin implements the NftServiceImpl interface.
func (s *NftServiceImpl) UserLogin(ctx context.Context, req *user.UserLoginRequest) (resp *user.UserLoginResponse, err error) {
	// TODO: Your code here...
	return
}

// SmsSend implements the NftServiceImpl interface.
func (s *NftServiceImpl) SmsSend(ctx context.Context, req *user.SmsSendRequest) (resp *user.SmsSendResponse, err error) {
	// TODO: Your code here...
	return
}

// Authenticate implements the NftServiceImpl interface.
func (s *NftServiceImpl) Authenticate(ctx context.Context, req *user.AuthenticateRequest) (resp *user.AuthenticateResponse, err error) {
	// TODO: Your code here...
	return
}

// UserDetail implements the NftServiceImpl interface.
func (s *NftServiceImpl) UserDetail(ctx context.Context, req *user.UserDetailRequest) (resp *user.UserDetailResponse, err error) {
	// TODO: Your code here...
	return
}

// AddProps implements the NftServiceImpl interface.
func (s *NftServiceImpl) AddProps(ctx context.Context, req *props.AddPropsRequest) (resp *props.AddPropsResponse, err error) {
	// TODO: Your code here...
	return
}

// UpdateProps implements the NftServiceImpl interface.
func (s *NftServiceImpl) UpdateProps(ctx context.Context, req *props.UpdatePropsRequest) (resp *props.UpdatePropsResponse, err error) {
	// TODO: Your code here...
	return
}

// AddDiamond implements the NftServiceImpl interface.
func (s *NftServiceImpl) AddDiamond(ctx context.Context, req *virtualCurrency.AddDiamondRequest) (resp *virtualCurrency.AddDiamondResponse, err error) {
	// TODO: Your code here...
	return
}

// ConsumeDiamond implements the NftServiceImpl interface.
func (s *NftServiceImpl) ConsumeDiamond(ctx context.Context, req *virtualCurrency.ConsumeDiamondRequest) (resp *virtualCurrency.ConsumeDiamondResponse, err error) {
	// TODO: Your code here...
	return
}

// QueryUserDiamond implements the NftServiceImpl interface.
func (s *NftServiceImpl) QueryUserDiamond(ctx context.Context, req *virtualCurrency.QueryUserDiamondRequest) (resp *virtualCurrency.QueryUserDiamondResponse, err error) {
	// TODO: Your code here...
	return
}

// AddOrUpdateCredits implements the NftServiceImpl interface.
func (s *NftServiceImpl) AddOrUpdateCredits(ctx context.Context, req *credits.AddOrUpdateCreditsRequest) (resp *credits.AddOrUpdateCreditsResponse, err error) {
	// TODO: Your code here...
	return
}

// GetUserCredits implements the NftServiceImpl interface.
func (s *NftServiceImpl) GetUserCredits(ctx context.Context, req *credits.GetUserCreditsRequest) (resp *credits.GetUserCreditsResponse, err error) {
	// TODO: Your code here...
	return
}

// GetCreditsRank implements the NftServiceImpl interface.
func (s *NftServiceImpl) GetCreditsRank(ctx context.Context, req *credits.GetCreditsRankRequest) (resp *credits.GetCreditsRankResponse, err error) {
	// TODO: Your code here...
	return
}

// DistributeProps implements the NftServiceImpl interface.
func (s *NftServiceImpl) DistributeProps(ctx context.Context, req *props.DistributePropsRequest) (resp *props.DistributePropsResponse, err error) {
	// TODO: Your code here...
	return
}

// QueryProps implements the NftServiceImpl interface.
func (s *NftServiceImpl) QueryProps(ctx context.Context, req *props.QueryPropsRequest) (resp *props.QueryPropsResponse, err error) {
	// TODO: Your code here...
	return
}

// ConsumeProps implements the NftServiceImpl interface.
func (s *NftServiceImpl) ConsumeProps(ctx context.Context, req *props.ConsumePropsRequest) (resp *props.ConsumePropsResponse, err error) {
	// TODO: Your code here...
	return
}
