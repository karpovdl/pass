docker stop $1_$2
docker rm -v -f $1_$2

docker run --restart always -d -it \
 -p $2:$2 \
 -e PORT="--port=$2" \
 -e PPROF="--pprof=$3" \
 -e PPROF_PORT="--pprof_port=$4" \
 --name $1_$2 \
 $1