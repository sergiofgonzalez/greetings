PATH_TO_CHROME := "/mnt/c/Program Files/Google/Chrome/Application/chrome.exe"
WSL_PATH_PREFIX := "wsl.localhost/Ubuntu"


clean:
	rm -f .coverage.html
	rm .*.out

test:
	go test ./...

cover:
	go test -cover

coverhtml:
	go test -coverprofile=.coverage.out
	go tool cover -html=.coverage.out -o .coverage.html
	${PATH_TO_CHROME} file://${WSL_PATH_PREFIX}${PWD}/.coverage.html &
