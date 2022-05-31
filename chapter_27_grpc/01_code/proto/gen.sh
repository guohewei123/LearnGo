#!/bin/bash

#protoc -I=. --go_out=plugins=grpc,paths=source_relative:gen/trip trip.proto
#protoc -I=. --go_out=plugins=grpc,paths=source_relative:gen/importtest importtest.proto
#protoc -I=. --go_out=plugins=grpc,paths=source_relative:gen/common common.proto

for proto in ./*.proto;
do
  protofile=${proto##*/}
  pack_dir=gen/${protofile%.*};
  mkdir -p ${pack_dir}
  cmd="protoc -I=. --go_out=plugins=grpc,paths=source_relative:${pack_dir} ${protofile}"
  echo $cmd
  bash -c "${cmd}"
done

echo "------------grpc-gateway--------"
gateway_cmd="protoc -I=. --grpc-gateway_out=paths=source_relative,grpc_api_configuration=trip.yaml:gen/trip trip.proto"
echo $gateway_cmd
bash -c "${gateway_cmd}"
