#!/bin/bash

docker build \
	-t mathapp-deploy \
	-f Dockerfile.deploy_1.13 .
