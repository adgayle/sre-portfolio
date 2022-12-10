#!/bin/bash
IMAGE='rancher/k3s:v1.24.4-k3s1'
PROJECT_NAME='cka'
SERVER_COUNT=3

echo "Creating k3d cluster ${PROJECT_NAME}."
k3d cluster create \
	${PROJECT_NAME} \
	--image ${IMAGE} \
	--no-lb \
	--registry-create ${PROJECT_NAME}-registry \
	--servers ${SERVER_COUNT}
