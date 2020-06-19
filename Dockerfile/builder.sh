#!/bin/bash

#Script make image
#./builder.sh

#$1 - docker id
#$2 - app name
#$3 - app version

docker_id="$1"
app_name="$2"
app_version="$3"

if [ $docker_id != "" && $app_name != "" && $app_version != "" ]; then

  # PreDelete artifactory
  docker ps -aqf "name=$app_name" | xargs -I'{}' docker stop $(docker ps -aqf "name=$app_name")
  docker ps -aqf "name=$app_name" | xargs -I'{}' docker rm -v -f $(docker ps -aqf "name=$app_name")

  docker images | grep $docker_id/$app_name | grep latest | xargs -I'{}' docker rmi -f $docker_id/$app_name
  docker images | grep $docker_id/$app_name | grep $app_version | xargs -I'{}' docker rmi -f $docker_id/$app_name:$app_version
  docker images | grep $docker_id/$app_name | grep builder | grep latest | xargs -I'{}' docker rmi -f $docker_id/$app_name-builder
  docker images | grep $docker_id/$app_name | grep builder | grep $app_version | xargs -I'{}' docker rmi -f $docker_id/$app_name-builder:$app_version

  # Build
  docker build --no-cache --build-arg app_df="/Dockerfile.run" app_version=$app_version -t $docker_id/$app_name-builder:$app_version -f Dockerfile.builder .
  docker run --rm --name $app_name-builder $docker_id/$app_name-builder:$app_version | docker build --no-cache -t $docker_id/$app_name:$app_version -f Dockerfile.run -
  docker tag $docker_id/$app_name:$app_version $docker_id/$app_name

  # PostDelete artifactory
  docker ps -aqf "name=$app_name" --filter status=exited | xargs -I'{}' docker rm -v -f $(docker ps -aqf "name=$app_name" --filter status=exited)
  docker images | grep $docker_id/$app_name | grep builder | xargs -I'{}' docker rmi -f $docker_id/$app_name-builder:$app_version

else

  echo None build

fi

echo Success