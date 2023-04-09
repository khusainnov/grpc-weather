#!/bin/bash

image_name=grpc-weather-app_v1.0:latest

docker pull image_name

if [ $? -eq 0 ]; then
  echo "Image pulled successfully"
else
  echo "Error pulling image"
  exit 1
fi