package config

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
	Charset    string

	RabbitMQ         string
	RabbitMQUser     string
	RabbitMQPassWord string
	RabbitMQHost     string
	RabbitMQPort     string

	EtcdHost string
	EtcdPort string

	GateWayServiceName    string
	GateWayServiceAddress string
	UserServiceName       string
	UserClientName        string
	UserServiceAddress    string
	TaskServiceName       string
	TaskClientName        string
	TaskServiceAddress    string

	ZipkinHost string
	ZipkinPort string
	ZipkinUrl  string

	RedisHost     string
	RedisPort     string
	RedisPassword string
	RedisDbName   int
)

func Init() {
	file, err := ini.Load("./config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
	}
	LoadMysqlData(file)
	LoadEtcd(file)
	LoadRabbitMQ(file)
	LoadServer(file)
	LoadZipkin(file)
	LoadRedisData(file)
}

func LoadMysqlData(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
	DbName = file.Section("mysql").Key("DbName").String()
	Charset = file.Section("mysql").Key("Charset").String()
}

func LoadRabbitMQ(file *ini.File) {
	RabbitMQ = file.Section("rabbitmq").Key("RabbitMQ").String()
	RabbitMQUser = file.Section("rabbitmq").Key("RabbitMQUser").String()
	RabbitMQPassWord = file.Section("rabbitmq").Key("RabbitMQPassWord").String()
	RabbitMQHost = file.Section("rabbitmq").Key("RabbitMQHost").String()
	RabbitMQPort = file.Section("rabbitmq").Key("RabbitMQPort").String()
}

func LoadEtcd(file *ini.File) {
	EtcdHost = file.Section("etcd").Key("EtcdHost").String()
	EtcdPort = file.Section("etcd").Key("EtcdPort").String()
}

func LoadServer(file *ini.File) {
	GateWayServiceName = file.Section("server").Key("GateWayServiceName").String()
	GateWayServiceAddress = file.Section("server").Key("GateWayServiceAddress").String()
	UserServiceName = file.Section("server").Key("UserServiceName").String()
	UserClientName = file.Section("server").Key("UserClientName").String()
	UserServiceAddress = file.Section("server").Key("UserServiceAddress").String()
	TaskServiceName = file.Section("server").Key("TaskServiceName").String()
	TaskClientName = file.Section("server").Key("TaskClientName").String()
	TaskServiceAddress = file.Section("server").Key("TaskServiceAddress").String()
}

func LoadZipkin(file *ini.File) {
	ZipkinHost = file.Section("zipkin").Key("ZipkinHost").String()
	ZipkinPort = file.Section("zipkin").Key("ZipkinPort").String()
	ZipkinUrl = file.Section("zipkin").Key("ZipkinUrl").String()
}

func LoadRedisData(file *ini.File) {
	RedisHost = file.Section("redis").Key("RedisHost").String()
	RedisPort = file.Section("redis").Key("RedisPort").String()
	RedisPassword = file.Section("redis").Key("RedisPassword").String()
}
