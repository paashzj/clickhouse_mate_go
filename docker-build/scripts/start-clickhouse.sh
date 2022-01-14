#!/bin/bash

mkdir -p $CLICKHOUSE_HOME/logs
nohup clickhouse server >>$CLICKHOUSE_HOME/logs/clickhouse.stdout.log 2>>$CLICKHOUSE_HOME/logs/clickhouse.stderr.log &

