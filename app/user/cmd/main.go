package main

import (
	"fmt"
	"github.com/go-micro/plugins/v4/registry/etcd"
	"github.com/go-micro/plugins/v4/wrapper/trace/opentracing"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"micro-todoList/app/gateway/wrappers"
	"micro-todoList/app/user/repository/db/dao"
	"micro-todoList/app/user/service"
	"micro-todoList/config"
	"micro-todoList/idl/pb"
)

func main() {
	config.Init()
	dao.InitDB()
	// etcd注册件
	etcdReg := etcd.NewRegistry(
		registry.Addrs(fmt.Sprintf("%s:%s", config.EtcdHost, config.EtcdPort)),
	)

	// 初始化 Tracer
	tracer := wrappers.GetTracer(config.UserServiceName, config.UserServiceAddress)
	tracerHandler := opentracing.NewHandlerWrapper(tracer)

	// 得到一个微服务实例
	microService := micro.NewService(
		micro.Name(config.UserServiceName), // 微服务名字
		micro.Address(config.UserServiceAddress),
		micro.Registry(etcdReg), // etcd注册件
		micro.WrapHandler(tracerHandler),
	)
	// 结构命令行参数，初始化
	microService.Init()
	// 服务注册
	_ = pb.RegisterUserServiceHandler(microService.Server(), service.GetUserSrv())
	// 启动微服务
	_ = microService.Run()
}
