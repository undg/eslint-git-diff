#!/usr/bin/env sh

go build.sh

while true
do
    inotifywait -qq -r -e create,close_write,modify,move,delete ./ && go build.sh
done

