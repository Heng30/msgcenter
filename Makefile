#!/bin/bash

APP_NAME=msgcenter
DATE=`date "+%Y_%m_%d"`
VERSION=`git tag | tail -n 1`

run: main.go version.go
	go run $^

build:
	rm -f ./msgcenter_v*
	echo "package main" > version.go
	echo "const VERSION string = \"${VERSION}\"" >> version.go
	go build -o ${APP_NAME}_${VERSION}_${DATE}

build-realse:
	rm -f ./msgcenter_*
	echo "package main" > version.go
	echo "const VERSION string = \"${VERSION}\"" >> version.go
	go build -ldflags "-s -w" -o ${APP_NAME}_${VERSION}_${DATE}

build-arm:
	echo "package main" > version.go
	echo "const VERSION string = \"${VERSION}\"" >> version.go
	env GOOS=linux GOARCH=arm go build -ldflags "-s -w" -o ${APP_NAME}_arm_${VERSION}_${DATE}

gen_denpendence:
	go mod tidy

download-dependence:
	go mod download

import-dependence:
	go mod vendor

gen-cert:
	cd cert && mkcert -install && mkcert localhost 127.0.0.1 ::1&& mv localhost+2-key.pem key.pem && mv localhost+2.pem cert.pem

ping:
	curl -v -H "Auth-Token: testToken" "http://localhost:8001/ping"

ping-tls:
	curl -v -H "Auth-Token: testToken" "https://localhost:8001/ping"

topic:
	curl -X POST -v -H "Auth-Token: testToken" -d '{"msg": "hello topic"}' "http://localhost:8001/topic/hello"


topic-tls:
	curl -X POST -v -H "Auth-Token: testToken" -d '{"msg": "hello topic"}' "https://localhost:8001/topic/hello"

