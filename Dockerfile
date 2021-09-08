FROM golang as builder
ENV APP_USER app
ENV APP_HOME /go/src/app
RUN groupadd $APP_USER && useradd -m -g $APP_USER -l $APP_USER
RUN mkdir -p $APP_HOME && chown -R $APP_USER:$APP_USER $APP_HOME
WORKDIR $APP_HOME
USER $APP_USER
COPY . .
RUN go mod download
RUN go mod verify
RUN go build -o app cmd/main.go

FROM debian:buster
FROM golang
ENV APP_USER app
ENV APP_HOME /go/src/app
RUN groupadd $APP_USER && useradd -m -g $APP_USER -l $APP_USER
RUN mkdir -p $APP_HOME
WORKDIR $APP_HOME
COPY --chown=0:0 --from=builder $APP_HOME/app $APP_HOME
EXPOSE 8000
USER $APP_USER
CMD ["sh", "-c", "./app"]