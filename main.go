package main

import (
	"log"

	"github.com/cloudwego/kitex/server"
	"github.com/mc_nft/dal"
	"github.com/mc_nft/middleware"

	mc_nft "github.com/mc_nft/kitex_gen/mc_nft/nftservice"
)

func main() {
	svr := mc_nft.NewServer(new(NftServiceImpl),
		server.WithMiddleware(middleware.RequestValidateMW))
	dal.InitDB()
	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
