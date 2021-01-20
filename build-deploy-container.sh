#!/bin/bash

docker build \
	-t go-pg-app-deploy \
	-f Dockerfile.deploy_1.13 .
