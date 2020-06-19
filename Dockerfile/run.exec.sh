#!/bin/bash

#./run.exec.sh

### PASS
#$1 - app name or commad
#$2 - app version
#$3 - port
#$4 - pprof mode
#$5 - pprof port

docker_id="karpovdl"
app_name="pass"
app_version="$1"

./builder.sh
./builder.sh $docker_id $app_name $app_version

app_port="19300"
app_pprof_mode="false"
app_pprof_port="19301"

./run.sh \
  $docker_id \
  $app_name \
  $app_version \
  $app_port \
  $app_pprof_mode \
  $app_pprof_port

echo Success