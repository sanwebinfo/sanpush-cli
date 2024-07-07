BUILD_DIR=./build

clean:
	rm -rf ${BUILD_DIR}

build:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64       go build -o build/sanpush-windows-amd64.exe sanpush.go
	CGO_ENABLED=0 GOOS=windows GOARCH=arm64       go build -o build/sanpush-windows-arm64.exe sanpush.go
	CGO_ENABLED=0 GOOS=windows GOARCH=386         go build -o build/sanpush-windows-386.exe   sanpush.go
	CGO_ENABLED=0 GOOS=linux   GOARCH=amd64       go build -o build/sanpush-linux-amd64       sanpush.go
	CGO_ENABLED=0 GOOS=linux   GOARCH=386         go build -o build/sanpush-linux-386         sanpush.go
	CGO_ENABLED=0 GOOS=linux   GOARCH=arm64       go build -o build/sanpush-linux-arm64       sanpush.go
	CGO_ENABLED=0 GOOS=linux   GOARCH=arm GOARM=7 go build -o build/sanpush-linux-arm-7       sanpush.go
	CGO_ENABLED=0 GOOS=darwin  GOARCH=amd64       go build -o build/sanpush-darwin-amd64      sanpush.go
	CGO_ENABLED=0 GOOS=darwin  GOARCH=arm64       go build -o build/sanpush-darwin-arm64      sanpush.go
	CGO_ENABLED=0 GOOS=android GOARCH=arm64       go build -o build/sanpush-android-arm64      sanpush.go