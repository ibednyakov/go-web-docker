#!/bin/bash

docker build \
	-t restful-db-app-deploy \
	-f Dockerfile.deploy_1.13 .
