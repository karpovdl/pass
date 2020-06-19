#!/bin/bash

#./run.sh

#$1 - docker id
#$2 - app name or commad
#$3 - app version
#$4 - port
#$5 - pprof mode
#$6 - pprof port

docker stop $2_$3_$4
docker rm -v -f $2_$3_$4

docker run --restart always -d -it \
 --name $2_$3_$4 \
 --hostname=dp-$2_$3_$4 \
 --network=dp-net \
 -p $4:$4 \
 -e PORT="--port=$4" \
 -e PPROF="--pprof=$5" \
 -e PPROF_PORT="--pprof_port=$6" \
 $1/$2:$3

echo Success