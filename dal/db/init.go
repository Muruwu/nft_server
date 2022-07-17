package db

import (
	"errors"
	"fmt"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/mc_nft/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DBWriter *gorm.DB
	DBReader *gorm.DB
)

func InitConf() {
	dbConf := conf.GetInstance().Mysql
	if err := checkDBConf(dbConf.Master); err != nil {
		panic(err)
	}
	if err := checkDBConf(dbConf.Slave); err != nil {
		panic(err)
	}

	var (
		err error
	)
	if DBWriter, err = gorm.Open(mysql.Open(fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConf.Master.User, dbConf.Master.Password, dbConf.Master.Host, dbConf.Master.Port, dbConf.Master.DbName),
	), &gorm.Config{}); err != nil {
		panic(err)
	}

	if DBReader, err = gorm.Open(mysql.Open(fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConf.Slave.User, dbConf.Slave.Password, dbConf.Slave.Host, dbConf.Slave.Port, dbConf.Slave.DbName),
	), &gorm.Config{}); err != nil {
		panic(err)
	}
	klog.Info("init gorm db success")
	// 设置连接池 wrap sql.DB
}

func checkDBConf(conf conf.MysqlConnConf) error {
	if conf.Host == "" ||
		conf.User == "" ||
		conf.Password == "" {
		return errors.New("invalid mysql conf")
	}
	return nil
}
