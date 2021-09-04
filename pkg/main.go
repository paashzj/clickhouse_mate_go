package main

import (
	"clickhouse_mate_go/pkg/clickhouse"
	"clickhouse_mate_go/pkg/util"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
)

func main() {
	util.Logger().Debug("this is a debug msg")
	util.Logger().Info("this is a info msg")
	util.Logger().Error("this is a error msg")
	clickhouse.Start()
	router := gin.Default()
	util.Logger().Info("clickhouse mate started")
	err := router.Run(":31018")
	if err != nil {
		util.Logger().Error("open http server failed")
		return
	}
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	for {
		select {
		case <-interrupt:
			return
		}
	}
}
