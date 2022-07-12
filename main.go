package main

import (
	"github.com/cloudwego/kitex/server"
	"github.com/mc_nft/middleware"
	"log"

	mc_nft "github.com/mc_nft/kitex_gen/mc_nft/nftservice"
)

func main() {
	svr := mc_nft.NewServer(new(NftServiceImpl),
		server.WithMiddleware(middleware.RequestValidateMW))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
