#!/bin/bash
echo What should the version be?
read VERSION

if [[ -z $VERSION ]];
then
    echo "Remove previous latest image"
    docker rmi go-server --force
    echo "Building with version:latest"
    docker build -t go-server:latest .
else
    echo "Building with version:"$VERSION
    docker build -t go-server:$VERSION .
fi

echo "Removing intermediate go builder image"
docker image prune --filter label=stage=go-builder -f