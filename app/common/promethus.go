package common

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"micro-todoList/config"
	"net/http"
)

func PrometheusBoot(port int) {
	http.Handle(fmt.Sprintf("/metrics"), promhttp.Handler())
	//启动web 服务
	go func() {
		address := fmt.Sprintf("%s:%d", config.PrometheusHost, port)
		err := http.ListenAndServe(address, nil)
		if err != nil {
			log.Fatal("启动失败")
		}
	}()

}
