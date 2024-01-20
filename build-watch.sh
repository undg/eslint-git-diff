#!/usr/bin/env sh

./build.sh

while true
do
    inotifywait -qq -r -e create,close_write,modify,move,delete ./ && ./build.sh
done

