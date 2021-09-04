package clickhouse

import (
	"clickhouse_mate_go/pkg/path"
	"clickhouse_mate_go/pkg/util"
	"github.com/paashzj/gutil"
	"go.uber.org/zap"
)

func startClickhousePlatform() error {
	stdout, stderr, err := gutil.CallScript(path.ClickhouseStartScript)
	util.Logger().Info("shell result ", zap.String("stdout", stdout), zap.String("stderr", stderr))
	return err
}