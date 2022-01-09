#!/bin/bash

export REMOTE_MODE=false
mkdir $CLICKHOUSE_HOME/logs
nohup $CLICKHOUSE_HOME/mate/clickhouse_mate >>$CLICKHOUSE_HOME/logs/clickhouse_mate.stdout.log 2>>$CLICKHOUSE_HOME/logs/clickhouse_mate.stderr.log

