#!/bin/bash
echo What should be the version of image?
read VERSION

echo What is the file_name go executable \(/dist/go_server_9999\)?
read FILENAME

if [[ -z $VERSION ]];
then
    FILENAME="go_server_9999"
fi


if [[ -z $VERSION ]];
then
    echo \n"Remove previous latest image"
    docker rmi go-server --force
    echo \n"Building with version:latest"
    docker build --build-arg file_name=$FILENAME -t go-server:latest .
else
    echo \n"Building with version:"$VERSION
    docker build -t go-server:$VERSION .
fi