#!/bin/bash
set -e

if [ "$#" -ne 1 ] ; then
  echo "Specify a package to generate, or use \"all\" to generate all" >&2
  exit 1
fi

# takes an argument specifying which directory to generate code
function gen() {
  protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. \
  --go-grpc_opt=paths=source_relative \
  ./$1/*.proto
}

if [ $1 == "all" ]; then
  dirs=$(ls -d -- */)
  for dir in ${dirs[@]}; do
    gen $dir
  done
else
  gen $1
fi
