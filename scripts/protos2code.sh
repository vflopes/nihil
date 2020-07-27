#!/bin/bash

cd pkg || exit 1

targets=(
  "analysis"
)

for path_name in "${targets[@]}"; do
  echo -n "$path_name: "

  echo "Compiling Golang..."
  protoc \
    -I ./ \
    --go_out=paths=source_relative:. \
    --go-grpc_out=paths=source_relative:. \
    ./"$path_name"/*.proto

  echo "Done."
done