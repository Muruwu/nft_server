// Code generated by Kitex v0.3.2. DO NOT EDIT.
package nftservice

import (
	"github.com/cloudwego/kitex/server"
	"github.com/mc_nft/kitex_gen/mc_nft"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler mc_nft.NftService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
