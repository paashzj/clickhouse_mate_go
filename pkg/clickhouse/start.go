package clickhouse

import (
	"clickhouse_mate_go/pkg/util"
	"go.uber.org/zap"
)

func Start() {
	err := startClickhousePlatform()
	if err != nil {
		util.Logger().Error("run start clickhouse scripts failed ", zap.Error(err))
		return
	}
}
