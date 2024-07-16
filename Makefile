binary=fileserver
ldflags="-s -w"
base_build=-v -trimpath

build:
	go build ${base_build} -ldflags ${ldflags} -o ${binary} ./

image:
	docker build -t ${binary}:dev --target release .