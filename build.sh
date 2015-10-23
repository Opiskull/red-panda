#!/bin/bash

function GenerateProtos {
	for f in "service.*/proto/*.proto";	do
    echo compile-proto: $f;
		protoc --go_out=plugins=grpc:. $f;
	done
}

function BuildServices {
  for service in $(ls -d $1.*/); do
    BuildService $service
  done
}

function BuildService {
  file="main.go"
  if [ -e "$1/$file" ]; then
    echo build: $1;
    cd $1;
    go build;
    cd ..
  fi
}

function BuildAllServices {
  BuildServices service
  BuildServices api
}

case $1 in
  protos|proto)
  GenerateProtos
  ;;
  build|build-all)
  if [ -n $2 ]; then
    BuildService $2
  else
    BuildAllServices
  fi
  ;;
  build-service)
  BuildServices service
  ;;
  build-api)
  BuildServices api
  ;;
  *)
  GenerateProtos
  BuildAllServices
  ;;
esac
