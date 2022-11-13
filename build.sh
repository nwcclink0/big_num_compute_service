#!/bin/bash

TARGET=$1
GO_OSARCH=
GO_BUILD_TAGS=
OUTPUT_SUBDIR=
GO_BUILD_TAGS=""
if [ "$TARGET" = "lb" ]; then
  export GOOS=linux
  export GOARCH=amd64
  go build -v -o="./load_balancer/service/big_num_compute_service"
  docker build -t yuantingwei/big_num_compute_service:v0.1 -f ./load_balancer/service/Dockerfile ./load_balancer/service
  cd load_balancer || echo "load_balancer don't exist" exit
  docker-compose build
  docker-compose up --scale big_num_compute_service=2
elif [ "$TARGET" = "native" ]; then
  go build -v
elif [ "$TARGET" = "github_action" ]; then
  go build -v -o="./load_balancer/service/big_num_compute_service"
else
  echo "./build [load_balancer/native]"
fi
