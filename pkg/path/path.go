package path

import (
	"os"
	"path/filepath"
)

// clickhouse
var (
	ClickhouseHome = os.Getenv("CLICKHOUSE_HOME")
)

// mate
var (
	ClickhouseMatePath    = filepath.FromSlash(ClickhouseHome + "/mate")
	ClickhouseScripts     = filepath.FromSlash(ClickhouseMatePath + "/scripts")
	ClickhouseStartScript = filepath.FromSlash(ClickhouseScripts + "/start-clickhouse.sh")
)
