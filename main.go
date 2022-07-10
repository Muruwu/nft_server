package main

import (
	"log"

	mc_nft "github.com/mc_nft/kitex_gen/mc_nft/nftservice"
)

func main() {
	svr := mc_nft.NewServer(new(NftServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
