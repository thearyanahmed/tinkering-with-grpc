#!/bin/bash

protoc sum/sumpb/sum.proto --go_out=plugins=grpc:.
protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.
