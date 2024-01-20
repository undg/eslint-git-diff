#!/usr/bin/env sh

echo "[$(date +%Y-%m-%d_%H-%m_%S)]"
echo "========================================"
go test ./...

echo ''

while true; do
	inotifywait -qq -r -e create,close_write,modify,move,delete ./ &&
		echo "[$(date +%Y-%m-%d_%H-%m_%S)]" &&
		echo "========================================" &&
		go test ./...

	echo ''
done
