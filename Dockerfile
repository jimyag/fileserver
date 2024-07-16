FROM golang:alpine as build
WORKDIR /app
COPY . .
RUN go build -v -trimpath -ldflags "-s -w" -o fileserver ./

FROM alpine:latest as release
WORKDIR /app
COPY --from=build /app/fileserver /app/fileserver
EXPOSE 8100

ENTRYPOINT [ "/app/fileserver"]