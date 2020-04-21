export ver = git -P tag --contains $(git rev-parse HEAD)
export dt = date -u '+%Y-%m-%d_%I:%M:%S%p'
darwin :
	GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.Version=`$(ver)` -X main.GitCommit=`$(dt)`" -o rtoolbox-v$(ver)-drawin-amd64 ./cmd/rtoolbox
