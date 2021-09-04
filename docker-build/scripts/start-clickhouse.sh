#!/bin/bash

nohup clickhouse server >>$CLICKHOUSE_HOME/clickhouse.stdout.log 2>>$CLICKHOUSE_HOME/clickhouse.stderr.log &