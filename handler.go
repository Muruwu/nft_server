package main

import (
	"context"

	"github.com/mc_nft/kitex_gen/ping"
)

// NftServiceImpl implements the last service interface defined in the IDL.
type NftServiceImpl struct{}

// Ping implements the NftServiceImpl interface.
func (s *NftServiceImpl) Ping(ctx context.Context, req *ping.PingRequest) (resp *ping.PingResponse, err error) {
	resp = &ping.PingResponse{Msg: "pong"}
	return
}
