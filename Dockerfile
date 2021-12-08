FROM golang:1.17.4-alpine3.15 AS golang

WORKDIR /app
COPY . /app
RUN go build -o execution-environment-tester -mod=mod main.go

FROM alpine:3.15

WORKDIR /app
COPY --from=golang /app/execution-environment-tester /app/execution-environment-tester
RUN chmod +x /app/execution-environment-tester

EXPOSE 1323
CMD ["/app/execution-environment-tester"]
