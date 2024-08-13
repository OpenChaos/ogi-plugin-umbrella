#!/usr/bin/env bash

export CONSUMERS=(ogi-http-service-consumer ogi-file-consumer)
export TRANSFORMERS=(ogi-yfinance-transformer)
export PRODUCERS=(ogi-http-producer ogi-file-producer ogi-yfin52low-producer ogi-slack-producer)


list-all-plugin-paths(){
  find . -type f -name 'go.mod' | sed 's/go.mod$//'
}

## build
build-plugin(){
  local plugin_name="$1"
  local plugin_path="$2"

  set -e
  echo "[+] BUILDING ${plugin_name}"
  pushd "${plugin_path}" > /dev/null
  GO111MODULE=on GOOS=linux GOARCH=amd64 \
    go build -o "../../out/${plugin_name}.so" -buildmode=plugin .
  popd > /dev/null
  ls -lah "./out/${plugin_name}.so"
  set +e
  echo '-------------------------------'
  echo ''
}

build-all-consumers(){
  for p in ${CONSUMERS[@]}; do
    build-plugin "${p}" "./consumers/${p}"
  done
}

build-all-transformers(){
  for p in ${TRANSFORMERS[@]}; do
    build-plugin "${p}" "./transformers/${p}"
  done
}

build-all-producers(){
  for p in ${PRODUCERS[@]}; do
    build-plugin "${p}" "./producers/${p}"
  done
}

build-all-plugins(){
  build-all-consumers
  build-all-transformers
  build-all-producers
}


case $1 in
  build-consumer*)
    build-all-consumers
    ;;
  build-transformer*)
    build-all-transformers
    ;;
  build-producer*)
    build-all-producers
    ;;
  build*)
    build-all-plugins
    ;;
  ls*)
    list-all-plugin-paths
    ;;
  **)
    echo "usage: $0 <build|ls|build-consumer|build-transformer|build-producer>"
    ;;
esac
