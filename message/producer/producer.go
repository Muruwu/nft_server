package producer

import (
	"fmt"
	"os"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

var Producer rocketmq.Producer

func InitProducer() {
	// 服务接入地址  (注意：需要在接入地址前面加上  http:// 或 https:// 否则无法解析)
	var serverAddress = "https://rocketmq-xxx.rocketmq.ap-bj.public.tencenttdmq.com:9876"
	// 授权角色名
	var secretKey = "admin"
	// 授权角色密钥
	var accessKey = "eyJrZXlJZC...."
	// 命名空间全称
	var nameSpace = "rocketmq-xxx|namespace_go"
	// 生产者组名称
	var groupName = "group1"
	// 创建消息生产者
	Producer, _ = rocketmq.NewProducer(
		// 设置服务地址
		producer.WithNsResolver(primitive.NewPassthroughResolver([]string{serverAddress})),
		// 设置acl权限
		producer.WithCredentials(primitive.Credentials{
			SecretKey: secretKey,
			AccessKey: accessKey,
		}),
		// 设置生产组
		producer.WithGroupName(groupName),
		// 设置命名空间名称
		producer.WithNamespace(nameSpace),
		// 设置发送失败重试次数
		producer.WithRetry(2),
	)
	// 启动producer
	err := Producer.Start()
	if err != nil {
		fmt.Printf("start producer error: %s", err.Error())
		os.Exit(1)
	}
}

// 异步发送消息

// 同步发送消息
