FROM golang:1.19-alpine

RUN apk update && apk add python3 py3-pip && pip install pyyaml

RUN mkdir /.cache && chmod 777 /.cache

RUN go install github.com/walterwanderley/sqlc-grpc@latest
RUN chmod -R 777 /go
CMD python3 misc/rpc/sqlc_grpc/downgrade_sqlc_config.py sqlc.yml sqlc.yaml && sqlc-grpc


