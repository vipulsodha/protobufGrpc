#!/usr/bin/env bash
protoc -I simple/ --go_out=simple/ simple/*.proto
protoc -I simple_enum/ --go_out=simple_enum/ simple_enum/*.proto
protoc -I complex/ --go_out=plugins=grpc:./complex complex/*.proto
