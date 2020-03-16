#!/bin/bash

#Script for building an application for *NIX

cd ..

git pull --all

go build -ldflags '-s -w' -v -o pass

pkill pass
cp --update pass /bin