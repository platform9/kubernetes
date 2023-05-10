#!/bin/bash

# This file will build k8s image. retag it and push image to dockerhub
make quick-release-images KUBE_EXTRA_WHAT="cmd/kubectl" KUBE_BUILD_CONFORMANCE=n KUBE_BUILD_HYPERKUBE=n KUBE_VERBOSE=0 KUBE_BUILD_PLATFORMS=linux/amd64 KUBE_DOCKER_REGISTRY=platform9

# Load the Docker image from the TAR file
docker load -i $TAR_FILE

# Extract the original tag from the image ID based on the original image name
ORIGINAL_IMAGE_NAME="platform9/kube-proxy-amd64"
ORIGINAL_TAG=$(docker images --format '{{.Tag}}' $ORIGINAL_IMAGE_NAME | head -n 1)

# Extract the version number from the original tag
VERSION=$(echo $ORIGINAL_TAG | cut -d '-' -f 1)

# Adding PMK build tag
VERSION="$VERSION-pmk-$TEAMCITY_BUILD_ID"

# Retag the Docker image with the new name and version number
docker tag $ORIGINAL_IMAGE_NAME:$ORIGINAL_TAG platform9/$NEW_IMAGE_NAME:$VERSION

echo platform9/$NEW_IMAGE_NAME:$VERSION >> "_output/release-images/amd64/image.txt"

# Verify that the image has been retagged
docker images platform9/$NEW_IMAGE_NAME

# Push the retagged image to Docker Hub
docker push platform9/$NEW_IMAGE_NAME:$VERSION