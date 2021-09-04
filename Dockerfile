FROM ttbb/base:go AS build
COPY . /opt/sh/compile
WORKDIR /opt/sh/compile/pkg
RUN go build -o clickhouse_mate .


FROM ttbb/clickhouse:nake

LABEL maintainer="shoothzj@gmail.com"

COPY docker-build /opt/sh/clickhouse/mate

COPY --from=build /opt/sh/compile/pkg/clickhouse_mate /opt/sh/clickhouse/mate/clickhouse_mate

CMD ["/usr/local/bin/dumb-init", "bash", "-vx", "/opt/sh/clickhouse/mate/scripts/start.sh"]