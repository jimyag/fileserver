FROM golang:alpine as build
WORKDIR /app
COPY . .
RUN make build

FROM alpine:latest as release
WORKDIR /app
COPY --from=build /app/fileserver /app/fileserver
EXPOSE 8100

ENTRYPOINT [ "/app/fileserver"]