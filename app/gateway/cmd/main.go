package main

import (
	"fmt"
	"github.com/go-micro/plugins/v4/registry/etcd"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/web"
	"micro-todoList/app/common"
	"micro-todoList/app/gateway/router"
	"micro-todoList/app/gateway/rpc"
	"micro-todoList/app/user/repository/cache"
	"micro-todoList/config"
	log "micro-todoList/pkg/logger"
	"time"
)

func main() {
	config.Init()
	rpc.InitRPC()
	cache.InitCache()
	log.InitLog()
	etcdReg := etcd.NewRegistry(
		registry.Addrs(fmt.Sprintf("%s:%s", config.EtcdHost, config.EtcdPort)),
	)
	// 初始化 Tracer
	tracer := common.GetTracer(config.GateWayServiceName, config.GateWayServiceAddress)
	// 初始化 Prometheus
	common.PrometheusBoot(config.PrometheusGateWayPath, config.PrometheusGateWayAddress)

	// 创建微服务实例，使用gin暴露http接口并注册到etcd
	server := web.NewService(
		web.Name(config.GateWayServiceName),
		web.Address(config.GateWayServiceAddress),
		// 将服务调用实例使用gin处理
		web.Handler(router.NewRouter(tracer)),
		web.Registry(etcdReg),
		web.RegisterTTL(time.Second*30),
		web.RegisterInterval(time.Second*15),
		web.Metadata(map[string]string{"protocol": "http"}),
	)
	// 接收命令行参数
	_ = server.Init()
	_ = server.Run()
}
