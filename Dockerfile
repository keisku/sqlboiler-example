FROM golang:1.13.1-alpine3.10

EXPOSE 3000

ENV APP_DIR $GOPATH/src/github.com/kskumgk63/sqlboiler-example
ENV PATH $GOPATH/bin:$PATH
ENV GO111MODULE on

RUN apk update && \
    apk add git make gcc g++

RUN go get -u github.com/pressly/goose/cmd/goose
RUN GO111MODULE=off && go get -u github.com/oxequa/realize
RUN go get -u golang.org/x/tools/cmd/goimports

ADD . $APP_DIR
WORKDIR $APP_DIR

CMD while ! nc -z ${PSQL_HOST} ${PSQL_PORT}; do sleep 1; done && \
    goose -dir sql postgres "host=${PSQL_HOST} \
    port=${PSQL_PORT} \
    user=${PSQL_USER} \
    password=${PSQL_PASSWORD} \
    dbname=${PSQL_DBNAME} \
    sslmode=${PSQL_SSLMODE}" up && \
    realize start
