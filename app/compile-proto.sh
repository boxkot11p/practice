#!/bin/bash

if [ $# -ne 1 ]; then
  echo "needs output folder arg" 1>&2
  exit 1
fi

files=`(find ./ -name '*.proto')`
for i in $files; do
  protoc -I . \
    --go_out $1 \
    --go_opt paths=source_relative \
    --go-grpc_out $1 \
    --go-grpc_opt paths=source_relative \
    "${i}"
done