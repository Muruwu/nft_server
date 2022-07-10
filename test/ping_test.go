package test

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/mc_nft/kitex_gen/mc_nft/nftservice"
	"github.com/mc_nft/kitex_gen/ping"
)

func TestPingClient(t *testing.T) {
	c, err := nftservice.NewClient("nft_service", client.WithHostPorts("127.0.0.1:8888"))

	if err != nil {
		log.Fatal(err)
	}
	for {
		req := &ping.PingRequest{}
		resp, err := c.Ping(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(resp)
		time.Sleep(time.Second)
	}

}
