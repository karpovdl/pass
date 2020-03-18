docker stop $(docker ps -aqf "name=pass")
docker rm -v -f $(docker ps -aqf "name=pass")

docker build --no-cache -t pass-builder -f Dockerfile.builder .
docker rmi -f pass
docker run --rm --name pass-builder pass-builder | docker build --no-cache -t pass -f Dockerfile.run -
docker rmi -f pass-builder