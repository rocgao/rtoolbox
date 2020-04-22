#!/bin/sh

commit_id=$(git -P rev-parse HEAD)
echo "${commit_id}"

ver=$(git -P tag --contains "${commit_id}")
echo "${ver}"

dt=$(date -u '+%Y-%m-%d_%I:%M:%S%p')
echo "${dt}"

flags="-X main.Version=${ver} -X main.GitCommit=${dt}"
echo "${flags}"

echo "Building package for Darwin"
out_file="rtoolbox-${ver}-drawin-amd64"
GOOS=darwin GOARCH=amd64  go build -ldflags "${flags}"  -o "${out_file}" ./cmd/rtoolbox
zip "${out_file}.zip" "${out_file}" > /dev/null
echo "Completed"

echo "Building package for Windows"
out_file="rtoolbox-${ver}-win-amd64"
GOOS=windows GOARCH=amd64  go build -ldflags "${flags}"  -o "${out_file}.exe" ./cmd/rtoolbox
zip "${out_file}.zip" "${out_file}.exe" > /dev/null
echo "Completed"

echo "Building package for Linux"
out_file="rtoolbox-${ver}-linux-amd64"
GOOS=linux GOARCH=amd64  go build -ldflags "${flags}"  -o "${out_file}" ./cmd/rtoolbox
zip "${out_file}.zip" "${out_file}" > /dev/null
echo "Completed"