FROM postgres:18-trixie AS pg-uuidv7-builder

RUN apt-get update \
    && apt-get install -y --no-install-recommends \
        build-essential git ca-certificates postgresql-server-dev-18 \
    && rm -rf /var/lib/apt/lists/*

ARG PG_UUIDV7_VERSION=v1.7.0
RUN git clone --depth 1 --branch ${PG_UUIDV7_VERSION} \
    https://github.com/fboulnois/pg_uuidv7.git /tmp/pg_uuidv7 \
    && make -C /tmp/pg_uuidv7 \
    && make -C /tmp/pg_uuidv7 install

FROM postgres:18-trixie

COPY --from=pg-uuidv7-builder /usr/lib/postgresql/18/lib/pg_uuidv7.so /usr/lib/postgresql/18/lib/
COPY --from=pg-uuidv7-builder /usr/share/postgresql/18/extension/pg_uuidv7* /usr/share/postgresql/18/extension/

EXPOSE 5432
COPY init/ /docker-entrypoint-initdb.d/
