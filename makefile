run_test=go test -v ./... | \
			sed ''/PASS/s//$(printf "\033[32mPASS\033[0m")/'' | \
			sed ''/FAIL/s//$(printf "\033[31mFAIL\033[0m")/'' | \
			sed ''/RUN/s//$(printf "\033[33mRUN\033[0m")/''

build:
	go build -ldflags "-X eslint-git-diff/cmd.stringVersion=$(git describe --abbrev=0 --tags)" -v -o ~/bin/eslint-git-diff
	# echo 'Save executable in ~/bin/eslint-git-diff'

watch:
	./build.sh
		go build -ldflags "-X eslint-git-diff/cmd.stringVersion=$(git describe --abbrev=0 --tags)" -v -o ~/bin/eslint-git-diff

	while true
	do
			inotifywait -qq -r -e create,close_write,modify,move,delete ./ && ./build.sh
	done

run:
	go run main.go

test:
	./test-watch.sh



test_coverage:
	go test ./... -coverprofile=coverage.out
#
# test_watch:
#
# 	echo "[$(date +%Y-%m-%d_%H-%m_%S)]"
# 	echo "========================================"
#
# 	run_test
#
# 	echo ''
#
# 	while true; do
# 		inotifywait -qq -r -e create,close_write,modify,move,delete ./ &&
# 			echo "[$(date +%Y-%m-%d_%H-%m_%S)]" &&
# 			echo "========================================" &&
# 			run_test
#
# 		echo ''
# 	done
