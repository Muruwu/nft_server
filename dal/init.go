package dal

import (
	"github.com/mc_nft/dal/db"
	"github.com/mc_nft/dal/kv"
)

func InitDB() {
	db.InitConf()
	kv.InitRedis()
}
