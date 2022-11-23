#!/bin/bash  

ROOT=..
export CONF_PATH=$ROOT/config/conf.yaml
BIN=jun2

mkdir -p $ROOT/bin

cd $ROOT/cmd

echo 'COMPILING...'
rm $ROOT/bin/$BIN
go build -o $BIN

mv $BIN $ROOT/bin

echo 'RUN'
$ROOT/bin/$BIN $@