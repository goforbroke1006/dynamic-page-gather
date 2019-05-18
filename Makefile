SERVICE_NAME=dynamic-page-gather
BUILD_DIR=./build/Release
ENTRY_POINT=cmd/dynamic-page-gather/main.go

all: deps build

clean:
	rm -rf ./build/
	rm -rf ./vendor/

deps:
	dep ensure

.PHONY: build
build:
	mkdir -p ${BUILD_DIR}
	GOOS=linux   GOARCH=amd64 go build -o ${BUILD_DIR}/${SERVICE_NAME}        ${ENTRY_POINT}

install: build
#	go install github.com/goforbroke1006/dynamic-page-gather/cmd/dynamic-page-gather/
	sudo cp ${GOPATH}/bin/${SERVICE_NAME} /usr/local/bin/${SERVICE_NAME}
	sudo chmod +x /usr/local/bin/${SERVICE_NAME}

release:
	mkdir -p ${BUILD_DIR}
	GOOS=linux   GOARCH=amd64 go build -o ${BUILD_DIR}/${SERVICE_NAME}        ${ENTRY_POINT}
	GOOS=windows GOARCH=386   go build -o ${BUILD_DIR}/${SERVICE_NAME}.exe    ${ENTRY_POINT}
	GOOS=windows GOARCH=amd64 go build -o ${BUILD_DIR}/${SERVICE_NAME}_64.exe ${ENTRY_POINT}
	zip -r -j ${BUILD_DIR}/bin.zip ${BUILD_DIR}/*
