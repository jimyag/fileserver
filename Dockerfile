FROM scratch
WORKDIR /app
COPY fileserver /app/fileserver
ENTRYPOINT [ "/app/fileserver"]