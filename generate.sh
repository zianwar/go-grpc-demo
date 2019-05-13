#bin/bash

protoc calcpb/calulator.proto --go_out=plugins=grpc:.
