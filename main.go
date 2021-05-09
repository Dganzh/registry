package main

import (
	log "github.com/Dganzh/zlog"
	"github.com/Dganzh/zrpc"
	"github.com/gin-gonic/gin"
)

var cfg *Config

// 提供http接口便于管控服务
func startHttpServer() {
	r := gin.Default()
	r.GET("/ping", Ping)
	r.GET("/stop_all", StopKVServer)
	if err := r.Run(cfg.HttpAddr); err != nil {
		log.Fatalw("http server run failed============>", "err", err)
	}
}


func main() {
	log.Info("start registry")
	cfg = defaultCfg // use default config
	go startHttpServer()
	server := zrpc.NewServer(cfg.RpcAddr)
	r := NewRegistry()
	SetGlobalRegistry(r)
	server.Register(r)
	server.Start()
}
