#!/bin/bash

export REMOTE_MODE=false

nohup $CLICKHOUSE_HOME/mate/clickhouse_mate >>$CLICKHOUSE_HOME/clickhouse_mate.stdout.log 2>>$CLICKHOUSE_HOME/clickhouse_mate.stderr.log